// controller/book_controller.go

package command

import (
	"api/handler"
	"api/repository"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service handler.EmployeeHandler
}

func NewEmployeeController(repo repository.EmployeeRepository) *EmployeeController {
	return &EmployeeController{handler.NewEmployeeHandler(repo)}
}

func (bc *EmployeeController) GetAllEmployees(c *gin.Context) {
	bc.service.GetAllEmployees(c)
}

func (bc *EmployeeController) CreateEmployee(c *gin.Context) {
	bc.service.CreateEmployee(c)
}

func (bc *EmployeeController) UpdateEmployee(c *gin.Context) {
	bc.service.UpdateEmployee(c)
}

func (bc *EmployeeController) GetEmployeeByID(c *gin.Context) {
	bc.service.GetEmployeeByID(c)
}

func (bc *EmployeeController) DeleteEmployee(c *gin.Context) {
	bc.service.DeleteEmployee(c)
}
