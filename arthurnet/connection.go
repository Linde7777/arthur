package arthurnet

import (
	"fmt"
	"main/arthurinterface"
	"net"
	"time"
)

type Connection struct {
	conn              *net.TCPConn
	connID            uint32
	IsClosed          bool
	notifyToCloseChan chan struct{}
	handleFunc        arthurinterface.HandleFunc
}

func (c *Connection) Start() {
	// 连接调用Start后，除了调用handlerFUnc，可能还会有其他操作，所以放在goroutine里
	go func() {
		err := c.handleFunc(c.conn)
		if err != nil {
			fmt.Println("handleFunc err: ", err)
			return
		}
	}()

	for {
		select {
		case <-c.notifyToCloseChan:
			return
		default:
			time.Sleep(2 * time.Second)
		}
	}
}

func (c *Connection) Stop() {
	if c.IsClosed {
		return
	}

	c.IsClosed = true
	c.notifyToCloseChan <- struct{}{}
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
