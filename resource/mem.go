package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		b := make([]byte, 1e7*i)
		fmt.Println(len(b))
		time.Sleep(time.Second * time.Duration(1))
	}
	time.Sleep(time.Second * time.Duration(1000))
}
