package main

import "fmt"

func main() {
	// variable declaration
	conferenceName := "GO Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// slices
	//var bookings = []string{"zihad", "hello", "golang"} // vector<string> bookings = {"zihad", "hello", "golang"}
	// or
	//bookings := []string{}
	var bookings []string // vector<string> bookings

	fmt.Printf("Welcome to our %v booking application\n", conferenceName) // %v to print any variable
	fmt.Printf("We have total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// getting user input
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName) // bookings.push_back("fkjf")
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice lenth: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
