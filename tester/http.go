package tester

import (
	"fmt"
	"net/http"
)

func TestHTTPConnection(link string) (bool, error) {

	resp, err := http.Get(link)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("got non 200 status code: %s", resp.Status)
	}

	return true, nil
}
