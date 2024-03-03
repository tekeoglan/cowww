# Cowww
This is a minimalistic web framework, developed  without using `net/http` package for learning inner workings of a web framework. It only supports HTTP/1.1 and lacks many features.

## Feature Overview
- Handling requests
- Parsing incoming requests
- Structuring responses

## Installation
```sh
git clone github.com/tekeoglan/cowww
```

## Usage
Create a `cmd` directory in root of project and add `main.go` in it.
```go
package main

import "github.com/tekeoglan/cowww"

func main() {
	cowww.Handle("/", func(req *cowww.HttpRequest, resp cowww.ResponseWriter) {
		resp.Header().Set("Content-Type", "text/plain")
		resp.Write([]byte("Hello, World!"))
	})

	config := cowww.ServerConfig{
		Host: "localhost",
		Port: "8080",
	}

	server := cowww.NewServer(config)

	server.Start()
}
```
Run `go run ./cmd`

After running, open <http://localhost:8080> in your browser.

## License
[MIT](https://github.com/tekeoglan/cowww/blob/main/LICENSE)
