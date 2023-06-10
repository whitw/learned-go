package main

import "fmt"

func bye(){
	fmt.Println("bye bye!")
}

func main(){
	fmt.Println("Hello World!")
	defer bye()
	for{
		
	}
}
