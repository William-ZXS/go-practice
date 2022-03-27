package packet

import (
	"bytes"
	"github.com/google/uuid"
	"testing"
)

func TestHandle(t *testing.T) {
	//command := int8(20)
	uuid := uuid.New().String()
	payload := bytes.Join([][]byte{[]byte{20}, []byte(uuid), []byte("hello")}, nil)
	ackPayload, err := Handle(payload)
	if err != nil {
		t.Errorf("want nil actual %s", err.Error())
	}
	if int8(ackPayload[0]) != int8(20) {
		t.Errorf("want 20 actual %d", int8(ackPayload[0]))
	}
	if string(ackPayload[1:37]) != uuid {
		t.Errorf("want %s actual %s", uuid, string(ackPayload[1:37]))
	}

}
