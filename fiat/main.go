package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Minute * 40)
	for _ = range ticker.C {
		fmt.Printf("ticked at %v", time.Now())
		data()
	}
}
