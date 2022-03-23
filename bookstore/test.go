package main

import (
	"fmt"
	"sync"
	"time"
)

var a int
var b int
func main() {
	var muA sync.Mutex
	var muB sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("==start==A")
		muA.Lock()
		time.Sleep(time.Second*1)
		muB.Lock()
		defer func() {
			muA.Unlock()
			muB.Unlock()
			wg.Done()
			fmt.Println("==end==A")
		}()
	}()

	go func() {
		fmt.Println("==start==B")
		muB.Lock()
		time.Sleep(time.Second*1)
		muA.Lock()
		defer func() {
			muA.Unlock()
			muB.Unlock()
			wg.Done()
			fmt.Println("==end==B")
		}()
	}()

	fmt.Println("==wait==")
	wg.Wait()
	fmt.Println("==success==")
}