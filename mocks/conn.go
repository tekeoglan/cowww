package mocks

import (
	"net"
	"time"
)

type mockConn struct {
	ReadBuf  []byte
	WriteBuf []byte
	Closed   bool
}

func NewMockConn() *mockConn {
	return &mockConn{WriteBuf: []byte{}}
}

func (c *mockConn) Read(b []byte) (int, error) {
	copy(b, c.ReadBuf)
	return len(c.ReadBuf), nil
}

func (c *mockConn) Write(b []byte) (int, error) {
	c.WriteBuf = append(c.WriteBuf, b...)
	return len(b), nil
}

func (c *mockConn) Close() error {
	c.Closed = true
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
