package main

import (
	"fmt"
	"github.com/William-ZXS/go-socket/frame"
	"github.com/William-ZXS/go-socket/packet"
	"net"
	"time"
)

//todo any类型怎么处理

//socket server
//解析处理数据
//一个packet有多少数据
func handleConn(conn net.Conn) {
	defer func() {
		//todo any 怎么处理
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		conn.Close()
	}()
	//循环处理链接上的数据
	for {
		conn.SetDeadline(time.Now().Add(time.Minute * 1))
		// read
		fmt.Println("==decode==")
		payload, err := frame.Decode(conn)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 处理payload得到返回值
		fmt.Println("==Handle==")
		ackPayload, err := packet.Handle(payload)
		if err != nil {
			return
		}
		//Encode返回值
		fmt.Println("==Encode==")
		frame.Encode(conn, ackPayload)
	}

}

func main() {
	address := ":9090"
	network := "tcp"
	l, err := net.Listen(network, address)
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}
	fmt.Printf("%s listen %s success\n", network, address)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			return
		}

		fmt.Printf("receive conn from %s \n", conn.RemoteAddr())
		go handleConn(conn)
	}
}
