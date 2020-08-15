package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/prajwalnayak7/Blog/api/controller"
)

var server = controller.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("[SERVER] Environment variables file not found")
	}
}

func Run() {
	fmt.Println("[SERVER] Starting the server")

	// Load the environment variables
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("[SERVER] Error loading environment variables, %v", err)
	} else {
		fmt.Println("[SERVER] Loading environment variables")
	}

	// Connect to the Database
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// Make initial entries to the database
	// seed.Load(server.DB)

	// Expose a port for the server to run
	server.Run(":8080")

}
