package database

import (
	"fmt"
	"testing"
)

func TestVerifyUser_ValidCredentials(t *testing.T) {
	result := VerifyUser("milton@mail.com", "mPass")

	if !result {
		t.Errorf("Expected VerifyUser to return true for valid credentials, but got false")
	}
}

func TestGetUserByMailOrName(t *testing.T) {
	var err error

	user, err := GetUserByEmailOrUsername("milton@mail.com", "milton")
	if user.Username != "milton" || err != nil {
		fmt.Println("USername: ", user.Username, " Passcode: ", user.Password, " Email: ", user.Email)
		t.Errorf("expected username to be milton, but got %s\n", user.Username)
	}
}
