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

	fmt.Printf("Welcome to our %v booking application\n", conferenceName) // %v to print any variable
	fmt.Printf("We have total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		if userTickets < remainingTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{} // create an empty slice
			for _, booking := range bookings {
				var names = strings.Fields(booking) // Fields returns slices elements when whitespace found. names is a slice here
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// noTicketsRemaining := remainingTickets == 0 // can be written like this
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.") // end of programe
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining. So, you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}
}
