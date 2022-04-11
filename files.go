package main

import (
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// Write the content to the filepath
func Write(path string, content []byte) error {
	dir, _ := filepath.Split(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LinkToFilepath retrieve filepath from that link URL
func LinkToFilepath(link *url.URL) string {
	dir, file := filepath.Split(link.Host + link.Path)
	if filepath.Ext(file) == "" {
		file += "/index.html"
	}
	return filepath.Join(dir, file)
}

// FetchExistingFilepaths resume website crawling by checking existing files
func FetchExistingFilepaths(baseURL *url.URL, destDir string) ([]string, error) {
	baseDir := filepath.Clean(destDir) + string(filepath.Separator)
	path := filepath.Join(baseDir, baseURL.Host, baseURL.Path)
	files, err := fetchDirectory(path, baseDir)
	if err != nil {
		return nil, err
	}
	return files, nil
}
func fetchDirectory(path string, destDir string) ([]string, error) {
	var filepaths = []string{}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return nil, err
		}
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			files, err := fetchDirectory(filepath.Join(path, file.Name()), destDir)
			if err != nil {
				continue
			}
			filepaths = append(filepaths, files...)
		} else {
			filepaths = append(filepaths, filepath.Join(strings.TrimPrefix(path, destDir), file.Name()))
		}
	}
	return filepaths, nil
}
