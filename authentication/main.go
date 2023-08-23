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

	return db
}

func main() {
	godotenv.Load("../.env")

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))

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
