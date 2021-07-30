package main

import (
	"fmt"
	"time"
)

func cpuIntensive(p *int) {
	for i := 1; i <= 100000; i++ {
		*p = i
		//runtime.Gosched()
	}
	fmt.Println("Done intensive thing")
}

func printVar(p *int) {
	fmt.Printf("x = %d\n", *p)
}

func main() {
	x := 0
	go cpuIntensive(&x)     // This should go into background
	go printVar(&x)         // This won't get scheduled until everything has finished.
	time.Sleep(time.Second) // Wait for goroutines to finish
}
