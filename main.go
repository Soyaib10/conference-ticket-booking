package main

import (
	"fmt"
	"strings"
)

func main() {
	// variable declaration
	conferenceName := "GO Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	// greeting user function
	greetingUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		// getting user input
		firstName, lastName, email, userTickets := getUserInput()

		// user input validation
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// book tickets
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			// print first names
			firstNames := getFirstNames(bookings)
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

func greetingUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName) // %v to print any variable
	fmt.Printf("We have total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string { // function takes input and returns slice
	firstNames := []string{} // create an empty slice
	for _, booking := range bookings {
		var names = strings.Fields(booking) // Fields returns slices elements when whitespace found. names is a slice here
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber // return multiple values
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

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
