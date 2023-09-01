package main

import (
	"fmt"
	"geeorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}

/*
	sqlite3 命令
conn to db:
		(base) ➜ day1-database-sql (main) ✗ sqlite3
		SQLite version 3.41.2 2023-03-22 11:56:21
		Enter ".help" for usage hints.
		Connected to a transient in-memory database.
		Use ".open FILENAME" to reopen on a persistent database.
		sqlite>

show dbs:
		sqlite> .databases
		main: /Users/hq/GeeVong/SimpleGo/db/gee-orm/day1-database-sql/gee.db r/w


open gee.db
		(base) ➜ day1-database-sql (main) ✗ sqlite3 gee.db
		SQLite version 3.41.2 2023-03-22 11:56:21
		Enter ".help" for usage hints.
		sqlite>

show table:
		sqlite> .table
		User
		sqlite>

show schema:
		sqlite> .schema User
		CREATE TABLE User(Name text);


show table data:
		sqlite> SELECT * FROM User;
		Tom
		Sam


creat table:
		sqlite> create table item(Name text, Price integer);


inert data:(需要注意大小写)
		sqlite> INSERT INTO item(Name, Price) VALUES ("华为mate60", 9999), ("apple14", 88888);
		sqlite> select * from item;
		华为mate60|9999
		apple14|88888

*/
