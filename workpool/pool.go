package workpool

import (
	"fmt"
	"sync"
)

//不要通过共享内存来通信，而是通过通信来共享内存

type Task func()

type Pool struct {
	//容量 worker数量
	capacity int
	//是否停止标志
	quit   chan struct{}
	task   chan Task
	active chan struct{}
	wg sync.WaitGroup
}

//创建pool
func New() *Pool {
	cap := 5
	p := &Pool{
		capacity: cap,
		quit:     make(chan struct{}),
		task:     make(chan Task),
		active: make(chan struct{},cap),
	}
	go p.run()
	return p
}

//创建 worker
func (p *Pool) newWorker(id int)  {
	//监听 quit 和task
	go func() {
		p.wg.Add(1)
		fmt.Println("==worker==",id,"==start==")
		defer func() {
			if err:= recover();err!=nil{
				fmt.Println("==worker err==")
			}
			p.wg.Done()
			<-p.active
			fmt.Println("==worker==",id,"==stop==")
		}()

		for  {
			select {
			case <-p.quit:
				fmt.Println("==receive stop==")
				return
			case t := <-p.task:
				fmt.Println("==receive task==")
				t()
			}
		}
	}()

}


//run
func (p *Pool) run() {
	//初始化worker
	id := len(p.active)
	for {
		select {
		case <-p.quit:
			fmt.Println("==stop run==")
			return
		case p.active<-struct{}{}:
			id++
			p.newWorker(id)
		}
	}
}

//schedule
func (p *Pool) Schedule(t Task)  {
	select {
	case p.task<-t:
		fmt.Println("==scheduled==")
	}
}


//stop
func (p *Pool)Stop()  {
	close(p.quit)
	p.wg.Wait()
	fmt.Println("==stoped==")
}

