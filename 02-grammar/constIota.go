package main

import "fmt"

const (
	Running = 1 << iota
	Waiting
	Send
	Receive
)

func main() {

	stat := Running | Send
	display(stat)
}

func display(stat int) int {
	if stat & Running == Running {
		fmt.Println("Running")
	}
	if stat & Waiting == Waiting {
		fmt.Println("Waiting")
	}
	if stat & Send == Send {
		fmt.Println("Sending")
	}
	if stat & Receive == Receive{
		fmt.Println("Receiving")
	}
	return stat
}
