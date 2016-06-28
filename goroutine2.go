package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create Goroutines")

	go func() {
		printPrime("A")
		wg.Done()
	}()

	go func() {
		printPrime("B")
		wg.Done()
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers.
func printPrime(prefix string) {
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}
