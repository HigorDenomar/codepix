package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/higordenomar/codepix/application/grpc/pb"
	"github.com/higordenomar/codepix/application/usecase"
	"github.com/higordenomar/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: &pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)

	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}

	log.Printf("gRPC server has ben started on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}
}
