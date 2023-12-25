package main

import (
	"github.com/rizface/breviori/database"
	"github.com/rizface/breviori/httpserver"
)

func main() {
	database.NewPGX()

	server := httpserver.NewHTTPServer()

	server.RegisterRoutes(httpserver.RegisterRoutes)

	server.Start()
}
