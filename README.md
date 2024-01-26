# DBHandlers

This package provides a set of utilities to facilitate executing SQL queries and scanning results into structs. It uses Go's generic features to provide a flexible and type-safe way to work with database rows.

## Functions

- **ExecuteQuery**: Execute SQL queries with parameters and scan the results into slices of a specified struct type.

## Prerequisites

- Go 1.18 or higher
- An SQL database driver compatible with Go's `database/sql` package (e.g., MySQL, PostgreSQL).

## Installation

Run this command in your project directory:

```sh
go get github.com/ChrisBolis99/DBHandlers
```

Remember to then import it in your .go file

```go
import "github.com/ChrisBolis99/DBHandlers"
