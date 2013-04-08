package bridge

import (
	"git.300brand.com/coverage/config"
	"net/rpc"
)

func New() *rpc.Client {
	return NewClient(&Conn{Address: config.RPC.Address})
}
