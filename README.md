# DBHandlers

This package provides a set of utilities for various database handling tasks, including executing SQL queries, scanning results into structs, and working with DBML (Database Markup Language) through its `DBMLParser` module. It leverages Go's generic features to offer flexible and type-safe database interaction and schema manipulation.

## Modules and Functions

### General

- **ExecuteQuery**: Execute SQL queries with parameters and scan the results into slices of a specified struct type.

### DBMLParser

- **ParseDBML**: Parses a string containing DBML content into a `Schema` struct, representing the database schema defined within. This function enables the easy manipulation and analysis of database schemas directly

- **GenerateSQLFromSchema**: Takes a `Schema` struct as input and generates SQL statements for creating the database schema as defined in the struct

## Prerequisites

- Go 1.18 or higher
- A SQL database driver compatible with Go's `database/sql` package (e.g., MySQL, PostgreSQL).

## Installation

To install `DBHandlers`, run the following command in your project directory:

```sh
go get github.com/ChrisBolis99/DBHandlers
