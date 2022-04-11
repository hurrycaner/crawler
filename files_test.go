package main

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"testing"
)

func TestLinkToFilepath(t *testing.T) {
	link1, _ := url.Parse("http://example.com/abc")
	assert.Equal(t, "example.com/abc/index.html", LinkToFilepath(link1))

	link2, _ := url.Parse("http://example.com/abc/def")
	assert.Equal(t, "example.com/abc/def/index.html", LinkToFilepath(link2))

	link3, _ := url.Parse("http://example.com/abc/xyz.html")
	assert.Equal(t, "example.com/abc/xyz.html", LinkToFilepath(link3))

	link4, _ := url.Parse("http://example.com/abc.svg")
	assert.Equal(t, "example.com/abc.svg", LinkToFilepath(link4))
}

func TestFetchExistingFilepaths(t *testing.T) {
	baseURL, _ := url.Parse("http://example.com/abc")
	filepaths, _ := FetchExistingFilepaths(baseURL, "testdata")
	assert.Equal(t, []string{
		"example.com/abc/def/index.html",
		"example.com/abc/index.html",
		"example.com/abc/xyz.html",
	}, filepaths)
}

func TestWrite(t *testing.T) {
	path := os.TempDir() + "example.com/abc/def/index.html"
	err := Write(path, []byte("test"))
	assert.Equal(t, nil, err)

	content, err := os.ReadFile(path)
	assert.Equal(t, nil, err)
	assert.Equal(t, "test", string(content))
}
