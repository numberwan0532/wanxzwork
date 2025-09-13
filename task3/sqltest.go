package main

import (
	"fmt"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

func createTable() {
	fmt.Println("create table...")
	GlobalDB.AutoMigrate(&Student{})
	fmt.Println("创建数据表成功")
}

func insert() {
	fmt.Println("insert...")
	GlobalDB.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})
	fmt.Println("插入数据成功")
	GlobalDB.Create(&[]Student{{Name: "李四", Age: 14, Grade: "二年级"}, {Name: "王五", Age: 16, Grade: "三年级"}})
}

func query() {
	fmt.Println("query...")
	var students []Student
	GlobalDB.Find(&students)
	fmt.Println("查询数据成功", students)
	var student Student
	GlobalDB.First(&student)
	fmt.Println("查询单条数据成功", student)
	GlobalDB.Where("age = ?", 16).Find(&students)
	fmt.Println("条件查询数据成功", students)
}

func update() {
	fmt.Println("update...")
	GlobalDB.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	fmt.Println("更新数据成功")
}

func delete() {
	fmt.Println("delete...")
	GlobalDB.Where("age < ?", 15).Delete(&Student{})
	fmt.Println("删除数据成功")
}
