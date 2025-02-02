package utils

import (
	"fmt"
	"os"
	"regexp"
)

// PasswordStrength checks password strength and length constraints
func PasswordStrength(password string) string {
	// Ensure password does not exceed 72 bytes
	if len([]byte(password)) > 72 {
		fmt.Println("Your password exceeds the maximum length of 72 bytes")
		os.Exit(1)
	}

	// Ensure password is at least 8 characters long
	if len(password) < 8 {
		fmt.Println("Password is too short: It should have at least 8 characters.")
		os.Exit(1)
	}

	// Check for at least one uppercase letter/lowercase/one digit/one special character
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	// Validate all conditions
	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		fmt.Println("Password too weak: It should have at least one uppercase letter, one lowercase letter, one number, and one special character.")
		os.Exit(1)
	}

	return password
}