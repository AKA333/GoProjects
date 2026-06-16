// Design a booking system where we have exactly 10 available tickets. Instead of using a Mutex to protect a variable, use a channel to represent the tickets.
// 100 customers will try to book at once. A customer 'gets a ticket' by successfully receiving from the channel. If the channel is empty, they should immediately get a 'Sold Out' message. You must ensure the program only exits once all 100 attempts are processed.
 
// Input: totalCustomers = 100, ticketCount = 10
// Expected Output:
// - 10 "Success" messages
// - 90 "Sold out" messages
// - Final "All attempts processed" message (Ensures no deadlocks)

package main

import (
	"fmt"
	"sync"
)
 var wg sync.WaitGroup

func main(){
	const totalCustomers int = 50
	const totalTickets int = 5

	ticketChan := make (chan int, totalTickets)

	for i:= 1; i<=totalTickets; i++ {
		ticketChan <- i
	}

	for i:= 1; i<=totalCustomers; i++ {
		wg.Add(1)
		go func () {
			defer wg.Done()
			select {
			case ticket:= <-ticketChan:
				fmt.Printf("Customer %v, Booked ticket success %v\n", i, ticket)
			default:
				fmt.Printf("Customer %d: Sold out\n", i)
			}
		}()
	}

	wg.Wait()
	fmt.Println("All request processed")

}