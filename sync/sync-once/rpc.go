package rpc

import (
	"net"
	"sync"
)

var twonce = &sync.Once{}
var lock sync.Mutex
var client *Client

// Client is a singleton to interact with a ficticious rpc layer
type Client struct {
	conn net.Conn
}

// NewClientOnce returns a new client or nil in case of any underlying error
// using sync.Once
func NewClientOnce() *Client {
	onceFn := func() {
		// expensive call with possible error goes here
		c, err := getConn()
		if err != nil {
			twonce = new(sync.Once)
			return
		}
		client = &Client{
			conn: c,
		}
	}

	twonce.Do(onceFn)
	return client
}

// NewClientLock returns a new client or nil in case of any underlying error
// using sync.Mutex
func NewClientLock() *Client {
	lock.Lock()
	defer lock.Unlock()
	if client == nil {
		c, err := getConn()
		if err != nil {
			return nil
		}
		client = &Client{
			conn: c,
		}
		return client
	}
	return client
}

// please don't do this, ever
func getConn() (net.Conn, error) {
	conn := &net.TCPConn{}
	return conn, nil
}
