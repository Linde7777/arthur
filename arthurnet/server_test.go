package arthurnet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServerBasic(t *testing.T) {
	s := NewServer("ArthurServer", "tcp4", "0.0.0.0", 1899)
	go s.Serve()

	fmt.Println("Sleep 3 seconds to wait server start")
	time.Sleep(time.Second * 3)
	go MockClient()
	for {
		// 防止主goroutine退出
		time.Sleep(5 * time.Second)
	}
}

func MockClient() {
	conn, err := net.Dial("tcp", "0.0.0.0:1899")
	if err != nil {
		fmt.Println("client dial err: ", err)
		return
	}

	for {
		_, err = conn.Write([]byte("I am client."))
		if err != nil {
			fmt.Println("conn write err: ", err)
			return
		}

		buf := make([]byte, 512)
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read err: ", err)
			return
		}
		fmt.Printf("Receive from server: %s\n", string(buf[:count]))
	}

}
