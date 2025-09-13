package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB
var db *sqlx.DB

func initDB() (err error) {

	dsn := "root:123456abc@tcp(192.168.124.13:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return err
	}
	fmt.Println("数据库连接成功", db)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return err
}
func main() {
	fmt.Println("start...")

	dsn := "root:123456abc@tcp(192.168.124.13:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	GlobalDB = db
	fmt.Println("数据库连接成功", db)
	createTable()
	insert()
	query()
	update()
	delete()

	createTowTables()
	insertAccount()
	err = transferFunds(1, 2, 100)
	if err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("转账成功")
	}
	go func() {
		for i := 0; i < 10; i++ {
			err = transferFunds(1, 2, 10)
			if err != nil {
				fmt.Println("转账失败:", err)
			} else {
				fmt.Println("转账成功")
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			err = transferFunds(1, 2, 10)
			if err != nil {
				fmt.Println("转账失败:", err)
			} else {
				fmt.Println("转账成功")
			}
		}
	}()

	time.Sleep(5 * time.Second)

	createEmployeeTable()
	insertEmployee()

	initDB()
	findEmployee()

	createBookTable()
	insertBooks()
	findBooks()

	createAllTable()

	insertAllData()
	findUserAllPostAndComment()
	findMostCommentsPost()
	createPost()
	deleteComment()
}
