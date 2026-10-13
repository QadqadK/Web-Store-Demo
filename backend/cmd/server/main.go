package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"

	"webstore/handlers"
	"webstore/services"
)

func main() {
	ctx := context.Background()
	dsn := os.Getenv("DATABASE_URL")
	//if dsn == "" {
	//	dsn = "postgresql://postgres:postgres@localhost:5432/postgres"
	//}
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		panic(err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to db")

	store := services.NewStore(conn)
	handler := handlers.NewHandler(store)

	// items
	http.HandleFunc("/items", handler.GetItems)
	http.HandleFunc("/items/{item_id}", handler.GetItemByID)

	// Start the server on port 8080
	println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
