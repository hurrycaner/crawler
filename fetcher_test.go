package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestFetchURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/error":
			w.WriteHeader(404)
			w.Write([]byte(strings.TrimSpace(`Not Found`)))
		case "/success":
			w.WriteHeader(200)
			w.Write([]byte(strings.TrimSpace(`OK`)))
		}
	}))
	linkSuccess, _ := url.Parse(ts.URL + "/success")
	body, err := FetchURL(linkSuccess)
	require.Equal(t, nil, err)
	assert.Equal(t, "OK", string(body))

	linkError, _ := url.Parse(ts.URL + "/error")
	body, err = FetchURL(linkError)
	require.Equal(t, "status code error: 404 404 Not Found", err.Error())
	assert.Equal(t, "", string(body))
}
