package ticketBooking

import (
	"fmt"
)

type Person struct {

	firstName, lastName, email string

	ticketsNeeded uint
}

var Bookings []Person

var TicketsRemaining uint

func Main_booking() {

	bookTickets := _bookTickets()

	TicketsRemaining = 50

	fmt.Printf("Welcome to GO CONFERENCE. Hurry Up ! Only %v tickets remain", TicketsRemaining)
	fmt.Println("Please provide details as asked below: ")

	customer := _getUserDetails()

	firstNameError, lastNameError, emailError := _validateUserDetails(&customer)

	if firstNameError || lastNameError || emailError {

		fmt.Println("\nPlease restart the booking process.")

	}
	
	fmt.Println("Please specify how many tickets do you need")
	fmt.Scan(&customer.ticketsNeeded)

	errorWhileBooking, ticketsRemainingNow := bookTickets(customer.ticketsNeeded)
	TicketsRemaining = ticketsRemainingNow

	if errorWhileBooking {
		fmt.Printf("Cannot book %v tickets as only %v are available", customer.ticketsNeeded, TicketsRemaining)
		fmt.Println("\nPlease re enter the details.")
		return
	}

	Bookings = append(Bookings, customer)
	fmt.Printf("%v tickets successfully booked. See you at the conference \n", customer.ticketsNeeded)

}




// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	conferenceName := "Go Conference"
// 	const conferenceTickets uint = 50
// 	var remainingTickets uint = conferenceTickets

// 	var bookings []string // is a slice. with size wud have been an array


// 	for {
		
// 		var firstName string
// 		var lastName string
// 		var email string
// 		var userTickets uint

// 		fmt.Printf("Welcome to %s ticket booking application. %v Tickets left.\n", conferenceName, remainingTickets)
// 		fmt.Printf("Please enter your login credentials.\n FIRST NAME: ")
// 		fmt.Scan(&firstName)

// 		fmt.Printf("Please enter your LAST NAME: ")
// 		fmt.Scan(&lastName)

// 		fmt.Printf("Please enter your EMAIL: ")
// 		fmt.Scan(&email)

// 		fmt.Printf("How many tickets do you want to book out of %v available: ", remainingTickets)
// 		fmt.Scan(&userTickets)

// 		if remainingTickets < userTickets {
// 			fmt.Printf("can't book %v tickets as only %v tickets are left.\nPlease try to book again", userTickets, remainingTickets)
// 			continue
// 		}

// 		remainingTickets -= userTickets

// 		// bookings[conferenceTickets - remainingTickets] = firstName + " " + lastName
// 		bookings = append(bookings, firstName+" "+lastName)

// 		firstNames := []string{}

// 		for _ , booking := range bookings {
// 			firstNames = append(firstNames, strings.Split(booking, " ")[0])
// 		}

// 		fmt.Printf("Dear %v, thanks for booking %v tickets. An email will be sent to %v \n", firstName, userTickets, email)
// 		fmt.Printf("User Interaction Ended, %v tickets left. First Names of bookings till now : %v \n", remainingTickets, firstNames)

// 		noTicketsLeft := remainingTickets <= 0

// 		if noTicketsLeft {
// 			break
// 		}

// 	}

// }
