package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestVerifyUserValidCredentials(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
		AddRow(1, "testuser", "test@example.com", "$2a$10$1234567890123456789012")

	mock.ExpectQuery("SELECT id, username, email, password FROM users WHERE email = ? OR username = ?").
		WithArgs("test@example.com", "test@example.com").
		WillReturnRows(rows)

	result := VerifyUser("test@example.com", "correctpassword")

	if !result {
		t.Errorf("Expected VerifyUser to return true for valid credentials, but got false")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
