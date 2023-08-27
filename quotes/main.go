package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/aalsa16/golang-microservices/proto"
	"github.com/aalsa16/golang-microservices/quotes/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 9002, "quote service port")
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

	lis, err := net.Listen("tcp", fmt.Sprintf("quotes:%d", port))

	if err != nil {
		log.Fatalf("Failed to set up server %v", err)
	}

	db := setupSql()

	grpcServer := grpc.NewServer()

	svc := service.NewQuoteService(db)

	proto.RegisterQuoteServiceServer(grpcServer, svc)

	log.Printf("Authentication service running on [::]:%d\n", port)

	grpcServer.Serve(lis)
}
