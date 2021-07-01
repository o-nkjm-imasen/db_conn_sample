// work well

// on windows
// 1. go get github.com/go-sql-driver/mysql

package main

import (
	"database/sql"
	"log"
	_  "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql",  "imasen:densanka1@tcp(192.168.19.149)/fuga")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db)
	defer db.Close()

	rows, err := db.Query("SELECT name FROM hoges")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			return
		}
		fmt.Printf("名前: %v\n", s)
	}
	// log.Println(rows)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
