package main

import (
	"Test_Go/server"
)

func main() {
	serverAttr := &server.ServerAttribute{}

	defer func() {
		if err := serverAttr.DB.Close(); err != nil {
			panic(err)
		}
	}()

	err := serverAttr.InitServer()
	if err != nil {
		panic(err)
	}

	// Start server
	port := serverAttr.Config.Port
	if port == "" {
		port = "8080"
	}
	serverAttr.Server.Logger.Fatal(serverAttr.Server.Start(":" + port))
}
