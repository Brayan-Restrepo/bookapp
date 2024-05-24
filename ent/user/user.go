// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeBooks holds the string denoting the books edge name in mutations.
	EdgeBooks = "books"
	// Table holds the table name of the user in the database.
	Table = "users"
	// BooksTable is the table that holds the books relation/edge.
	BooksTable = "books"
	// BooksInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	BooksInverseTable = "books"
	// BooksColumn is the table column denoting the books relation/edge.
	BooksColumn = "user_books"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByBooksCount orders the results by books count.
func ByBooksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBooksStep(), opts...)
	}
}

// ByBooks orders the results by books terms.
func ByBooks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBooksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBooksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BooksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BooksTable, BooksColumn),
	)
}
