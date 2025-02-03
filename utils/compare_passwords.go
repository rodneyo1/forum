package utils

// Helper function to compare hashed passwords
func ComparePasswords(hashedPassword, inputPassword string) bool {
	// implement your password comparison logic here
	return hashedPassword == inputPassword // Replace with secure password comparison
}
