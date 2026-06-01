package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"divar_cookie_harvestor/utils"
	"divar_cookie_harvestor/interface"
)

func main() {
	var phone_num string

	fmt.Println("Enter the phone number: ")
	fmt.Scan(&phone_num)

	utils.Confirm_phone_number(phone_num)

	// sample for print a message with colors
	// utils.Print("test message with color red", utils.Red)

	// making request
	req := interface.Make_phone_number_request(phone_num)
	data, err := json.Marshal(req)
	if err != nil {
		utils.Print("failed to marshal request", utils.Red)
		return
	}

	// sending request
	resp, err := http.Post(utils.API_SENDING_PHONE_NUMBER, "application/json", bytes.NewBuffer(data))
	if err != nil {
		utils.Print("failed to send request", utils.Red)
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	// reading response


}
