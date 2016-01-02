package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

// sql 패키지에 등록될
// PostgresDriver 타입
type PostgresDriver struct{}

// 데이터베이스에 대한 연결을 수행하는 Open 함수.
func (dr PostgresDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *PostgresDriver

// main 함수에 앞서 호출될 init 함수
func init() {
	d = new(PostgresDriver)
	sql.Register("postgres", d)
}
