package bridge

import (
	"git.300brand.com/coverage/config"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func Client() *rpc.Client {
	conn := &Conn{Address: config.RPC.Address}
	return jsonrpc.NewClient(conn)
}
