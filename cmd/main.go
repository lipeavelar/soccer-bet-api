package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		fmt.Println("Hello, World!")
		time.Sleep(10 * time.Second)
	}
}
