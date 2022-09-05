package main

import (
	"fmt"
	"go-tutorial/helper"
)

// var totalTickets uint = 50
var remainingTickets uint = 50
var names = []string{}
var emails = []string{}
var userTickets = []uint{}

// type areaInterface interface {
// 	Tarea() float64
// 	Tperimeter() float64
// }

// type shape struct {
// 	length  int
// 	breadth int
// }

// func (a shape) Tarea() float64 {
// 	return float64(a.length * a.breadth)
// }

// func (a shape) Tperimeter() float64 {
// 	return float64(a.length + a.breadth)
// }

func main() {

	greetUser()

	for true {
		name, email, ticket := acceptInputFromUser()

		isValidName, isValidEmail, isValidTickets := helper.ValidInputFields(name, email, ticket, remainingTickets)

		if isValidTickets {

			if isValidName && isValidEmail {

				storeData(name, email, ticket)
			} else {

				invalidNameEmailMsg()
				continue
			}
		} else {

			fmt.Printf("%v tickets can not be booked, %v tickets are remaining!\n", ticket, remainingTickets)
			continue
		}

		remainingTickets -= ticket
		showConfirmBooking(ticket)

		if remainingTickets == 0 {
			break
		}
	}

	showBookedTicketsUserName()

	// // fmt.Printf("%v check \n", testFunc(1))
	// var area areaInterface
	// area = shape{10, 20}

	// fmt.Printf("Area of rectangle : %v\n", area.Tarea())
	// fmt.Printf("Perimeter of rectangle : %v\n", area.Tperimeter())
}

func greetUser() {

	fmt.Println("Welcome to our conference booking portal!")
	fmt.Println("Please enter following details.")
}

func acceptInputFromUser() (string, string, uint) {

	fmt.Printf("Enter your name : ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Enter your email : ")
	var email string
	fmt.Scan(&email)

	fmt.Printf("Enter number of booking you want : ")
	var ticket uint
	fmt.Scan(&ticket)

	return name, email, ticket
}

func storeData(name string, email string, ticket uint) {

	userTickets = append(userTickets, ticket)
	emails = append(emails, email)
	names = append(names, name)
}

func invalidNameEmailMsg() {

	fmt.Printf("Entered name or email is not valid, please enter valid name or email, %v tickets remaining hurry up!\n", remainingTickets)
}

func showConfirmBooking(ticket uint) {

	fmt.Printf("You have booked %v tickets of our conference\n", ticket)
}

func showBookedTicketsUserName() {
	fmt.Printf("Total bookings done by %v\n", names)
}

// func testFunc(a int) *int {
// 	var aa int = a
// 	return &aa
// }
