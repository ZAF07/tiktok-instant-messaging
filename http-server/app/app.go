package app

import (
	"fmt"
	"log"

	httpmanager "github.com/ZAF07/tiktok-instant-messaging/http-server/pkg/http-manager"
)

/*
	The app package is where we initialise all required dependencies of the application.

	Example:
	1. Datastore (Driven Adapter)
	2. Services (Core)
	3. Handlers aka controllers (Driving Adapters)
	4. HTTP/RPC Servers (Driving Adapter)
*/

type App struct {
	datastore
	httpmanager.HTTPManager
	// Cache
	// Services
	// Handlers
	// HTTPServer
}

func InitApplication() *App {
	//  Here we initialise all dependencies and construct the App struct
	/*
		1. Init HTTP Server
		2. Init DB
		3. Init Cache
		4. Init Services
		5. Init Handlers
	*/
	httpServer := httpmanager.NewHTTPServer()
	db := newDatastore()
	a := &App{
		*db,
		*httpServer,
	}
	return a
}
func (a *App) Start() {
	// Run the server
	fmt.Println("SERVER STARTING..")
	if err := a.GetServer().ListenAndServe(); err != nil {
		log.Fatalf("error starting server. error msg: %v", err)
	}
}

// DB STUFF
type datastore struct {
	db string
}

func (ds *datastore) GetDatastore() string {
	return ds.db
}

func newDatastore() *datastore {
	return &datastore{
		db: "DATABASE",
	}
}
