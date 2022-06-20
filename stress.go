package main

import (
	"time"
)

func main() {
	for {
		for start := time.Now(); time.Since(start) < 10*time.Second; {
		}
		time.Sleep(20 * time.Second)
	}
}
