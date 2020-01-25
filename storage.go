package main

import (
	"database/sql"
	"log")

func connect() *sql.DB {
	// 2 variabel sekaligus kembaliannya *db,eror (nil = null)
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/latihan")

	if err != nil {
		log.Fatal(err) // fatal eror warna merah
	}

	return db
}