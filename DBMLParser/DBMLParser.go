package DBMLParser

import (
	"fmt"
	"strings"
)

type Column struct {
	Name string
	Type string
}

type Table struct {
	Name    string
	Columns []Column
}

type Schema struct {
	Tables []Table
}

func ParseDBML(dbmlContent string) (Schema, error) {
	var schema Schema
	var currentTable *Table

	lines := strings.Split(dbmlContent, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(trimmedLine, "Table"):
			tableName := strings.Fields(trimmedLine)[1]
			currentTable = &Table{Name: tableName}
			schema.Tables = append(schema.Tables, *currentTable)

		case strings.Contains(trimmedLine, ":"):
			parts := strings.Split(trimmedLine, ":")
			if currentTable != nil {
				columnName := strings.TrimSpace(parts[0])
				columnType := strings.TrimSpace(parts[1])
				column := Column{Name: columnName, Type: columnType}
				currentTable.Columns = append(currentTable.Columns, column)
			}

		case strings.TrimSpace(trimmedLine) == "}":
			currentTable = nil
		}
	}

	return schema, nil
}

func generateSQLForTable(table Table) string {
	var columnDefinitions []string

	for _, column := range table.Columns {
		columnDefinitions = append(columnDefinitions, fmt.Sprintf("%s %s", column.Name, column.Type))
	}

	return fmt.Sprintf("CREATE TABLE %s (\n  %s\n);", table.Name, strings.Join(columnDefinitions, ",\n  "))
}

func GenerateSQLFromSchema(schema Schema) string {
	var sqlStatements []string

	for _, table := range schema.Tables {
		sqlStatements = append(sqlStatements, generateSQLForTable(table))
	}

	return strings.Join(sqlStatements, "\n\n")
}
