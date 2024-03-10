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
	router            arthurinterface.IRouter
}

func (c *Connection) Start() {
	req := &Request{
		conn: c,
	}

	// 连接调用Start后，除了调用路由里面的方法，可能还会有其他操作，所以放在goroutine里
	go func(req arthurinterface.IRequest) {
		err := c.router.PreHandle(req)
		if err != nil {
			fmt.Println("PreHandle err: ", err)
			return
		}

		err = c.router.Handle(req)
		if err != nil {
			fmt.Println("Handle err: ", err)
			return
		}

		err = c.router.PostHandle(req)
		if err != nil {
			fmt.Println("PostHandle err: ", err)
			return
		}
	}(req)

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

func (c *Connection) GetRouter() arthurinterface.IRouter {
	return c.router
}

func NewConnection(conn *net.TCPConn, connID uint32, router arthurinterface.IRouter) *Connection {
	return &Connection{
		conn:              conn,
		connID:            connID,
		IsClosed:          false,
		notifyToCloseChan: make(chan struct{}),
		router:            router,
	}
}
