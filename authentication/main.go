package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aalsa16/golang-microservices/authentication/service"
	"github.com/aalsa16/golang-microservices/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 9001, "authentication service port")
	flag.Parse()
}

func setupSql() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URI"))

	if err != nil {
		log.Fatalf("Failed to connect to db %v", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY AUTO_INCREMENT, email VARCHAR(255) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL, uuid VARCHAR(255) NOT NULL UNIQUE, refresh_token TEXT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS quotes (id INT AUTO_INCREMENT, quote VARCHAR(255), author VARCHAR(255) NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, owner_uuid VARCHAR(255) NOT NULL, PRIMARY KEY (id), FOREIGN KEY (owner_uuid) REFERENCES users(uuid))")
	if err != nil {
		log.Fatalf(err.Error())
	}

	return db
}

func main() {
	godotenv.Load("../.env")

	lis, err := net.Listen("tcp", fmt.Sprintf("authentication:%d", port))

	if err != nil {
		log.Fatalf("Failed to set up server %v", err)
	}

	db := setupSql()

	grpcServer := grpc.NewServer()

	svc := service.NewAuthService(db)

	proto.RegisterAuthenticationServiceServer(grpcServer, svc)

	log.Printf("Authentication service running on [::]:%d\n", port)

	grpcServer.Serve(lis)
}
