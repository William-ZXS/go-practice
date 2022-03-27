package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	uuid, _ := uuid.NewUUID()
	fmt.Println("==uuid==", uuid)
}
