package frame

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var ErrShortRead = errors.New("short read")
var ErrShortWrite = errors.New("short write")

// length 4字节
func Decode(r io.Reader) ([]byte, error) {
	var length int32
	err := binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	payload := make([]byte, length-4)
	n, err := io.ReadFull(r, payload)
	if err != nil {
		return nil, err
	}
	if int32(n) != length-4 {
		return nil, ErrShortRead
	}

	return payload, nil
}

func Encode(w io.Writer, payload []byte) error {
	var length int32 = int32(len(payload) + 4)
	err := binary.Write(w, binary.BigEndian, &length)
	if err != nil {
		return err
	}
	n, err := w.Write(payload)
	if err != nil {
		return err
	}
	if n != len(payload) {
		return ErrShortWrite
	}
	fmt.Println("write success")
	return nil
}
