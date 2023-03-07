package api

import (
	"fmt"
	"log"
	"os"
	"github.com/bachvu236/go-jwt/api/controllers"
	"github.com/joho/godotenv"
)
func init() {}
var server = controllers.Server{}

func Run() {
	var err error = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
		fmt.Printf(os.Getenv("DB_DRIVER"))
	}
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Run(":8080")
}