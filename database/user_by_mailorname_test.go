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

func TestVerifyUser_InvalidCredentials(t *testing.T) {
	result := VerifyUser("milton@mail.com", "password")

	if result {
		t.Errorf("Expected VerifyUser to return false for valid credentials, but got true")
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

func TestVerifyUser_NonExistentUser(t *testing.T) {
	result := VerifyUser("nonexistent@mail.com", "password")

	if result {
		t.Errorf("Expected VerifyUser to return false for non-existent user, but got true")
	}
}

func TestVerifyUser_NonExistentUsername(t *testing.T) {
	result := VerifyUser("nonexistent@mail.com", "password")

	if result {
		t.Errorf("Expected VerifyUser to return false for non-existent username, but got true")
	}
}

func TestVerifyUser_CaseSensitiveEmail(t *testing.T) {
	// Test with a lowercase email
	resultLower := VerifyUser("milton@mail.com", "mPass")
	if !resultLower {
		t.Errorf("Expected VerifyUser to return true for lowercase email, but got false")
	}

	// Test with an uppercase email
	resultUpper := VerifyUser("MILTON@MAIL.COM", "mPass")
	if resultUpper {
		t.Errorf("Expected VerifyUser to return false for uppercase email, but got true")
	}
}

func TestVerifyUser_CaseSensitiveUsername(t *testing.T) {
	// Test with correct username but different case
	result := VerifyUser("MILTON@mail.com", "mPass")

	if result {
		t.Errorf("Expected VerifyUser to return false for case-sensitive username, but got true")
	}

	// Test with correct username and correct case
	result = VerifyUser("milton@mail.com", "mPass")

	if !result {
		t.Errorf("Expected VerifyUser to return true for correct case-sensitive username, but got false")
	}
}
