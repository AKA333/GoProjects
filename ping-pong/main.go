package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	oddChan := make(chan int)
	evenChan := make(chan int)

	var limit int = 10

	wg.Add(2)
	go func(){
		defer wg.Done()
		for val := range oddChan {
			if(val>limit){
				close(evenChan)
				return
			} else {
				fmt.Println("odd:", val)
				evenChan <- val+1
			}
		}
	}()
	
	go func(){
		defer wg.Done()
		for val := range evenChan {
			if(val > limit){
				close(oddChan)
				return
			} else {
				fmt.Println("even:", val)
				oddChan <- val+1
			}
		}
	}()

	oddChan <-1

	wg.Wait()
}