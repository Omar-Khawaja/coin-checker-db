package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func openDB(user, pw, host, coin, price, time string) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/", user, pw, host))
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer db.Close()

	createdb := "create database if not exists prices"
	_, err = db.Exec(createdb)
	if err != nil {
		fmt.Println(err)
		return
	}

	usedb := "use prices"
	_, err = db.Exec(usedb)
	if err != nil {
		fmt.Println(err)
		return
	}

	createTable := "create table if not exists prices (coin varchar(20), price varchar(20), time varchar(30))"
	_, err = db.Exec(createTable)
	if err != nil {
		fmt.Println(err)
		return
	}

	insert := fmt.Sprintf("insert into prices values('%s', '%s', '%s')", coin, price, time)
	_, err = db.Exec(insert)
	if err != nil {
		fmt.Println(err)
		return
	}
}
