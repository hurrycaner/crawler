package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// FetchURL fetch the URL and return the content
func FetchURL(link *url.URL) ([]byte, error) {
	resp, err := http.Get(link.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
