package main

import (
	"fmt"
	"time"
)

func main() {
	const number = 10
	message := make(chan int, number)
	flag := make(chan bool)
	done := make(chan bool)
	consumer := func(msg chan int, done chan bool, flag chan bool) {
		go func() {
			Ticker := time.NewTicker(1 * time.Second)
			for range Ticker.C {
				select {
				case <-done:
					fmt.Println("child process exit!")
					flag <- false
					return
				default:
					if m, ok := <-msg; ok {
						INfo := fmt.Sprintf("send msg :%d", m)
						fmt.Println(INfo)
						if len(msg) == 0 {
							fmt.Println("Msg Send All End !!")
							flag <- true
						}
					}
				}
			}
		}()
	}
	consumer(message, done, flag)
	product := func(n int, message chan int) {
		for i := 1; i <= n+5; i++ {
			fmt.Println("Rece MSG(product)", i)
			message <- i
		}
	}

	product(number, message)

	if flagv, ok := <-flag; ok {
		fmt.Println("Ready Send_SIG Close done(channel)")
		close(done)
		fmt.Println("Send Close.done Rece Sig flagv:", flagv)
	}
	<-flag
	defer close(message)
	fmt.Println("Main Exit")

}
