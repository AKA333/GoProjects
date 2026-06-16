package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	const totalCustomers int = 50
	const totalTickets int = 5

	ticketChan := make(chan int, totalTickets)

	for i := 1; i <= totalTickets; i++ {
		ticketChan <- i
	}

	for i := 1; i <= totalCustomers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case ticket := <-ticketChan:
				fmt.Printf("Customer %v, Booked ticket success %v\n", i, ticket)
			default:
				fmt.Printf("Customer %d: Sold out\n", i)
			}
		}()
	}

	wg.Wait()
	fmt.Println("All request processed")

}
