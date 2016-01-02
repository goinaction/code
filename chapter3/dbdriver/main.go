// sql 패키지와의 연동을 간단히 보여주기 위한
// 예제 프로그램
package main

import (
	"database/sql"

	_ "github.com/webgenie/go-in-action/chapter3/dbdriver/postgres"
)

// 애플리케이션 진입점
func main() {
	sql.Open("postgres", "mydb")
}
