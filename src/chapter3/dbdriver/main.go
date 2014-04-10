package main

import (
	_ "chapter2/dbdriver/postgres" //<co id="anon-import" />
	"database/sql"
)

func main() {
	sql.Open("postgres", "mydb") //<co id="sql-open" />
}
