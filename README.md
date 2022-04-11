# Somewhat web crawler mirror

Usage:

```shell
go run main.go <url> <destination>

> go run main.go http://www.w3school.com/cpp websites
```

Developed using Golang version 1.18

Missing features:
- File processing to guarantee that links are relative to base directory
- Handle Ctrl-C
  > Not top priority feature while no concurrency implemented
- Concurrency 
  > - I would add a guard channel to limit number of concurrent goroutines
  > - Mutex to handle queue pop and push operations
  > - Mutex to handle filepaths map read and write operations
  > - At first, would add concurrency to:
  >   - FetchURL
  >   - Extract / ParseNode
- Better error handling and logging
- Better code isolation
  > Would add each file to a different package, with interfaces.

- Missing tests for main.Crawl
  > Would use mockery and testify/mock package

Considerations: 
- I've never created anything similar to a web crawler/mirror before, it was interesting, but it felt a bit odd, since it's more focused in coding than architectural challenge (maybe another challenge will come after this?)
- Most interesting features couldn't be completed in 4 hours, sorry :() 
- I've lost some time with ResolveReference not working as I expected with relative URLs and base URLs without a trailing slash. Example: 
 
```go
package main

import (
  "fmt"
  "net/url"
)

func main() {
  u, _ := url.Parse("cpp_intro.asp")
  
  b, _ := url.Parse("https://www.w3scools.com/cpp")
  fmt.Println(b.ResolveReference(u))
  // returns "https://www.w3scools.com/cpp_intro.asp"
  b, _ = url.Parse("https://www.w3scools.com/cpp/")
  fmt.Println(b.ResolveReference(u))
  // returns of "https://www.w3scools.com/cpp/cpp_intro.asp"
}
``` 
