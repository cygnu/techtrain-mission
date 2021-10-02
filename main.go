package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
}

func main() {
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/sample_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var user User

	err = db.QueryRow("SELECT * FROM users WHERE id = ?", 1).Scan(&user.ID, &user.Name)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("レコードが存在しません")
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(user.ID, user.Name)
	}
}
