package main

import "google.golang.org/grpc"

func getListenPort() int {
	// TODO implement fetchign from sysenv
	return 9009
}

func getGRPCOptions() []grpc.ServerOption {
	return []grpc.ServerOption{}
}
