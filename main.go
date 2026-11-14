package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

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

func getHelth(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "healthy"})

}

func createmployee(c *gin.Context) {

	var request EmployeeRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "request error"})
		return
	}

	var data Employee
	mu.Lock()
	defer mu.Unlock()
	data.ID = int64(nextID)
	nextID++

	data.Name = request.Name
	data.Salary = request.Salary
	data.Department = request.Department

	employees[int(data.ID)] = data

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
		"data":    data,
	})

}

func getAll(c *gin.Context) {

	mu.RLock()
	defer mu.RLock()
	result := make([]Employee, 0, len(employees))

	for _, emp := range employees {
		result = append(result, emp)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    result,
	})
}

func update(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id incorrect",
		})
		return
	}
	var request EmployeeRequest

	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "binding error",
		})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	emp, exists := employees[int(id)]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id not available"})
		return
	}

	if request.Department != "" {
		emp.Department = request.Department
	}
	if request.Name != "" {
		emp.Name = request.Name
	}
	if request.Salary != emp.Salary {
		emp.Salary = request.Salary
	}

	employees[int(id)] = emp

	c.JSON(http.StatusCreated, gin.H{"message": "update completed",
		"data": employees[int(id)]})
}

func getById(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Param Error"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	emp, exists := employees[int(id)]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success",
		"data": emp})

}
