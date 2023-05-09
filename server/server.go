package main

import (
	"log"
	"net"

	"grpc-tet/storage"
	"grpc-tet/upload"

	"google.golang.org/grpc"

	uploadpb "grpc-tet/proto"
)

func main() {
	// Initialise TCP listener.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	// Bootstrap upload server.
	uplSrv := upload.NewServer(storage.New("tmp/"))

	// Bootstrap gRPC server.
	rpcSrv := grpc.NewServer()

	// Register and start gRPC server.
	uploadpb.RegisterUploadServiceServer(rpcSrv, uplSrv)
	log.Fatal(rpcSrv.Serve(lis))
}
