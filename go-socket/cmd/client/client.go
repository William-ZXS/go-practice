package main

import (
	"fmt"
	"github.com/William-ZXS/go-socket/frame"
	"github.com/William-ZXS/go-socket/packet"
	"github.com/google/uuid"
	"net"
	"time"
)

func main() {

	address := "127.0.0.1:9090"
	network := "tcp"
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("conn error: ", err)
		return
	}
	defer func() {
		conn.Close()
	}()

	conn.SetDeadline(time.Now().Add(time.Second * 5))

	p := &packet.Packet{
		Command: 20,
		Uuid:    uuid.New().String(),
		Body:    []byte("request1"),
	}
	payload, err := packet.Encode(p)
	if err != nil {
		return
	}
	//write
	frame.Encode(conn, payload)
	fmt.Println("send request", p)
	//read
	ackPayload, err := frame.Decode(conn)
	if err != nil {
		return
	}
	ackP, err := packet.Decode(ackPayload)
	if err != nil {
		return
	}
	fmt.Println("receive ack", ackP)

}
