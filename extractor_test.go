package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestExtract(t *testing.T) {
	baseURL, _ := url.Parse("http://example.com/abc")
	content := []byte(`<html>
		<head>
			<title>Example</title>
		</head>
		<body>
			<h1>Hello, World!</h1>
			<a href="http://example.com/def">Link 1</a>
			<a href="http://example.com/abc">Link 2</a>
			<a href="http://example.com/">Link 3</a>
			<a href="http://example.com/abc/def">Link 4</a>
			<a href="test.html">Link 5</a>
			<a href="/abc/xyz">Link 6</a>
		</body>
	</html>`)

	links, err := Extract(content, baseURL)
	require.Equal(t, nil, err)
	require.Equal(t, 4, len(links))
	assert.Equal(t, "http://example.com/abc/", links[0].String())
	assert.Equal(t, "http://example.com/abc/def/", links[1].String())
	assert.Equal(t, "http://example.com/abc/test.html", links[2].String())
	assert.Equal(t, "http://example.com/abc/xyz/", links[3].String())
}
