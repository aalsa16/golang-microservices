package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aalsa16/golang-microservices/api/routes"
	"github.com/aalsa16/golang-microservices/api/server"
	"github.com/aalsa16/golang-microservices/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	addr          string
	authGrpcPort  int
	quoteGrpcPort int
)

func init() {
	flag.StringVar(&addr, "api_addr", "0.0.0.0:8080", "api service address")
	flag.IntVar(&authGrpcPort, "auth_port", 9001, "grpc auth service port")
	flag.IntVar(&quoteGrpcPort, "quote_port", 9002, "grpc quote service port")
	flag.Parse()
}

func main() {
	godotenv.Load("../.env")

	authConn, err := grpc.Dial(fmt.Sprintf("authentication:%d", authGrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}

	quoteConn, err := grpc.Dial(fmt.Sprintf("quotes:%d", quoteGrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}

	authSvcClient := proto.NewAuthenticationServiceClient(authConn)
	quoteSvcClient := proto.NewQuoteServiceClient(quoteConn)
	handlers := routes.NewHandlers(authSvcClient, quoteSvcClient)

	apiServer := server.NewServer(addr, handlers)

	apiServer.Run()
}
