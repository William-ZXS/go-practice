package frame

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestEncode(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	payload := []byte("hello")
	err := Encode(buf, payload)
	if err != nil {
		t.Errorf("Encode err %s", err.Error())
	}

	var totalLen int32
	err = binary.Read(buf, binary.BigEndian, &totalLen)
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}
	if totalLen != 9 {
		t.Errorf("want 9, actual %d", totalLen)
	}

	if buf.String() != "hello" {
		t.Errorf("want hello actual %s", buf)
	}

}

func TestDecode(t *testing.T) {
	data := []byte{0x0, 0x0, 0x0, 0x9, 'h', 'e', 'l', 'l', 'o'}

	payload, err := Decode(bytes.NewReader(data))
	if err != nil {
		t.Errorf("want nil actual %s", err.Error())
	}
	if string(payload) != "hello" {
		t.Errorf("want hello actual %s", string(payload))
	}
}
