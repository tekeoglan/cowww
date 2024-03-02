package cowww

import (
	"errors"
	"fmt"
	"log"
	"net"
)

const DefaultProto = "HTTP/1.1"

var handlers = make(map[string]Handler)

type Handler func(req *HttpRequest, resp ResponseWriter)

type Header map[string]string

func (h Header) Set(key, val string) {
	if h == nil {
		return
	}

	h[key] = val
}

func (h Header) Get(key string) string {
	if h == nil {
		return ""
	}

	return h[key]
}

func (h Header) Del(key string) {
	if h == nil {
		return
	}

	delete(h, key)
}

// Handle registers a handler for the given path
func Handle(path string, handler Handler) {
	handlers[path] = handler
}

func handleRequest(c net.Conn) error {
	if c == nil {
		return errors.New("Connection is nil")
	}

	defer c.Close()

	req, err := parseHttpRequest(c)
	if err != nil {
		return errors.New(fmt.Sprintf("Error parsing request: %v", err))
	}

	log.Printf("%s %s %s", req.Method, req.Url, req.Proto)

	if handler, ok := handlers[req.Url]; ok {
		res := &response{
			c:             c,
			req:           req,
			statusCode:    StatusOk,
			status:        statusText(StatusOk),
			handlerHeader: Header{},
		}
		handler(req, res)
	} else {
		c.Write([]byte(NotFoundError.Error()))
	}

	return nil
}
