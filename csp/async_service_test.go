package csp

import (
	"fmt"
	"testing"
	"time"
)

func service1() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask1() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Task is done.")
}

func AsyncService1() chan string {
	retCh := make(chan string, 1)
	//retCh := make(chan string, 1)
	go func() {
		ret := service1()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}
func TestSelect(t *testing.T)  {
	select {
	case ret1:=<-AsyncService1():
		t.Log(ret1)
	case <-time.After(time.Millisecond*100):
		t.Error("time out")
	}


}