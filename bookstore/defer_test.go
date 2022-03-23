package main

import "testing"

func sum(x int) int  {
	m := 0
	for i:=0;i<x;i++ {
		m += i
	}
	return  m
}

func sumWithDefer()  {
	defer func() {
		sum(10)
	}()
}

func sumWithOutDefer()  {
	sum(10)
}

func BenchmarkBenchWithDefer(b *testing.B)  {
	sumWithDefer()
}

func BenchmarkBenchWithOutDefer(b *testing.B)  {
	sumWithOutDefer()
}
