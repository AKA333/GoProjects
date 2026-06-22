package main

import (
	"fmt"
)

func main(){
	f()
	fmt.Println("returned normal from f")
}

func f(){
	defer func(){
		if r:= recover(); r!= nil{
			fmt.Println("recovered in f", r)
		}
	}()
	fmt.Println("Calling g")
	g(0)
	fmt.Println("returned normal from g")
}

func g(i int){
	if i>3 {
		fmt.Println("panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in g", i)
	fmt.Println("printing in g", i)
	g(i+1)
}

// Output:
// Calling g
// printing in g 0
// printing in g 1
// printing in g 2
// printing in g 3
// panicking
// defer in g 3
// defer in g 2
// defer in g 1
// defer in g 0
// recovered in f 4
// returned normal from f