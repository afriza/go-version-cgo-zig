package main

import (
	"database/sql"
	"log"

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
	var n int
	err = db.QueryRow("SELECT 1").Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(n)
}
