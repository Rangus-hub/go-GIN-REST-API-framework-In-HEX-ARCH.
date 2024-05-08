package controller

import (
	"api/command"
	"api/repository"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "PostGrePsd"
	dbname   = "Company"
)

func StartServer() {

	// Create a connection pool
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}
	defer pool.Close()

	fmt.Printf("Connected to PostgreSQL database: %s\n", dbname)

	// Create repositories and controllers
	employeeRepository := repository.NewPostgresEmployeeRepository(pool)
	employeeController := command.NewEmployeeController(employeeRepository)

	// Setup Gin router
	router := gin.Default()
	router.GET("/employee", employeeController.GetAllEmployees)
	router.POST("/employees", employeeController.CreateEmployee)
	router.PUT("/employees/:id", employeeController.UpdateEmployee)
	router.GET("/employees/:id", employeeController.GetEmployeeByID)
	router.DELETE("/employees/:id", employeeController.DeleteEmployee)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
