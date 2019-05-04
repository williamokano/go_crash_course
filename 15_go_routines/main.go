package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/williamokano/go_crash_course/15_go_routines/fib"
)

func main() {
	atoi, _ := strconv.Atoi(os.Args[1])
	poolSize := int(atoi)

	jobs := make(chan uint, poolSize)
	results := make(chan *big.Int, poolSize)

	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < poolSize; i++ {
		jobs <- uint(i)
	}
	close(jobs)

	for j := 0; j < poolSize; j++ {
		fmt.Println((<-results).String())
		fmt.Println()
	}
	close(results)
}

func worker(jobs <-chan uint, results chan<- *big.Int) {
	for n := range jobs {
		results <- fib.Fib(n)
	}
}
