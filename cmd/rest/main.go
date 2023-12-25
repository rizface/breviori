package main

import (
	"github.com/rizface/breviori/database"
	"github.com/rizface/breviori/httpserver"
)

func main() {
	database.StartPG()

	database.StartRedis()

	server := httpserver.NewHTTPServer()

	server.RegisterRoutes(httpserver.RegisterRoutes)

	server.Start()
}
