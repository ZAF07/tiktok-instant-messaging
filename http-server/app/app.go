package app

import (
	"github.com/ZAF07/tiktok-instant-messaging/http-server/config"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/service"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/handlers/httphandler"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/router"

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
	httphandler.HTTPHandler
	// Cache
	// Services
	// Handlers
	// HTTPServer
}

func InitApplication() *App {
	//  Here we initialise all dependencies and construct the App struct
	/* 💡 TODO:
	1. Init HTTP Server 👍
	2. Init DB
	3. Init Cache
	4. Init Services 👍
	5. Init Handlers 👍
	6. Init RPC Client
	*/

	// Load application configs into global struct
	config.LoadConfig()

	// Init all the dependencies
	httpServer := httpmanager.NewHTTPServer()
	db := newDatastore()
	services := service.NewHTTPService()
	handlers := httphandler.NewHTTPHandler(services)

	// Construct the App struct
	a := &App{
		*db,
		*httpServer,
		*handlers,
	}
	return a
}

// Start starts the entire application. It ca
func (a *App) Start() {
	h := a.GetHandler()
	router.NewRouter(h, a.HTTPHandler)
	a.StartServer()

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
