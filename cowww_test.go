package cowww

import (
	"strings"
	"testing"

	"github.com/tekeoglan/cowww/mocks"
)

func TestHeader(t *testing.T) {
	var h Header

	t.Log("Inserting a value inside not initialized header should not panic")
	h.Set("Content-Type", "text/plain")

	t.Log("Getting a value inside not initialized header should not panic")
	h.Get("Content-Type")

	t.Log("Deleting a value inside not initialized header should not panic")
	h.Del("Content-Type")

	h = Header{}

	h.Set("Content-Type", "text/plain")
	if h["Content-Type"] != "text/plain" {
		t.Error("Handler header should contain 'text/plain'")
	}

	got := h.Get("Content-Type")
	if got != "text/plain" {
		t.Error("Handler header should contain 'text/plain'")
	}

	h.Del("Content-Type")
	if h.Get("Content-Type") != "" {
		t.Error("Handler header should be empty")
	}
}

func TestHandle(t *testing.T) {
	handler := func(req *HttpRequest, resp ResponseWriter) {
		resp.WriteHeader(201)
	}

	Handle("/", handler)

	if handlers["/"] == nil {
		t.Error("Handler should be registered")
	}
}

func TestHandleRequest(t *testing.T) {
	err := handleRequest(nil)
	if err.Error() != "Connection is nil" {
		t.Error("Error should be 'Connection is nil'")
	}

	c := mocks.NewMockConn()

	err = handleRequest(c)
	if !strings.HasPrefix(err.Error(), "Error parsing request:") {
		t.Error("Error should be 'Error parsing request:'")
	}

	if c.Closed != true {
		t.Error("Connection should be closed")
	}

	handlers["/"] = func(req *HttpRequest, resp ResponseWriter) {
		t.Log("Handler called")

		if req.Url != "/" {
			t.Error("Request URL should be '/'")
		}

		if req.Method != "GET" {
			t.Error("Request method should be 'GET'")
		}

		if req.Proto != "HTTP/1.1" {
			t.Error("Request proto should be 'HTTP/1.1'")
		}

		if resp.Header() == nil {
			t.Error("Response header should not be nil")
		}
	}

	c.ReadBuf = []byte("GET / HTTP/1.1\r\n\r\n")
	c.Closed = false

	err = handleRequest(c)

	c.ReadBuf = []byte("GET /notvalid HTTP/1.1\r\n\r\n")
	c.Closed = false

	err = handleRequest(c)
	if string(c.WriteBuf) != string([]byte(NotFoundError.Error())) {
		t.Error("Error should be 'Not Found'")
	}

	// after the request is handled properly, err should be nil
	if err != nil {
		t.Error("Error should be nil")
	}
}
