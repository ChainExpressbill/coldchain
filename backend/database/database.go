package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ChainExpressbill/coldchain/ent"
	"github.com/ChainExpressbill/coldchain/ent/migrate"
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
	SSL_MODE := os.Getenv("SSL_MODE")
	/**
	@see https://www.postgresql.org/docs/current/libpq-ssl.html
	@desc Table 34.1. SSL Mode Descriptions
	|	sslmode	| Eavesdropping protection | MITM protection | Statement
	|	disable |	No	|	No	|	I don't care about security, and I don't want to pay the overhead of encryption.
	|	allow	|	Maybe	|	No	|	I don't care about security, but I will pay the overhead of encryption if the server insists on it.
	|	prefer	|	Maybe	|	No	|	I don't care about encryption, but I wish to pay the overhead of encryption if the server supports it.
	|	require	|	Yes	|	No	|	I want my data to be encrypted, and I accept the overhead. I trust that the network will make sure I always connect to the server I want.
	|	verify-ca	|	Yes	|	Depends on CA policy	|	I want my data encrypted, and I accept the overhead. I want to be sure that I connect to a server that I trust.
	|	verify-full	|	Yes	|	Yes	|	I want my data encrypted, and I accept the overhead. I want to be sure that I connect to a server I trust, and that it's the one I specify.
	**/
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD, SSL_MODE)

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

// scheme 생성
func MigrationDatabase() {
	err := GetInstance().Schema.Create(
		context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
