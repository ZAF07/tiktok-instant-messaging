package rpcmanager

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
	This package provides an RPC client needed to communicate to upstream services via RPC

	This is where i should init a new instance of the upstream's RPC client
*/

type RPCClient struct {
	conn *grpc.ClientConn
	// This should be the rpc-server's RPC client imported here and passing the conn to init a new instance
	// client
}

// Returns a new instance of RPCClient{}
func NewGRPCClient(port string) *RPCClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connecting to RPC server. error msg: %v", err)
	}

	return &RPCClient{
		conn: conn,
	}
}

func (r *RPCClient) GetConn() *grpc.ClientConn {
	return r.conn
}
