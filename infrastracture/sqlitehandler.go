package infrastructure

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // driver

	"github.com/CanobbioE/web-service-example/interfaces/repositories"
)

// SqliteHandler wraps a connection to a SQLite database.
type SqliteHandler struct {
	Conn *sql.DB
}

// Execute performs the given statement on the database
// the handler is connected to.
func (sh *SqliteHandler) Execute(statement string) error {
	_, err := sh.Conn.Exec(statement)
	return err
}

// Query performs the given statement on the database
// the handler is connected to and returns the result.
func (sh *SqliteHandler) Query(statement string) (repositories.Row, error) {
	var result SqliteRow
	rows, err := sh.Conn.Query(statement)
	if err != nil {
		return result, err
	}

	result.Rows = rows
	return result, nil
}

// SqliteRow wraps sql results
// and implements our interfaces/repositories.Row
// to allow the interfaces to comunicate with the database
// without handling low level details.
type SqliteRow struct {
	Rows *sql.Rows
}

// Scan copies the columns from the matched row
// into the values pointed at by dest.
// See the documentation on Rows.Scan for details.
// If more than one row matches the query, Scan uses the first row
// and discards the rest.
// If no row matches the query, Scan returns ErrNoRows.
func (r SqliteRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest)
}

// Next prepares the next result row for reading with the Scan method.
// It returns true on success, or false if there is no next result row
// or an error happened while preparing it.
// Err should be consulted to distinguish between the two cases.
func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}

// NewSqliteHandler istannciate a new connection to the sqlite db.
// A filepath to the database file name must be supplied.
func NewSqliteHandler(fileName string) (*SqliteHandler, error) {
	conn, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}
	return &SqliteHandler{conn}, nil
}
