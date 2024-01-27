package DBHandlers

import (
	"database/sql"
	"fmt"
	"reflect"
)

// ExecuteQuery executes a SQL query with the provided parameters and scans the resulting rows into a slice of the specified prototype type.
//
// Parameters:
// - queryToCall: The SQL query string to be executed. This query can include placeholders for parameters, "?" for SQL databases.
// - params: A slice of interface{} that contains the parameters to be used with the queryToCall. The order of parameters in this slice should match the order of placeholders in the query string.
// - db: A pointer to an sql.DB object that represents an open database connection. This connection is used to prepare and execute the SQL query.
// - prototype: An instance of the type T that serves as a prototype for scanning each row of the result set. This prototype determines the structure into which each row's data will be scanned.
//
// Returns:
// - A slice of type T, where T is the type of the prototype parameter. Each element in the slice represents a row from the query result, with its fields populated according to the data in the row.
// - An error if the query preparation, execution, or row scanning fails. If the query and row scanning succeed, the error is nil.
func ExecuteQuery[T any](queryToCall string, params []interface{}, db *sql.DB, prototype T) ([]T, error) {
	statement, err := db.Prepare(queryToCall)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.Query(params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items, err := scanMultipleRows(rows, prototype)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// scanMultipleRows iterates over the rows returned from a SQL query and scans the data from each row into new instances of type T,
// T is a generic type parameter that must be a struct compatible with the structure of the row data.
//
// Parameters:
//   - rows: *sql.Rows representing the result set from a SQL query execution. It is assumed that the caller has already executed a SQL query
//     and that `rows` is ready for iteration.
//   - prototype: An instance of type T that serves as a prototype for creating new instances of T for each row in the result set.
//     The fields of T should correspond to the columns selected in the SQL query
//
// Returns:
// - A slice of type T, where each element is populated with data from a row in the result set.
// - An error if any occurs during row iteration or data scanning.
func scanMultipleRows[T any](rows *sql.Rows, prototype T) ([]T, error) {
	var results []T
	structType := reflect.TypeOf(prototype)
	if structType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("prototype must be a struct")
	}

	for rows.Next() {
		structValue := reflect.New(structType).Elem()

		fields := make([]interface{}, structValue.NumField())
		for i := range fields {
			fields[i] = structValue.Field(i).Addr().Interface()
		}

		if err := rows.Scan(fields...); err != nil {
			return nil, err
		}

		results = append(results, structValue.Interface().(T))
	}

	return results, nil
}
