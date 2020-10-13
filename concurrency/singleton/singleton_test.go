package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
	
}
var singleInstance *Singleton
var once sync.Once

func GetSingleton() *Singleton {
	once.Do(func() {
		fmt.Println("create object")
		singleInstance=new(Singleton)
	})
	return singleInstance
}
func TestGetSingleton(t *testing.T)  {
	var wg sync.WaitGroup
	for i:=1;i<10;i++{
		wg.Add(1)
		go func() {
			obj:=GetSingleton()
			fmt.Printf("%x\n",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
