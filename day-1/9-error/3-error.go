package main

import (
	"errors"
	"fmt"
	"regexp"
)

// ValidateEmail checks if a given email is valid or not.
func ValidateEmail(email string) (bool, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return false, errors.New("invalid email address")
	}
	return true, nil
}

func main() {
	email := "test@example.com" // replace with the email you want to validate
	isValid, err := ValidateEmail(email)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !isValid {
		fmt.Println(email, "is not valid.")
		return
	}
	fmt.Println(email, "is valid.")

}

/*
Question 1:
Todo:
- Create a function ValidateEmail(email string) that checks if a given email is valid or not.
- This function should take in an email in the string format and return a boolean value and an error.
- The function should match the email against a regular expression to check if it is valid.
- If the email is valid, return true, otherwise, return an error with message "Invalid email address".
- Use the errors.New function to create the error message.
- In the main function, call the ValidateEmail function with an email and check the error.
- If there is an error, print the error message, otherwise print that the email is valid.
- Create a seperate package to write all the email validation functionality

Hint: The regular expression for an email address is: `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

/*


re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
if !re.MatchString(email) {}
*/
