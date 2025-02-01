package database

import (
	"fmt"
	"strings"
)

// Create inserts data into the specified table using provided field names and values.
func Create(tableName string, columns []string, values []interface{}) (int64, error) {
	var identifier int64 = -1
	var err error

	// the number of columns must match the number of values
	if len(columns) != len(values) {
		return identifier, fmt.Errorf("number of columns (%d) does not match number of values (%d)", len(columns), len(values))
	}

	// prepare the SQL query
	// using the "?" ensures that we avoid sql injection
	placeholders := make([]string, len(values))
	for i := range values {
		placeholders[i] = "?"
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	// execute the query
	result, err := db.Exec(query, values...)
	if err != nil {
		return identifier, err
	}

	// get the last inserted ID
	identifier, err = result.LastInsertId()
	return identifier, err
}
