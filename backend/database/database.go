package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ChainExpressbill/coldchain/ent"
	_ "github.com/lib/pq"
)

var client *ent.Client
var connErr error
var once sync.Once

func createConnectionURL() string {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)

	return connectionUrl
}

func GetInstance() *ent.Client {
	if client == nil {
		once.Do(func() {
			client, connErr = ent.Open("postgres", createConnectionURL())

			if connErr != nil {
				log.Fatal(connErr)
			}

			fmt.Println("Database connection success")
		})
	}
	return client
}
