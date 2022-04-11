# Test task: Backend Developer

Implement a recursive, mirroring web crawler.

The crawler should be a command-line tool that accepts a starting URL and a destination directory.
The crawler will then download the page at the URL, save it in the destination directory, and then recursively proceed to any valid links in this page.

A valid link is the value of an `href` attribute in an `<a>` tag the resolves to urls that are children of the initial URL. For example, given initial URL `https://start.url/abc`, URLs that resolve to `https://start.url/abc/foo` and `https://start.url/abc/foo/bar` are valid URLs, but ones that resolve to `https://another.domain` or to `https://start.url/baz` are not valid URLs, and should be skipped.

Additionally, the crawler should:
- Correctly handle being interrupted by `Ctrl-C`
- Perform work in parallel where reasonable
- Support resume functionality by checking the destination directory for downloaded pages, and skip downloading and processing where not necessary

The expected timeframe for this assignment is ~4 hours. Please do not spend more time on it. If you do not think you can implement all the features in this timeframe, prioritize the features you consider most important.

The implementation must be written in Go.

Some tips:

- If you’re not familiar with this kind of software, see `wget --mirror` for very similar functionality
- Document features that are missing, and any other changes you would make if you had a longer time to implement the assignment
