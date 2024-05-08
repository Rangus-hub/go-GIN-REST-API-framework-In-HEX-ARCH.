// package app

// import (

// 	"context"
// 	"fmt"
// 	"log"
// 	"api/controller"
// 	"github.com/jackc/pgx/v4/pgxpool"
// )

// func StartServer() {
// 	// Create a connection pool
// 	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	pool, err := pgxpool.Connect(context.Background(), connectionString)
// 	if err != nil {
// 		log.Fatalf("Error connecting to PostgreSQL: %v", err)
// 	}
// 	defer pool.Close()

// 	fmt.Printf("Connected to PostgreSQL database: %s\n", dbname)

// 	DevRoutes(pool)

// }
