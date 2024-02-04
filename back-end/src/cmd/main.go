package main

import "src/internal/services"

func main() {
	server := services.NewVanillaServer()
	server.Initialize()
	server.Run()
}
