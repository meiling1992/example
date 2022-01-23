package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//https://api.github.com/users/meiling1992
	requestfuture := func(url string) <-chan []byte {
		cr := make(chan []byte, 1)
		go func() {
			var body []byte
			defer func() {
				cr <- body
			}()
			res, err := http.Get(url)
			if err != nil {
				return
			}
			defer res.Body.Close()
			body, _ = ioutil.ReadAll(res.Body)
		}()
		return cr
	}

	future := requestfuture("https://api.github.com/users/meiling1992")
	if body, ok := <-future; ok {
		Msg := fmt.Sprintf("response len(%d)", len(body))
		fmt.Println(Msg)
	}
	fmt.Println("Main future exit")

}
