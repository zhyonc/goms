package listener

import (
	"goms/network/server/api"
	"net"
)

type Listener interface {
	Start()
	Stop()
}

type TCPHandleFunc func(conn net.Conn)

type UDPHandleFunc func(msg *api.Message)
