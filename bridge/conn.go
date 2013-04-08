package bridge

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

type Conn struct {
	Address  string
	response *http.Response
}

var (
	_      io.ReadWriteCloser = &Conn{}
	client *http.Client
)

func init() {
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
	transport := http.DefaultTransport
	client = &http.Client{
		Transport: transport,
	}
}

func (c *Conn) Close() error {
	return c.response.Body.Close()
}

func (c *Conn) Read(p []byte) (n int, err error) {
	for c.response == nil {
		<-time.After(time.Microsecond)
	}
	return c.response.Body.Read(p)
}

func (c *Conn) Write(p []byte) (n int, err error) {
	c.response, err = client.Post(c.Address, "application/json", bytes.NewReader(p))
	return len(p), err
}
