package utils

import (
	"fmt"
)

func Confirm_phone_number(phone_number string) {
	var choice string

	fmt.Println("Phone number is: ", phone_number)
	Print("Proceed? (y|n)", Yellow)

	fmt.Scan(&choice)

	if choice == "y" || choice == "Y" {
		Print("OK OK", Green)
	} else if choice == "n" || choice == "N" {
		Print("Exitting...", Red)
		return
	} else {
		Print("Please insert y or n!, Exitting...", Red)
		return
	}

}
