package main

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	employees = make(map[int]Employee)
	nextID    = 2
	mu        sync.RWMutex
)

func main() {

	data := Employee{ID: 1, Name: "one", Department: "HR", Salary: 0.01}
	employees[1] = data

	r := gin.Default()

	r.GET("/health", getHelth)
	r.POST("/employees", createmployee)
	r.GET("/getEmployees", getAll)
	r.PUT("/employees/:id", update)
	r.GET("/employees/:id", getById)

	r.Run(":8080")

}
