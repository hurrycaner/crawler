package main

import (
	"bytes"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

// Extract the URLs from the file
func Extract(content []byte, baseURL *url.URL) ([]*url.URL, error) {
	b := bytes.NewReader(content)
	page, err := html.Parse(b)
	if err != nil {
		return nil, err
	}
	baseURL = fixURL(baseURL)
	links := parseNode(page, baseURL)
	return links, nil
}

// parseNode extract links, used this as example: https://pkg.go.dev/golang.org/x/net/html#example-Parse
func parseNode(n *html.Node, baseURL *url.URL) []*url.URL {
	links := []*url.URL{}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, parseNode(c, baseURL)...)
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				href, err := url.Parse(a.Val)
				if err != nil {
					break
				}
				href = fixURL(href)
				uri := baseURL.ResolveReference(href)
				if strings.HasPrefix(uri.String(), baseURL.String()) {
					links = append(links, uri)
				}
				break
			}
		}
	}
	return links
}

func fixURL(baseURL *url.URL) *url.URL {
	urlParts := strings.Split(baseURL.Path, "/")
	if !strings.Contains(urlParts[len(urlParts)-1], ".") && baseURL.String()[len(baseURL.String())-1] != '/' {
		baseURL.Path += "/"
	}
	return baseURL
}
