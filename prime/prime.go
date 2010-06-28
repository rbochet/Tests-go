package main

import (
	"fmt"
)

// Generate all the numbers and add them to the global channel
func generate(allChan chan int) {
	for i := 2; ; i++ {
		allChan <- i
	}
}

func main() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}


// Take a chan, and put the relative primes in another chan
func filter(inCh, outCh chan int, prime int) {
	for {
		i := <-inCh
		if i%prime != 0 {
			outCh <- i
		}
	}
}
