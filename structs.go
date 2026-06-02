package main

import (
	"errors"
	"fmt"
	"time"
)



func main(){
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user 

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	//...do something awsesome with that gathered data!
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}



func getUserData(promptText string) string{
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}