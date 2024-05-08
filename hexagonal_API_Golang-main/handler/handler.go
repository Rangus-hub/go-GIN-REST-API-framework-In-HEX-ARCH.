// handler/employee_handler.go
/*


 */
package handler

import (
	model "api/models"
	"api/repository"
	"context"
	"fmt"
	"net/http"

	//"github.com/georgysavva/scany/pgxscan"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler interface {
	GetAllEmployees(c *gin.Context)
	CreateEmployee(c *gin.Context)
	UpdateEmployee(c *gin.Context)
	GetEmployeeByID(c *gin.Context)
	DeleteEmployee(c *gin.Context)
}

type employeeHandler struct {
	repo repository.EmployeeRepository
}

func NewEmployeeHandler(repo repository.EmployeeRepository) EmployeeHandler {
	return &employeeHandler{repo}
}

func (h *employeeHandler) GetAllEmployees(c *gin.Context) {
	var employees []model.Employee
	getAllEmployeesQuery := `select * from Employee`

	rows, err := h.repo.Query(context.Background(), getAllEmployeesQuery)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var Employee model.Employee

		err = rows.Scan(&Employee.ID, &Employee.Name, &Employee.Mobile, &Employee.Email)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		employees = append(employees, Employee)
	}

	c.JSON(http.StatusOK, gin.H{
		"Employees": employees,
	})
}

func (h *employeeHandler) CreateEmployee(c *gin.Context) {
	var newEmployee model.Employee
	var insertedEmployee model.Employee

	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	insertEmployeeQuery := `insert into Employee (name, mobile, email)
						values ($1, $2, $3)
						returning *`

	// Execute the query and retrieve the result rows
	rows, err := h.repo.Query(context.Background(), insertEmployeeQuery, newEmployee.Name, newEmployee.Mobile, newEmployee.Email)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	// Scan the first row into insertedEmployee
	if rows.Next() {
		err = rows.Scan(&insertedEmployee.ID, &insertedEmployee.Name, &insertedEmployee.Mobile, &insertedEmployee.Email)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"Employee": insertedEmployee,
	})
}

func (h *employeeHandler) UpdateEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	var updateEmployee model.Employee

	if err := c.ShouldBindJSON(&updateEmployee); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updateEmployeeQuery := `update Employee
					set name = $2, mobile = $3, email = $4
					where id = $1;`

	res, err := h.repo.Exec(context.Background(), updateEmployeeQuery, employeeID, updateEmployee.Name, updateEmployee.Mobile, updateEmployee.Email)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	count := res.RowsAffected()
	if count < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Employee with id %v is not found", employeeID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v row affected. Employee with id %v has been successfully updated", count, employeeID),
	})
}

func (h *employeeHandler) GetEmployeeByID(c *gin.Context) {
	employeeID := c.Param("id")
	var Employee model.Employee

	getEmployeeQuery := `select * from Employee where id = $1;`

	err := h.repo.QueryRow(context.Background(), getEmployeeQuery, employeeID).Scan(&Employee.ID, &Employee.Name, &Employee.Mobile, &Employee.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Employee with id %v is not found", employeeID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Employee": Employee,
	})
}

func (h *employeeHandler) DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	deleteEmployeeQuery := `delete from Employee where id = $1;`

	res, err := h.repo.Exec(context.Background(), deleteEmployeeQuery, employeeID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	count := res.RowsAffected()
	if count < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Employee with id %v is not found", employeeID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v row affected. Employee with id %v has been successfully deleted", count, employeeID),
	})
}

// func (h *employeeHandler) CreateEmployee(c *gin.Context) {
// 	var newEmployee model.Employee
// 	var insertedEmployee model.Employee

// 	if err := c.ShouldBindJSON(&newEmployee); err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	insertEmployeeQuery := `insert into Employee (name, mobile, email)
// 						values ($1, $2, $3)
// 						returning *`

// 	// Execute the query and retrieve the result row
// 	row := h.repo.QueryRow(context.Background(), insertEmployeeQuery, newEmployee.Name, newEmployee.Mobile, newEmployee.Email)

// 	// Declare a slice of model.Employee to store the scanned rows
// 	//var insertedEmployees model.Employee

// 	// Scan all rows directly into the insertedEmployees slice using pgxscan.ScanAll
// 	//err := pgxscan.ScanAll(row, &insertedEmployees)
// 	err := row.Scan(&insertedEmployee.ID, &insertedEmployee.Name, &insertedEmployee.Mobile, &insertedEmployee.Email)

// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

//		c.JSON(http.StatusCreated, gin.H{
//			"Employee": insertedEmployee,
//		})
//	}

// func (h *employeeHandler) GetEmployeeByID(c *gin.Context) {
// 	employeeID := c.Param("id")
// 	var employee model.Employee

// 	getEmployeeQuery := `select * from Employee where id = $1;`

// 	err := h.repo.QueryRow(context.Background(), getEmployeeQuery, employeeID).Scan(&employee.ID, &employee.Name, &employee.Mobile, &employee.Email)

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"message": fmt.Sprintf("Employee with id %v is not found", employeeID),
// 		})
// 		return
// 	}

//		c.JSON(http.StatusOK, gin.H{
//			"Employee": employee,
//		})
//	}
