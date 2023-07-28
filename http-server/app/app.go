package app

import (
	"github.com/ZAF07/tiktok-instant-messaging/http-server/config"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/adapters/datastore/cache"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/adapters/handlers/httphandler"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/service"

	cachemanager "github.com/ZAF07/tiktok-instant-messaging/http-server/internal/infrastructure/cache-manager"
	httpmanager "github.com/ZAF07/tiktok-instant-messaging/http-server/internal/infrastructure/http-manager"
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
	httpmanager.HTTPManager
	httphandler.HTTPHandler
	cache.Cache
}

func InitApplication() *App {
	//  Here we initialise all dependencies and construct the App struct
	/* 💡 TODO:
	1. Init HTTP Server 👍
	2. Init Cache 👍
	4. Init Services 👍
	5. Init Handlers 👍
	6. Init RPC Client
	*/

	// Load application configs into global struct
	config.LoadConfig()

	// Init all the dependencies
	httpServer := httpmanager.NewHTTPServer()

	cacheClient := cachemanager.NewRedisClient()
	cache := cache.NewCache(cacheClient.Client) // 💡 TODO: Implement the initialisation of the redis client (NOT THE ADAPTER)
	services := service.NewHTTPService(cache)
	handlers := httphandler.NewHTTPHandler(services)

	// Construct the App struct
	a := &App{
		*httpServer,
		*handlers,
		*cache,
	}
	return a
}

// Start starts the entire HTTP-SERVER application
func (a *App) Start() {
	// Initialise all service routes and start the server
	h := a.GetHandler()
	a.HTTPHandler.InitRoutes(h)
	a.StartServer()

}
