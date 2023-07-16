package main

import (
	"fmt"

	// "github.com/ZAF07/tiktok-instant-messaging/http-server/app"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/app"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/config"
)

func main() {

	ap := config.InitAppConfig()
	fmt.Println("HTTP PORT: ", ap.GetPortHTTP())
	fmt.Println("HTTP READ TIMEOUT: ", ap.GetReadTimeoutHTTP())
	fmt.Println("HTTP WRITE TIMEOUT: ", ap.GetWriteTimeoutHTTP())
	fmt.Println("HTTP CORS: ", ap.GetCorsAllowOrigins())

	// RPC
	fmt.Println("RPC PORT: ", ap.GetPortRPC())
	fmt.Println("RPC METHOD: ", ap.GetMethodRPC())
	fmt.Println("RPC AGE: ", ap.GetMaxConAgeRPC())
	fmt.Println("RPC GRACE: ", ap.GetMaxConGraceRPC())
	httpApp := app.InitApplication()
	httpApp.Start()
	// fmt.Println(httpApp.GetDatastore())
	/*
		This is where we start the application

		The main package calls the app package to initialise the entire application
		It then starts the HTTP server (initialised by the app package) passing the required dependencies
	*/

}
