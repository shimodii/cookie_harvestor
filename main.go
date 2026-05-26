package main

import (
	"fmt"
)

func confirm_phone_number(phone_num string) {
	var choice string

	fmt.Println("Phone number is: ", phone_num, ", procceed? (y|n)")
	fmt.Scan(&choice)

	if choice == "y" {
		fmt.Println("ok test")
	} else if choice == "n" {
		fmt.Println("Exitting...")
		return
	} else {
		fmt.Println("Please enter y or n")
		return
	}
}

func main() {
	var phone_num string

	fmt.Println("Enter the phone number: ")
	fmt.Scan(&phone_num)

	confirm_phone_number(phone_num)

}
