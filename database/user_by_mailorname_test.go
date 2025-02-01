package database

import (
	"fmt"
	"os"
	"testing"
)

func TestGetUserByMailOrName(t *testing.T) {
	var err error

	_, err = CreateUser("milton", "milton@mail.com", "mPass")
	if err != nil {
		t.Errorf("Unable to save user")
		os.Exit(1)
	}
	user, err := GetUserByEmailOrUsername("milton@mail.com", "milton")
	if user.Username != "milton" || err != nil {
		fmt.Println("USername: ", user.Username, " Passcode: ", user.Password, " Email: ", user.Email)
		t.Errorf("expected username to be milton, but got %s\n", user.Username)
	}
}
