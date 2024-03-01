package mocks

import (
	"net"
	"time"
)

type mockConn struct {
	writeBuf []byte
}

func NewMockConn() *mockConn {
	return &mockConn{writeBuf: []byte{}}
}

func (c *mockConn) Read(b []byte) (int, error) {
	return 0, nil
}

func (c *mockConn) Write(b []byte) (int, error) {
	c.writeBuf = append(c.writeBuf, b...)
	return len(b), nil
}

func (c *mockConn) Close() error {
	return nil
}

func (c *mockConn) LocalAddr() net.Addr {
	return nil
}

func (c *mockConn) RemoteAddr() net.Addr {
	return nil
}

func (c *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func (c *mockConn) WrittenBuf() []byte {
	return c.writeBuf
}
