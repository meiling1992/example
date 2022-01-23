package main

import (
	inf "example/ex01/InfiniteChannel/infinitechannel"
	"fmt"

	"time"
)

// 	"example/ex01/InfiniteChannel/api/infinitechannel"

func main() {
	// t.Log("Starting InfiniteChannel")
	cinstance := inf.NewInfiniteChannel()
	go func() {
		for i := 1; i < 20; i++ {
			cinstance.In(i)
			time.Sleep(time.Second)
		}
		cinstance.Close()
	}()
	for i := 1; i < 50; i++ {
		val := cinstance.Out()
		fmt.Printf("%v", val)
		// t.Log()
		time.Sleep(time.Microsecond)

	}
}

// package infinitechannel

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func TestInfiniteChannel(t *testing.T) {
// 	t.Log("Starting InfiniteChannel")
// 	cinstance := NewInfiniteChannel()
// 	go func() {
// 		for i := 1; i < 20; i++ {
// 			cinstance.In(i)
// 			time.Sleep(time.Second)
// 		}
// 		cinstance.Close()
// 	}()
// 	for i := 1; i < 50; i++ {
// 		val := cinstance.Out()
// 		t.Log(fmt.Sprintf("rece.Out val:%v", val))
// 		time.Sleep(time.Microsecond)

// 	}
// }
