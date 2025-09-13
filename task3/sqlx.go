package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func createEmployeeTable() {
	GlobalDB.AutoMigrate(&Employee{})
}

func insertEmployee() {
	GlobalDB.Create(&Employee{Name: "Alice", Department: "HR", Salary: 60000})
	GlobalDB.Create(&[]Employee{
		{Name: "Bob", Department: "Engineering", Salary: 80000},
		{Name: "Charlie", Department: "Marketing", Salary: 70000},
	})
}

func findEmployee() {
	var sql = "select * from employees where department=?"
	var employees []Employee
	if err := db.Select(&employees, sql, "技术部"); err != nil {
		fmt.Printf("查询失败，%v", err)
	}
	fmt.Println(employees)

	var employee Employee
	sqlStr := "select * from employees order by salary desc limit 1"
	if err := db.Get(&employee, sqlStr); err != nil {
		fmt.Printf("查询失败，%v", err)
	}
	fmt.Println(employee)
}
