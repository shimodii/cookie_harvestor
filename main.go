package main

import (
	"fmt"

	"divar_cookie_harvestor/utils"
)

func main() {
	var phone_num string

	fmt.Println("Enter the phone number: ")
	fmt.Scan(&phone_num)

	utils.Confirm_phone_number(phone_num)

	// sample for print a message with colors
	// utils.Print("test message with color red", utils.Red)

}
