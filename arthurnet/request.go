package arthurnet

import "main/arthurinterface"

type Request struct {
	conn arthurinterface.IConnection
}

func (r *Request) GetConn() arthurinterface.IConnection {
	return r.conn
}
