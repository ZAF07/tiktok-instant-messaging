package apis

import (
	rpcmanager "github.com/ZAF07/tiktok-instant-messaging/message-service/internal/infrastructure/rpc-manager"
	"google.golang.org/grpc"
)

/* Driven Adapter (Driven by the core to the outside)
This package implements the interface (ports) of the core service to interact with external services via RPC

This is only the adapter. The initialisation of the RPC client lives in the /pkg directory.

Each core/services needs to have an instance of this package in its struct field to be able to make calls to the outside of the current application ( just like the database implementation)

*/

type RPCServiceClient struct {
	conn   *grpc.ClientConn
	client *rpcmanager.RPCClient
}

func NewRPCServiceClient() *RPCServiceClient {
	return &RPCServiceClient{}
}
