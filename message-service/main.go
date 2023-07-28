package main

import (
	"github.com/ZAF07/tiktok-instant-messaging/http-server/app"
)

/*
	This is the entry point of the application

	The main package calls the app package to initialise and start the entire application
*/

func main() {

	httpApp := app.InitApplication()
	httpApp.Start()
	// fmt.Println(httpApp.GetDatastore())

}
