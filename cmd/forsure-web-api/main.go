package main

import (
	"fmt"
	"os"

	"github.com/forsure-tech/forsure-web-api/internal/server"
)

func main() {

	var port string

	port = os.Getenv("API_PORT")

	if port == "" {
		port = "8080"
	}

	r := server.New()
	r.Run(fmt.Sprintf(":%s", port))

}
