package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SequentialSum(inputSize int) int {
	start := time.Now()
	sum := 0
	for i := 1; i <= inputSize; i++ {
		sum += process(i)
	}
	elapsed := time.Since(start)
	fmt.Printf("Sequential time taken: %v\n", elapsed)
	return sum
}

// ParallelSum implement this method.
func ParallelSumChan(inputSize int) int {
	start := time.Now()
	ch := make(chan int)
	var mu sync.Mutex
	var sum int = 0
	for i := 1; i <= inputSize; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- process(i)
		}(i)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			sum += <-ch
		}()
	}
	wg.Wait()
	close(ch)

	fmt.Printf("Parallel sum using Channel time taken: %v\n", time.Since(start))
	return sum
}

func ParallelSum(inputSize int) int {
	var mu sync.Mutex
	start := time.Now()
	var sum int = 0
	for i := 1; i <= inputSize; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			temp := process(i)

			mu.Lock()
			sum += temp
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	// wg.Wait()
	fmt.Printf("Parallel time taken: %v\n", time.Since(start))
	return sum
}

func process(num int) int {
	time.Sleep(time.Millisecond) // simulate processing time
	return num * num
}

func main() {
	in := 10000

	var wgg sync.WaitGroup

	fmt.Println("*************************Calling these methods sequentially*************************")
	seq_sum := SequentialSum(in)
	fmt.Printf("Sequential sum: %v\n", seq_sum)

	parallel_sum := ParallelSum(in)
	fmt.Printf("Parallel sum: %v\n", parallel_sum)

	para_chan := ParallelSumChan(in)
	fmt.Printf("Parallel sum using channel: %v\n", para_chan)

	fmt.Println("*************************Calling these methods concurrently*************************")

	wgg.Add(1)
	go func() {
		defer wgg.Done()
		seq_sum := SequentialSum(in)
		fmt.Printf("Sequential sum: %v\n", seq_sum)
	}()

	wgg.Add(1)
	go func() {
		defer wgg.Done()
		parallel_sum := ParallelSum(in)
		fmt.Printf("Parallel sum: %v\n", parallel_sum)
	}()

	wgg.Add(1)
	go func() {
		defer wgg.Done()
		para_chan := ParallelSumChan(in)
		fmt.Printf("Parallel sum using channel: %v\n", para_chan)
	}()

	wgg.Wait()
}

// *************************Calling these methods sequentially*************************
// Sequential time taken: 11.627910927s
// Sequential sum: 333383335000
// Parallel time taken: 9.866007ms
// Parallel sum: 333383335000
// Parallel sum using Channel time taken: 15.754297ms
// Parallel sum using channel: 333383335000
// *************************Calling these methods concurrently*************************
// Parallel sum using Channel time taken: 20.820779ms
// Parallel sum using channel: 333383335000
// Parallel time taken: 20.81495ms
// Parallel sum: 333383335000
// Sequential time taken: 11.612824545s
// Sequential sum: 333383335000
