package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/routes"
	"github.com/ChandelShikha/go-cogoport/models"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	config.Connect()
	// workers.Workers()

	r := routes.SetupRouter()

	port := os.Getenv("port")

	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080" //localhost
	}
	type Job interface {
		Run()
	}

	r.Run(":" + port)

}
