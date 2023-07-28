package app

import (
	"log"
	"net"

	"github.com/ZAF07/tiktok-instant-messaging/http-server/config"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/adapters/datastore/cache"
	msghandler "github.com/ZAF07/tiktok-instant-messaging/http-server/internal/adapters/handlers/message_handler"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/ports"
	"github.com/ZAF07/tiktok-instant-messaging/http-server/internal/core/service"
	"github.com/soheilhy/cmux"

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
	msghandler.HTTPHandler
	// cache.Cache
	c      ports.ICacheStore
	config *config.ApplicationConfig
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
	config := config.LoadConfig()

	// Init all the dependencies
	httpServer := httpmanager.NewHTTPServer()

	cacheClient := cachemanager.NewRedisClient()
	cache := cache.NewRedisCache(cacheClient.Client)

	messageService := service.NewHTTPService(cache)
	handlers := msghandler.NewHTTPHandler(messageService)

	// Construct the App struct
	a := &App{
		*httpServer,
		*handlers,
		cache,
		config,
	}
	return a
}

// Start starts the entire HTTP-SERVER application
func (a *App) Start() {
	h := a.GetHandler()
	a.HTTPHandler.InitRoutes(h)
	a.startServer()
}

func (a *App) startServer() {
	port := a.config.GetPortHTTP()
	network := a.config.GetHTTPNetwork()

	lis, err := net.Listen(network, port)
	if err != nil {
		log.Fatalf("Error connecting tcp port %s: error: %v", port, err)
	}
	Mux := cmux.New(lis)

	httpListener := Mux.Match(cmux.HTTP1Fast())

	go a.StartHTTPServer(httpListener)

	if err := Mux.Serve(); err != nil {
		log.Fatalf("MUX ERR : %+v", err)
	}
}
