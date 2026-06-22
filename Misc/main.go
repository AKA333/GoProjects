package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex

var temp int = 10

var keySync chan int = make(chan int)

func main() {
	fmt.Println("Hello, World!")

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("go routine")
	}()
	// wg.Wait()

	fmt.Println("main")

	wg.Add(1)
	go update(10)

	wg.Add(1)
	go update(20)

	wg.Wait()
}

func update(val int) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	temp += val
	keySync <- temp
	fmt.Printf("Updated value temp: %v\n", <-keySync)
}
