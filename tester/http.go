package tester

import (
	"fmt"
	"net/http"
)

func TestHTTPConnection(link string, expectedCode int) (bool, error) {

	resp, err := http.Get(link)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != expectedCode {
		return false, fmt.Errorf("got status code: %s not %d", resp.Status, expectedCode)
	}

	return true, nil
}
