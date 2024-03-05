package arthurnet

import (
	"fmt"
	"main/arthurinterface"
	"net"
	"time"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	fmt.Printf("Server %s is starting, IPVersion: %s, IP: %s, Port: %d\n",
		s.Name, s.IPVersion, s.IP, s.Port)

	// TODO: goroutine泄露
	// Listen业务单独一个goroutine，后面还会有其他业务放在不同的goroutine
	go func() {
		listener, err := net.ListenTCP(s.IPVersion, &net.TCPAddr{
			IP:   net.ParseIP(s.IP),
			Port: s.Port,
		})
		if err != nil {
			fmt.Println("net.ListenTCP err: ", err)
			panic(err)
		}
		fmt.Println("listen ok")

		for {
			// 肯定得是按顺序建立实际连接，如果是
			//for{
			//	go func() {
			//		conn,err:=listener.AcceptTCP()
			//      do something
			//	}()
			//}
			// 那就是无限循环创建goroutine, 会让机器爆炸
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listener.AcceptTCP err: ", err)
				continue
			}

			// 这个用户建立了实际连接后，单独开一个goroutine处理他的业务，
			// 不要挡住其他用户建立实际连接.
			// 这里简单模拟一下业务，后面会替换.
			go func() {
				buf := make([]byte, 512)
				for {
					count, err := conn.Read(buf)
					if err != nil {
						fmt.Println("conn.Read err: ", err)
						continue
					}

					// 先简单模拟一下业务逻辑，做一个简单的回复功能
					fmt.Printf("Receive data: %s\n", string(buf[:count]))
					_, err = conn.Write([]byte("This is Arthur Morgan"))
					if err != nil {
						fmt.Println("conn.Write err: ", err)
						continue
					}
				}
			}()

		}
	}()

}

func (s *Server) Stop() {
	fmt.Println("stopping")
	// todo: 释放掉服务资源，如goroutine
}

func (s *Server) Serve() {
	s.Start()

	//todo: 这里可以做一些启动服务后的额外业务
	for {
		//阻塞，不然主goroutine退出了，服务就退出了
		time.Sleep(20 * time.Second)
	}
}

func NewServer(name, ipVersion, ip string, port int) arthurinterface.IServer {
	return &Server{
		Name:      name,
		IPVersion: ipVersion,
		IP:        ip,
		Port:      port,
	}
}
