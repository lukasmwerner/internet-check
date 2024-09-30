package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lukasmwerner/internet-check/tester"
)

type Entry struct {
	Host       string
	StatusCode int
}

func main() {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err.Error())
		return
	}

	configFilePath := path.Join(home, ".config", "internet-check", "hosts.json")

	hosts := []Entry{}

	b, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		err := json.Unmarshal(b, &hosts)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	p := tea.NewProgram(initalModel(hosts))

	go func() {
		var wg sync.WaitGroup
		type Status struct {
			Index int
			Error error
			Ok    bool
		}
		statuses := make(chan Status)
		for i, host := range hosts {
			wg.Add(1)
			go func(idx int, host Entry) {
				ok, err := tester.TestHTTPConnection(host.Host, host.StatusCode)
				statuses <- Status{
					Index: idx,
					Error: err,
					Ok:    ok,
				}
				wg.Done()
			}(i, host)

		}

		go func() {
			wg.Wait()
			close(statuses)
		}()

		for v := range statuses {
			if v.Ok {
				p.Send(successMsg{
					Index:    v.Index,
					Finished: true,
				})
			} else {
				p.Send(failureMsg{
					Index: v.Index,
					Error: v.Error.Error(),
				})
			}
		}
		p.Send(tea.QuitMsg{})

	}()

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
