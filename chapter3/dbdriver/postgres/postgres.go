//<start id="init"/>
package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

func init() {
	d = new(PostgresDriver) //<co id="driver-create" />
	sql.Register("postgres", d)
}

//<end id="init"/>

type PostgresDriver struct{}

func (dr PostgresDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *PostgresDriver
