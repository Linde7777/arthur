package arthurinterface

import (
	"net"
)

type IConnection interface {
	Start()
	Stop()
	GetConnID() uint32
}

// HandleFunc 是用户的实际业务，和链接绑定
type HandleFunc func(conn *net.TCPConn, data []byte, cnt int) error
