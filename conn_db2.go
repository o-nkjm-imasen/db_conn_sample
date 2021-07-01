// work well

// on windows
// 1.install client solutions packages(ODBC driver)

// 2.go get github.com/alexbrainman/odbc

// on Linux
// 1.sudo apt install libodbc1 odbcinst odbcinst1debian2 unixodbc unixodbc-dev

// 1.install client solutions packages(ODBC driver)
// 2.go get github.com/alexbrainman/odbc

package main

import (
	"database/sql"
	"log"

	_ "github.com/alexbrainman/odbc" // https://github.com/alexbrainman/odbc/wiki/GettingStartedOnWindows

	_ "flag"
	"fmt"
	_ "os"
)

func main() {
    //win
    // db, err := sql.Open("odbc", "DRIVER=iSeries Access ODBC Driver;SYSTEM=192.168.10.7; UID=ftpprf; PWD=ftpprf")
    //linux
    db, err := sql.Open("odbc","DRIVER=IBM i Access ODBC Driver 64-bit;SYSTEM=192.168.10.7; UID=ftpprf; PWD=ftpprf")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db)
	defer db.Close()

	rows, err := db.Query("SELECT VMBCO FROM IMAWRK.VDMFPDMY LIMIT 10")
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
		fmt.Printf("部品: %v\n", s)
	}
	// log.Println(rows)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
