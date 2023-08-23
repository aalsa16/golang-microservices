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
	addr     string
	grpcPort int
)

func init() {
	flag.StringVar(&addr, "auth_addr", "127.0.0.1:8080", "api service address")
	flag.IntVar(&grpcPort, "port", 9001, "grpc service port")
	flag.Parse()
}

func main() {
	godotenv.Load("../.env")

	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", grpcPort), grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}

	authSvcClient := proto.NewAuthenticationServiceClient(conn)
	handlers := routes.NewHandlers(authSvcClient)

	apiServer := server.NewServer(addr, handlers)

	apiServer.Run()
}
