package cowww

import (
	"github.com/tekeoglan/cowww/mocks"
	"testing"
)

func TestResponseHeader(t *testing.T) {

	conn := mocks.NewMockConn()
	r := &response{
		c:             conn,
		handlerHeader: Header{},
	}

	r.Header().Set("Content-Type", "text/plain")
	r.Header().Set("Content-Length", "4")

	if r.handlerHeader["Content-Type"] != "text/plain" {
		t.Error("Handler header should contain 'text/plain'")
	}

	if r.handlerHeader["Content-Length"] != "4" {
		t.Error("Handler header should contain '4'")
	}
}

func TestResponseWrite(t *testing.T) {

	conn := mocks.NewMockConn()
	r := &response{
		c:             conn,
		handlerHeader: Header{"Content-Type": "text/plain"},
	}
	r.Write([]byte("test"))

	expected := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 4\r\n\r\ntest"

	if string(conn.WriteBuf) != expected {
		t.Errorf("\nExpected: '%s'\nGot: '%s'", expected, string(conn.WriteBuf))
	}
}

func TestResponseWriteHeader(t *testing.T) {
	conn := mocks.NewMockConn()
	r := &response{
		c: conn,
	}

	r.WriteHeader(StatusOk)

	if r.handlerHeader == nil {
		t.Error("Handler header should not be nil")
	}

	if r.statusCode != StatusOk {
		t.Errorf("Expected status code %d, got %d", StatusOk, r.statusCode)
	}

	if r.status != statusText(StatusOk) {
		t.Errorf("Expected status %s, got %s", statusText(StatusOk), r.status)
	}
}
