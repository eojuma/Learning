package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	waitGrpup sync.WaitGroup
	mut       sync.Mutex
)
var signals []string

func main() {
	// for i := 0; i < 4; i++ {
	// 	 greeter("Hello")
	// 	 greeter("World")
	// }

	websitelist := []string{
		"https://youtube.com",
		"https://fb.com",
		"https://github.com",
		"https://safaricom.com",
		"https://go.dev",
	}

	for _, web := range websitelist {
		waitGrpup.Add(1)
		go getStatusCode(web)
	}
	waitGrpup.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	fmt.Println(s)
// }

func getStatusCode(endpoint string) {
	defer waitGrpup.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops error", err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Println(res.StatusCode, endpoint)
}


