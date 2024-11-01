package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/codescalersinternships/Secret-Note-API-SPA-Marwan-Radwan.git/internal/db"
	"github.com/codescalersinternships/Secret-Note-API-SPA-Marwan-Radwan.git/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	dbClient, err := db.NewDbClient()
	if err != nil {
		log.Fatalf("Error connecting to the database.\n %v", err)
	}
	r := server.NewServer(dbClient)

	err = r.Run(uint(port))
	if err != nil {
		log.Fatal("server run error:", err)
	}
	fmt.Println("\nGracefully shutting down HTTP server...")
	fmt.Println("Shutdown complete.")
}
