package main

import "fmt"

func main() {
	// variable declaration
	var conferenceName string = "GO Conference" 
	const conferenceTickets int = 50
	remainingTickets := 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName) // %v to print any variable
	fmt.Printf("We have total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}