package main

import (
	"fmt"
	"net"

	"github.com/markamdev/notekeeper/internal/server"
	"github.com/markamdev/notekeeper/proto/notekeeper"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting remote-db server")

	lst, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", getListenPort()))
	if err != nil {
		fmt.Println("failed to start listener", err.Error())
		return
	}
	defer lst.Close()

	grpcServer := grpc.NewServer(getGRPCOptions()...)
	notekeeper.RegisterNoteKeeperServer(grpcServer, server.NewGRPCServer())
	grpcServer.Serve(lst)
}
