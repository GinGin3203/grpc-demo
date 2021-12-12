package main

import (
	"context"
	pb "github.com/GinGin3203/grpc-demo/proto"
	"github.com/GinGin3203/grpc-demo/server/repository"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type usersServer struct {
	pb.UnimplementedUsersServiceServer
	repo repository.Repository
}

func main() {
	conn, err := initializeDBConn()
	if err != nil {
		log.Fatalln("error during db init: ", err)
	}
	defer conn.Close(context.Background())
	log.Println("connected to db")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}

	srv := &usersServer{
		repo: repository.New(conn),
	}
	grpcServer := grpc.NewServer(grpc.MaxConcurrentStreams(5))

	reflection.Register(grpcServer)
	pb.RegisterUsersServiceServer(grpcServer, srv)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve: ", err)
	}
}

func initializeDBConn() (conn *pgx.Conn, err error) {
	conn, err = pgx.Connect(context.Background(), "postgres://postgres:example@pg:5432/postgres")
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS users(
		id         int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name       VARCHAR(100) NOT NULL,
		role       VARCHAR(100) NOT NULL,
		last_updated_at timestamp    NOT NULL DEFAULT now()
	);`)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
