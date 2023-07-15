package main

import (
	"fmt"

	"github.com/ZAF07/tiktok-instant-messaging/http-server/config"
)

func main() {
	ap := config.InitAppConfig()
	fmt.Println("HTTP PORT: ", ap.GetPortHTTP())
	fmt.Println("HTTP READ TIMEOUT: ", ap.GetReadTimeoutHTTP())
	fmt.Println("HTTP WRITE TIMEOUT: ", ap.GetWriteTimeoutHTTP())

	// RPC
	fmt.Println("RPC PORT: ", ap.GetPortRPC())
	fmt.Println("RPC METHOD: ", ap.GetMethodRPC())
	fmt.Println("RPC AGE: ", ap.GetMaxConAgeRPC())
	fmt.Println("RPC GRACE: ", ap.GetMaxConGraceRPC())

	// rpcMaxConn := ap.GetRPCServiceMaxConAge()
	// rpcMaxGrace := ap.GetRPCServiceMaxConGrace()
	// fmt.Println("rpc max con: ", rpcMaxConn)
	// fmt.Println("rpc grace: ", rpcMaxGrace)
	// // fmt.Println("http timeout: ", ap.GetReadTimeout())
	// fmt.Println("PORT: ", ap.Http.GetHTTPPort())
	// fmt.Println("PORT: ", ap.Http.Port)

}
