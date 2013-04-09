package bridge

// Shamelessly ripped from net/rpc/json and modified slightly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/rpc"
	"sync"
)

type clientCodec struct {
	url string

	// temporary work space
	req  clientRequest
	resp clientResponse

	// JSON-RPC responses include the request id but not the request method.
	// Package rpc expects both.
	// We save the request method in pending when sending a request
	// and then look it up by request ID when filling out the rpc Response.
	mutex   sync.Mutex   // protects pending
	pending chan pending // map request id to method name
}

type pending struct {
	response *http.Response
	method   string
}

var _ rpc.ClientCodec = &clientCodec{}
var client *http.Client

func init() {
	transport := http.DefaultTransport
	client = &http.Client{
		Transport: transport,
	}
}

// NewClient returns a new rpc.Client to handle requests to the
// set of services at the other end of the connection.
func NewClient(url string) *rpc.Client {
	return rpc.NewClientWithCodec(NewClientCodec(url))
}

// NewClientCodec returns a new rpc.ClientCodec using JSON-RPC on conn.
func NewClientCodec(url string) rpc.ClientCodec {
	return &clientCodec{
		url:     url,
		pending: make(chan pending, 1),
	}
}

type clientRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      uint64      `json:"id"`
	Jsonrpc string      `json:"jsonrpc"`
}

func (c *clientCodec) WriteRequest(r *rpc.Request, param interface{}) error {
	c.req.Method = r.ServiceMethod
	c.req.Params = param
	c.req.Id = r.Seq
	c.req.Jsonrpc = "2.0"

	body, err := json.Marshal(&c.req)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}

	c.pending <- pending{resp, r.ServiceMethod}

	return nil
}

type clientResponse struct {
	Id     uint64           `json:"id"`
	Result *json.RawMessage `json:"result"`
	Error  interface{}      `json:"error"`
}

func (r *clientResponse) reset() {
	r.Id = 0
	r.Result = nil
	r.Error = nil
}

func (c *clientCodec) ReadResponseHeader(r *rpc.Response) error {
	p := <-c.pending

	dec := json.NewDecoder(p.response.Body)
	defer p.response.Body.Close()

	c.resp.reset()
	if err := dec.Decode(&c.resp); err != nil {
		return err
	}

	r.Error = ""
	r.Seq = c.resp.Id
	r.ServiceMethod = p.method

	if c.resp.Error != nil {
		x, ok := c.resp.Error.(string)
		if !ok {
			return fmt.Errorf("invalid error %v", c.resp.Error)
		}
		if x == "" {
			x = "unspecified error"
		}
		r.Error = x
	}
	return nil
}

func (c *clientCodec) ReadResponseBody(x interface{}) error {
	if x == nil {
		return nil
	}
	return json.Unmarshal(*c.resp.Result, x)
}

func (c *clientCodec) Close() error {
	return nil
}
