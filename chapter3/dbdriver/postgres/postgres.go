package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

// PostgresDriver provides our implementation for the
// sql package.
type PostgresDriver struct{}

// Open provides a connection to the database.
func (dr PostgresDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *PostgresDriver

// init is called prior to main.
func init() {
	d = new(PostgresDriver)
	sql.Register("postgres", d)
}
