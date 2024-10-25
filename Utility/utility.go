package Utility

//This is a different package.
// We have to name the package as we named the main package

import (
	"fmt"
)

// This is a function which is available to all the files in the package
// The way of exporting a function is by making the first letter of the function name as capital
// Similarly we can export a variable by making the first letter of the variable as capital
func GeneralGreetings(name string) {
	fmt.Println("Welcome to the ticket booking system")
}

func GreetUser(name string, availableTickets int) {
	fmt.Printf("Hello %v  get ready to book the tickets \n", name)
	fmt.Printf("There are %v tickets available now\n", availableTickets)
}

func GetUser() (string, string, int, int) {
	var fname string
	var lname string
	var ageOfUser int
	var noOfTicketRequired int
	fmt.Println("Please enter your first name")
	fmt.Scan(&fname)
	fmt.Println("Please enter last name")
	fmt.Scan(&lname)
	fmt.Println("Please enter age")
	fmt.Scan(&ageOfUser)
	fmt.Println("Please enter number of tickets required")
	fmt.Scan(&noOfTicketRequired)
	return fname, lname, ageOfUser, noOfTicketRequired
}
