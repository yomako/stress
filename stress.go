package main

import (
	"time"
)

func stress() {
	for {
		for start := time.Now(); time.Since(start) < 10*time.Second; {
		}
		time.Sleep(20 * time.Second)
	}
}
