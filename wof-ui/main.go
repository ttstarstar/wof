package main

import (
	"os"

	"wof/wof-ui/server"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8100"
	}

	server := server.NewServer()
	server.Run(":" + port)
}
