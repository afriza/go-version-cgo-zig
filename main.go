package main

import (
	"database/sql"
	"fmt"
	"log"
	"runtime/debug"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	params := "?_pragma=foreign_keys(1)" +
		"&_pragma=journal_mode(WAL)" +
		"&_pragma=synchronous(NORMAL)" +
		"&_pragma=busy_timeout(5000)" +
		"&_pragma=journal_size_limit(200000000)" +
		"&_txlock=immediate" +
		""
	db, err := sql.Open("sqlite3", "file::memory:"+params)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	q := "SELECT 1"
	var n int
	err = db.QueryRow(q).Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%q => %v\n", q, n)

	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, s := range bi.Settings {
			fmt.Println(s.Key+":", s.Value)
		}
	}

}
