package main

import "github.com/rizface/breviori/httpserver"

func main() {
	server := httpserver.NewHTTPServer()

	server.RegisterRoutes(httpserver.RegisterRoutes)

	server.Start()
}
