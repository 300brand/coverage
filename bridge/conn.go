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
	Response *http.Response
}

var (
	_      io.ReadWriteCloser = &Conn{}
	client *http.Client
)

func init() {
	log.SetFlags(log.Lmicroseconds)
	transport := http.DefaultTransport
	client = &http.Client{
		Transport: transport,
	}
}

func (c *Conn) Close() error {
	log.Println("Close()")
	return c.Response.Body.Close()
}

func (c *Conn) Read(p []byte) (n int, err error) {
	for c.Response == nil {
		<-time.After(time.Microsecond)
	}
	log.Println("Read()")
	log.Println("len(p):", len(p))
	log.Printf("%v", c.Response)
	return c.Response.Body.Read(p)
}

func (c *Conn) Write(p []byte) (n int, err error) {
	log.Println("Write()")
	log.Printf("%s", p)
	c.Response, err = client.Post(c.Address, "application/json", bytes.NewReader(p))
	log.Printf("%v", err)
	return len(p), err
}
