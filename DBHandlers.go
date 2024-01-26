package dbhandlers

import (
	"database/sql"
	"reflect"
)

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

func scanMultipleRows[T any](rows *sql.Rows, prototype T) ([]T, error) {
	var results []T
	structType := reflect.TypeOf(prototype)

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
