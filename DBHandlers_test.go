package DBHandlers

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestExecuteQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	type Dummy struct {
		ID   int
		Name string
	}

	query := "SELECT id, name FROM dummy WHERE name LIKE '%Doe%'"
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "John Doe").
		AddRow(2, "Jane Doe")

	mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectQuery().WillReturnRows(rows)

	var dummyPrototype Dummy
	results, err := ExecuteQuery(query, nil, db, dummyPrototype)
	if err != nil {
		t.Errorf("error was not expected while executing query: %s", err)
	}

	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
