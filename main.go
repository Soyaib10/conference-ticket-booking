package main

import (
	"fmt"
)

// package level variables: it does not support a := 5 type initialization. They are accessable from any function under the
// package they are. so the functions that uses these variables as parameter don't need this in parameter list.
const conferenceTickets int = 50

var conferenceName = "GO Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // empty slice of struct

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	// greeting user function
	greetingUsers()

	for {
		// getting user input
		firstName, lastName, email, userTickets := getUserInput()

		// user input validation
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// book tickets
			bookTicket(userTickets, firstName, lastName, email)
			// print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			// noTicketsRemaining := remainingTickets == 0 // can be written like this
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.") // end of programe
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetingUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName) // %v to print any variable
	fmt.Printf("We have total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string { // function takes input and returns slice
	firstNames := []string{} // create an empty slice
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// create a struct for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
