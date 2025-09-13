package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Auhtor string
	Price  float64
}

func createBookTable() {
	GlobalDB.AutoMigrate(&Book{})
}

func insertBooks() {
	GlobalDB.Create(&[]Book{
		{Title: "book1", Auhtor: "author1", Price: 20},
		{Title: "book2", Auhtor: "author1", Price: 40},
		{Title: "book3", Auhtor: "author1", Price: 50},
		{Title: "book4", Auhtor: "author1", Price: 60},
		{Title: "book5", Auhtor: "author1", Price: 70}},
	)
}

func findBooks() {
	sqlStr := "select * from books where price > ?"
	var books []Book
	if err := db.Select(&books, sqlStr, 50); err != nil {
		fmt.Printf("查询失败，%v", err)
	}
	fmt.Println(books)
}
