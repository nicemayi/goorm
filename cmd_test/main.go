package main

import (
	"fmt"
	"goorm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := goorm.NewEngine("sqlite3", "goorm.db")
	defer engine.Close()

	s := engine.NewSession()
	s.Raw("DROP TABLE IF EXISTS User;").Exec()
	s.Raw("CREATE TABLE User(Name test);").Exec()
	s.Raw("CREATE TABLE User(Name test);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`)  VALUES (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}