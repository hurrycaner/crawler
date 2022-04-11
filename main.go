package main

import (
	"flag"
	"fmt"
	"net/url"
	"path/filepath"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Println("Usage: go run main.go <url> <destination>")
		return
	}
	RootURL := flag.Arg(0)
	DestDir := filepath.Clean(flag.Arg(1)) + string(filepath.Separator)

	var filepaths = map[string]bool{}
	var queue = []*url.URL{}
	link, err := url.Parse(RootURL)
	if err != nil {
		panic(err)
	}
	queue = append(queue, link)

	existing, err := FetchExistingFilepaths(link, DestDir)
	if err != nil {
		panic(err)
	}
	for _, file := range existing {
		if _, ok := filepaths[file]; !ok {
			filepaths[file] = false
		}
	}
	Crawl(queue, filepaths, link, DestDir)
	fmt.Printf("%d files downloaded\n", len(filepaths))
}

func Crawl(queue []*url.URL, filepaths map[string]bool, baseURL *url.URL, destination string) {
	for len(queue) > 0 {
		link := queue[0]
		queue = queue[1:]
		path := LinkToFilepath(link)
		var content []byte
		var err error
		if extracted, exists := filepaths[path]; exists {
			if extracted {
				// already extracted
				continue
			}
			content, err = Read(filepath.Join(destination, path))
			if err != nil {
				continue // should log and handle error here
			}
		} else {
			// only download if file does not exist
			content, err = FetchURL(link)
			if err != nil {
				continue // should log and handle error here
			}
			err = Write(filepath.Join(destination, path), content)
			if err != nil {
				continue // should log and handle error here
			}
			filepaths[path] = false
		}
		newLinks, err := Extract(content, baseURL)
		if err != nil {
			continue // should log and handle error here
		}
		filepaths[path] = true
		for _, link := range newLinks {
			newPath := LinkToFilepath(link)
			if extracted, exists := filepaths[newPath]; !exists || !extracted {
				// would add a func with mutex to handle concurrency
				queue = append(queue, link)
			}
		}
	}
}
