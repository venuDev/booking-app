package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	type UserData struct {
		firstName       string
		lastName        string
		email           string
		numberOfTickets uint
	}

	var bookings = make([]UserData, 0)
	var tuser UserData
	fmt.Printf("Welcome to %s booking App\n", conferenceName)
	for {
		fmt.Printf("Please enter First Name : \n")
		fmt.Scan(&tuser.firstName)
		fmt.Printf("Please enter Last Name : \n")
		fmt.Scan(&tuser.lastName)
		fmt.Printf("Please enter Email Address : \n")
		fmt.Scan(&tuser.email)
		fmt.Printf("Please enter No.Of.Tickets You need : \n")
		fmt.Scan(&tuser.numberOfTickets)
		tmptickets := remainingTickets - tuser.numberOfTickets
		if remainingTickets > tuser.numberOfTickets && tmptickets >= 1 {
			remainingTickets = tmptickets
			bookings = append(bookings, tuser)
			fmt.Printf("Hello %v Sucessfully  booked %v tickets.!! Happy Holidays \n", tuser.firstName, tuser.numberOfTickets)
		} else {
			fmt.Printf("Unable to book your %v tickets please try with some number <= %v\n", tuser.numberOfTickets, tmptickets)
		}
		var quit string
		fmt.Println("Do you want continue Any key/No(N)")
		fmt.Scan(&quit)
		if (strings.EqualFold(quit, "N") ||strings.EqualFold(quit, "NO")) {
			break
		}
	}
	fmt.Printf("All Booking in this session\n Name      Tickets \n")
	for _, booking := range bookings {
		fmt.Printf("%v - %v \n", booking.email, booking.numberOfTickets)
	}

}
