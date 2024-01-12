package main

import "fmt"

func main() {
	// variable declaration
	conferenceName := "GO Conference" 
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// array
	//var bookings = [50]string{"zihad", "hello", "golang"} // string bookings[50] = {"zihad", "hello", "golang"}
	var bookings [50]string // string booking[50]

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

	remainingTickets -= userTickets;
	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings);
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Println("Array lenth: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}