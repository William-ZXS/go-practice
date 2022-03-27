package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

//定义packet的结构体 拥有的方法

// 结构体
// length 4字节
// command 1字节
//

const (
	CommandConn   int8 = 10
	CommandSubmit int8 = 20
)

var ErrShortRead = errors.New("short read")
var ErrShortWrite = errors.New("short write")

type Packet struct {
	Command int8
	Uuid    string //36字节
	Body    []byte
}

func Decode(payload []byte) (*Packet, error) {
	command := payload[0]
	uuid := payload[1:37]
	body := payload[37:]
	return &Packet{
		Command: int8(command),
		Uuid:    string(uuid),
		Body:    body,
	}, nil
}

func Encode(packet *Packet) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, packet.Command)
	return bytes.Join([][]byte{buf.Bytes(), []byte(packet.Uuid), []byte(packet.Body)}, nil), nil
}

func Handle(payload []byte) ([]byte, error) {
	packet, err := Decode(payload)
	if err != nil {
		return nil, err
	}

	//do business
	ackPacket := &Packet{
		Command: packet.Command,
		Uuid:    packet.Uuid,
		Body:    []byte(fmt.Sprintf("command %d success", packet.Command)),
	}

	ackPayload, err := Encode(ackPacket)
	if err != nil {
		return nil, err
	}
	return ackPayload, nil
}
