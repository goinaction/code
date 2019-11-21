// Sample program to show how to show you how to briefly work
// with the sql package.
package main

import (
	"database/sql"
//空白符 让init运行 还不会因为包未使用产生错误
	_ "github.com/goinaction/code/chapter3/dbdriver/postgres"
)

// main is the entry point for the application.
func main() {
	sql.Open("postgres", "mydb")
}
