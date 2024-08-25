package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	checkMark = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	xMark     = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).SetString("X")
)

type successMsg struct {
	Index    int
	Finished bool
}
type failureMsg struct {
	Index int
	Error string
}

type model struct {
	hosts    []string
	finished []bool
	success  []bool
	errors   []string

	spinner spinner.Model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}
	case successMsg:
		m.success[msg.Index] = true
		m.finished[msg.Index] = true
		return m, nil
	case failureMsg:
		m.success[msg.Index] = false
		m.finished[msg.Index] = true
		m.errors[msg.Index] = msg.Error
		return m, nil
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	header := "internet-check \n"

	render := ""
	for i, host := range m.hosts {
		if !m.finished[i] {
			render += fmt.Sprintf(" %s %s\n", m.spinner.View(), host)
		} else if m.success[i] {
			render += fmt.Sprintf("  %s  %s\n", checkMark.String(), host)
		} else {
			render += fmt.Sprintf("  %s  %s: %s\n", xMark.String(), host, m.errors[i])
		}
	}

	return header + render
}

func initalModel(hosts []string) model {

	m := model{
		hosts:    hosts,
		finished: make([]bool, len(hosts)),
		success:  make([]bool, len(hosts)),
		errors:   make([]string, len(hosts)),
		spinner:  spinner.New(spinner.WithSpinner(spinner.Points)),
	}
	return m
}
