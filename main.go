package main

import (
	"github.com/tejas-cogo/go-cogoport/routes"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	r := route.SetupRouter()

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
