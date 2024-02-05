package DBMLParser

import (
	"fmt"
	"strings"
)

type Column struct {
	Name        string
	Type        string
	NotNull     bool
	PrimaryKey  bool
	Unique      bool
	Constraints string
	Default     string
}

type Table struct {
	Name    string
	Columns []Column
}

type Schema struct {
	Tables []Table
}

// ParseDBML parses a string containing DBML (Database Markup Language) content and constructs a Schema struct representing the database schema defined within the DBML.
// The function processes the DBML content line by line, identifying table definitions, column definitions, and the end of table definitions to construct the schema
//
// Parameters:
// - dbmlContent: A string containing the DBML content to be parsed
//
// Returns:
// - Schema: A struct representing the parsed database schema. The Schema struct contains a slice of Table structs, each of which includes the table name and a slice of Column structs representing the table's columns
// - error: An error value that will be nil if the parsing succeeds without issues
func ParseDBML(dbmlContent string) (Schema, error) {
	var schema Schema
	lines := strings.Split(dbmlContent, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(trimmedLine, "Table"):
			handleTableDefinition(&schema, trimmedLine)

		case strings.Contains(trimmedLine, ":"):
			handleColumnDefinition(&schema, trimmedLine)

		case strings.TrimSpace(trimmedLine) == "}":
			endCurrentTable(&schema)
		}
	}

	return schema, nil
}

// GenerateSQLFromSchema takes a Schema struct as input and generates SQL statements for creating tables as defined in the schema. Each table in the schema is translated into a CREATE TABLE statement,
// with columns, types, and constraints (if any) properly formatted as part of the SQL statement
//
// Parameters:
// - schema: A Schema struct that represents the database schema. This struct should have been previously filled by parsing DBML content, containing all necessary information about tables, columns, and their types and constraints.
//
// Returns:
// - string: A single string containing the SQL statements for creating all tables defined in the schema, separated by two newlines (\n\n) for readability
func GenerateSQLFromSchema(schema Schema) string {
	var sqlStatements []string

	for _, table := range schema.Tables {
		sqlStatements = append(sqlStatements, generateSQLForTable(table))
	}

	return strings.Join(sqlStatements, "\n\n")
}

func generateSQLForTable(table Table) string {
	var columnDefinitions []string

	for _, column := range table.Columns {
		specificColumnDefinition := fmt.Sprintf("%s %s", column.Name, column.Type)

		if column.NotNull {
			specificColumnDefinition += " NOT NULL"
		}

		if column.PrimaryKey {
			specificColumnDefinition += " PRIMARY KEY"
		}

		if column.Unique {
			specificColumnDefinition += " UNIQUE"
		}

		if column.Default != "" {
			specificColumnDefinition += " DEFAULT " + column.Default
		}

		if column.Constraints != "" {
			specificColumnDefinition += " " + column.Constraints
		}

		columnDefinitions = append(columnDefinitions, specificColumnDefinition)
	}

	return fmt.Sprintf("CREATE TABLE %s (\n  %s\n);", table.Name, strings.Join(columnDefinitions, ",\n  "))
}

func handleTableDefinition(schema *Schema, line string) {
	tableName := strings.Fields(line)[1]
	currentTable := Table{Name: tableName}
	schema.Tables = append(schema.Tables, currentTable)
}

func handleColumnDefinition(schema *Schema, line string) {
	parts := strings.SplitN(line, ":", 2)
	columnName := strings.TrimSpace(parts[0])
	columnDef := strings.TrimSpace(parts[1])

	currentTable := &schema.Tables[len(schema.Tables)-1]

	column := parseColumnDetails(columnName, columnDef)
	currentTable.Columns = append(currentTable.Columns, column)
}

func parseColumnDetails(columnName, columnDef string) Column {
	var column Column
	column.Name = columnName

	parts := strings.SplitN(columnDef, " ", 2)
	column.Type = parts[0]

	if len(parts) > 1 {
		constraintsStr := strings.Trim(parts[1], "[]")
		constraints := strings.Split(constraintsStr, ",")

		for _, c := range constraints {
			constraint := strings.TrimSpace(c)
			switch {
			case strings.EqualFold(constraint, "pk") || strings.EqualFold(constraint, "primaryKey"):
				column.PrimaryKey = true
				column.NotNull = true

			case strings.EqualFold(constraint, "notNull"):
				column.NotNull = true

			case strings.EqualFold(constraint, "unique"):
				column.Unique = true

			case strings.HasPrefix(strings.ToLower(constraint), "default"):
				column.Default = strings.TrimPrefix(constraint, "default ")
				column.Default = strings.TrimSpace(column.Default)
			}
		}
	}

	return column
}

func endCurrentTable(schema *Schema) {
	// No action is needed to 'end' a table.
	// This function exists for symmetry and future use.
}
