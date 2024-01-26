# DBHandlers

The `DBHandlers` package provides a set of utilities to facilitate executing SQL queries and scanning results into structs. It uses Go's generic features to provide a flexible and type-safe way to work with database rows.

## Features

- **ExecuteQuery**: Execute SQL queries with parameters and scan the results into slices of a specified struct type.

## Prerequisites

- Go 1.18 or higher (due to the use of generics).
- An SQL database driver compatible with Go's `database/sql` package (e.g., MySQL, PostgreSQL).

## Installation

Run the following command in your project directory:

```sh
go get github.com/ChrisBolis99/DBHandlers
