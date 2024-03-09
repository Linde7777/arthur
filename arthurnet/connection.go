package arthurnet

import (
	"fmt"
	"main/arthurinterface"
	"net"
)

type Connection struct {
	conn              *net.TCPConn
	connID            uint32
	IsClosed          bool
	notifyToCloseChan chan struct{}
	handleFunc        arthurinterface.HandleFunc
}

func (c *Connection) Start() {
	data := []byte("this is a message from demoHandleFunc")
	err := c.handleFunc(c.conn, data)
	if err != nil {
		fmt.Println("handleFunc err: ", err)
		return
	}
}

func (c *Connection) Stop() {
	if c.IsClosed {
		return
	}

	c.IsClosed = true
	err := c.conn.Close()
	if err != nil {
		fmt.Println("conn.Close err: ", err)
		return
	}

}

func (c *Connection) GetConnID() uint32 {
	return c.connID
}

func NewConnection(conn *net.TCPConn, connID uint32, handleFunc arthurinterface.HandleFunc) *Connection {
	return &Connection{
		conn:              conn,
		connID:            connID,
		IsClosed:          false,
		notifyToCloseChan: make(chan struct{}),
		handleFunc:        handleFunc,
	}
}
