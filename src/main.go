package main

import (
	"runtime"
	"template-api-fiber-golang/src/healthcheck"
	_server "template-api-fiber-golang/src/server"
)

func main() {
	// Run go routines in parallelism based on CPUs of your server
	runtime.GOMAXPROCS(runtime.NumCPU())
	// create server
	server := _server.CreateServer(8080)
	// register routes
	healthcheck.RegisterRoutes(*server.Router)
	// start server
	server.StartServer()
}
