package main

type Employee struct {
	ID         int64   `json: "id" binding:"required"`
	Name       string  `json: "name" binding:"required"`
	Department string  `json: "dept" binding:"required"`
	Salary     float64 `json: "salary" binding: "required"`
}

type EmployeeRequest struct {
	Name       string  `json: "name, omitempty" `
	Department string  `json: "dept, omitempty" `
	Salary     float64 `json: "salary,  omitempty" `
}
