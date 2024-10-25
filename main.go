// https://pkg.go.dev/fmt#pkg-overview

// in go everything is organizated in packages
// every go file starts with a package declaration
// main package is special, it's used for creating an executable file
// when we create our own application, we also need to include it in a package

package main

//fmt is a package which provides support for the Print function
//Th
import (
	"booking-app/Utility"
	"fmt"
	"strings"
)

// Go needs to know the entry point of our appication
// We need explicitly tell go that this is the entry point of our application, start executing from here

// entry point of our application is a named function
// For one go prgram we will have 1 main function

//--------------------------------------------------//

// when a variable is created and assigned with value in go the type is automatically inferred (bit like JavaScript)
// However if  a variable is created and not assigned with any value, we need to explicitly mention the type of the variable
// Unlike JavaScript where we can create a variable without assigning any value and it will be undefined.

func main() {
	//fmt.Print("Hello world")
	//var name = "John" // string
	//name := "Jon" // string same as the above line ; this is a short hand declaration
	// We can not use short hand declaration for constants
	const age = 25 // integer
	//name = "Doe"
	var totoalQuota = 100
	var usedQuota = 52
	var conferenceName = "GopherCon"
	var conferenceLocation = "San Francisco"
	fmt.Println("We have a total of ", totoalQuota, "and", usedQuota, "is already been booked")
	//With Printf we can use the reference of the variable with %{letter} in between the DOUBLE QUOTES.
	//We however need to pass the variable as the second and following arguments to the Printf function
	// \n is used for new line
	fmt.Printf("We are organizing %s in %s \n", conferenceName, conferenceLocation)

	var availableQuota int
	availableQuota = totoalQuota
	fmt.Printf("%v tickets are avaiable right now.\n", availableQuota)

	// TYPE OF VARIABLES can be checked using %T
	fmt.Printf("The type of availableQuota is %T\n", availableQuota)

	// We can also use the Scan function to take input from the user
	// We need to declare a variable and pass the reference of the variable to the Scan function
	// The scan function will take in pointers
	// The pointer is a variable which points to the memory location
	// of another variable which points to the memory address of the value . Refer pointer pic in assets folder
	// pointers in go are also called special variable
	// This will help store the value entered by the user in the variable

	// fmt.Println("Please enter your first name")
	// fmt.Scan(&fname)
	// fmt.Println("Please enter last name")
	// fmt.Scan(&lname)
	// fmt.Println("Please enter age")
	// fmt.Scan(&ageOfUser)
	// fmt.Println("Please enter number of tickets required")
	// fmt.Scan(&noOfTicketRequired)

	//availableQuota = availableQuota - noOfTicketRequired

	// %s is used for string %d is used for integer %v is used for any type of variable
	// fmt.Printf("Hello %s %s, you are %d years old \n", fname, lname, ageOfUser)
	// fmt.Printf("You have successfully booked %d tickets. \n", noOfTicketRequired)
	// fmt.Printf("Total number of tickets available for booking are %d\n", availableQuota)

	//Array declaration in go
	// Array size has to be defined first and then the values can be assigned
	// Much like Java and unlike JavaScript where we can directly assign values to an array
	// var booking [10]string
	// booking[0] = "John"

	// fmt.Printf("The first booking is %s\n", booking[0])
	// fmt.Printf("The first booking is %s\n", booking)
	// fmt.Printf("The type of booking array %T\n", booking)
	// fmt.Printf("The length of booking array %v\n", len(booking))

	//When we do not provide the size of the array, it is called a slice
	//Adding values to a slice has to be done using the append function

	var bookingList []string
	// bookingList = append(bookingList, "John")
	// bookingList = append(bookingList, "Ram")
	// bookingList = append(bookingList, "Abhi")
	// bookingList = append(bookingList, "Shyam")

	// fmt.Printf("The first booking is %v\n", bookingList[0])
	// fmt.Printf("Booking list is  %v\n", bookingList)
	// fmt.Printf("Booking list length is %d\n", len(bookingList))

	// Infinite loop

	for {

		if len(bookingList) < 3 {
			fname, lname, ageOfUser, noOfTicketRequired := Utility.GetUser()
			fullname, isAgeValid, address := checkAndReturnUserDetails(fname, lname, ageOfUser, conferenceLocation)
			fmt.Printf("The full name is %v\n", fullname)
			fmt.Printf("The age is valid %v\n", isAgeValid)
			fmt.Printf("The address %v\n", address)
			if isAgeValid {
				availableQuota = availableQuota - noOfTicketRequired

				fmt.Printf("Hello %s %s, you are %d years old \n", fname, lname, ageOfUser)
				fmt.Printf("You have successfully booked %d tickets. \n", noOfTicketRequired)
				fmt.Printf("Total number of tickets available for booking are %d\n", availableQuota)

				bookingList = append(bookingList, fname+" "+lname)

				fmt.Printf("Booking list is  %v\n", bookingList)
			} else {
				fmt.Println("The age is not valid, user not allowed to book the tickets")
			}

			noOfTicketRequired = 0
		} else {
			break
		}

		checkBooking(bookingList)
	}

	// this is a finite for loop
	// it gives the index of the list in index field and the value in the booking field
	// much like for each loop in JS
	for index, booking := range bookingList {
		fmt.Printf("Booking %d is %v\n", index, booking)
		// Fields function is used to split the string into words based on space
		var name = strings.Fields(booking)
		fmt.Printf("First name is %v\n", name[0])
	}

}

func checkBooking(bookings []string) {
	// connect to the database
	// insert the bookings into the database
	// return the status of the booking
	fmt.Println("Connecting to the database")
	fmt.Println("Inserting the bookings into the database")
	fmt.Println("Returning the status of the booking")

	for index, booking := range bookings {
		var flag string
		fmt.Printf("Booking %d is %v\n", index, booking)
		if len(booking) > 3 && len(booking) < 10 {
			flag = "short"
		} else if len(booking) >= 10 && len(booking) < 15 {
			flag = "medium"
		} else {
			flag = "long"
		}

		switch flag {
		case "short":
			fmt.Println("The booking is short")
		case "medium":
			fmt.Println("The booking is medium")
		case "long":
			fmt.Println("The booking is long")
		default:
			fmt.Println("The booking is not valid")

		}

	}
}

// go allows a function to return multiple values
// The type of values returned needs to be mentioned in the function signature
// The values returned can be accessed by the calling function and assigned to variables
func checkAndReturnUserDetails(fname string, lname string, ageOfUser int, location string) (string, bool, string) {
	fullname := fname + " " + lname
	isAgeValid := ageOfUser > 20
	address := location
	return fullname, isAgeValid, address
}
