package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var firstName, lastName, emailAddress string
	var userTickets uint
	bookings := []string{}

	fmt.Printf("Welcome to %s booking App\n", conferenceName)
	fmt.Printf("We have %d avilable out off %d\n", remainingTickets, conferenceTickets)
	for {
		fmt.Printf("Please enter First Name : \n")
		fmt.Scan(&firstName)
		if firstName == "quit" {
			break
		}
		fmt.Printf("Please enter Last Name : \n")
		fmt.Scan(&lastName)
		fmt.Printf("Please enter Email Address : \n")
		fmt.Scan(&emailAddress)
		fmt.Printf("Please enter No.Of.Tickets You need : \n")
		fmt.Scan(&userTickets)
		tmptickets := remainingTickets - userTickets
		if remainingTickets > userTickets && tmptickets >= 1 {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+""+lastName+" - "+strconv.FormatUint(uint64(userTickets), 10))
			fmt.Printf("Hello %v Sucessfully  booked %v tickets.!! Happy Holidays \n", firstName, userTickets)
		} else {
			fmt.Printf("Unable to book your %v tickets please try with some number <= %v\n", userTickets, tmptickets)
		}
	}
	for _, booking := range bookings {
		fmt.Printf("All Booking in this session %v\n", strings.Fields(booking)[1])
	}

}
