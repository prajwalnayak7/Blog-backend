package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	// Connect to the database server
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	fmt.Printf("[SERVER] Connecting to %s ", DBURL)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("[SERVER] Cannot connect to %s database ", Dbdriver)
		log.Fatal(" due to the error: ", err)
	} else {
		fmt.Printf("[SERVER] Successfuly connected to %s database ", Dbdriver)
	}

	// Database migration
	// server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("[SERVER] Listening on port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
