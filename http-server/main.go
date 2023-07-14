package main

import (
	"fmt"

	"github.com/ZAF07/tiktok-instant-messaging/http-server/config"
)

func main() {
	ap := config.NewAppConfig()
	rpcMaxConn := ap.GetRPCServiceMaxConAge()
	rpcMaxGrace := ap.GetRPCServiceMaxConGrace()
	fmt.Println(rpcMaxConn)
	fmt.Println(rpcMaxGrace)
	fmt.Println("PORT: ", ap.GetHTTPPort())
}
