package ticketBooking

import (
	"fmt"
	"strings"
)

func _validateUserDetails (user *Person) (bool, bool, bool) {

	firstNameError := false
	lastNameError := false
	emailError := false

	if len(user.firstName) < 2 {
		fmt.Printf("First Name has length less than 3")
		firstNameError = true
	}

	if len(user.lastName) < 2 {
		fmt.Printf("Last Name has length less than 3")
		lastNameError = true
	}

	if !strings.Contains(user.email, "@") {
		fmt.Printf("Last Name has length less than 3")
		emailError = true
	}

	return firstNameError, lastNameError, emailError

}

func _getUserDetails () Person {

	user := Person{}

	fmt.Printf("Please enter your first name\n")
	fmt.Scan(&user.firstName)

	fmt.Printf("Please enter your last name\n")
	fmt.Scan(&user.lastName)

	fmt.Printf("Please enter your email\n")
	fmt.Scan(&user.email)

	return user

}

func _bookTickets() func(uint) (bool, uint) {

	var remainingTickets uint = 50

	return func (needToBook uint) (bool, uint) {

		tooManyTicketsError := false

		if needToBook > remainingTickets {
			tooManyTicketsError = true
		} else {
			remainingTickets -= needToBook
		}

		return tooManyTicketsError, remainingTickets

	}
}