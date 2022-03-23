package main

import (
	"fmt"
	"github.com/William-ZXS/workpool"
	"time"
)

func main() {
	p := workpool.New()
	time.Sleep(time.Second*2)
	for i:=0;i<10;i++ {
		p.Schedule(
			func() {
				fmt.Println("==doing work==")
				time.Sleep(time.Second*2)
				fmt.Println("==finish work==")
			},
		)
	}

	time.Sleep(time.Second*5)
	fmt.Println("==p begin stop==")
	p.Stop()
	fmt.Println("==p stoped==")
}

