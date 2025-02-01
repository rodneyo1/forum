package database // Global setup before tests
import (
	"fmt"
	"os"
	"testing"
)

const dbFile = "test.db"

// Setup function to delete the database file before each test
func deleteTestDb() {
	// Delete the database file if it exists
	if _, err := os.Stat(dbFile); err == nil {
		err := os.Remove(dbFile)
		if err != nil {
			fmt.Printf("Failed to delete the database file: %v\n", err)
		}
	}
}


func TestMain(m *testing.M) {
	// Setup before tests
	fmt.Println("Global setup: init db handle")
	deleteTestDb()
	Init(dbFile)
	defer db.Close()

	// Run tests
	code := m.Run()

	// Teardown after tests
	fmt.Println("Global teardown: after tests")
	deleteTestDb()
	os.Exit(code) // exit with the code from m.Run()
}
