package main

import (
	"flag"
	"fmt"
	"olin-assessment-muhammad-diffa/package/logger"
	"olin-assessment-muhammad-diffa/routers"
)

func init() {
	logger.InitLogger()
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()

	app := routers.SetupRouter()
	app.Run(fmt.Sprintf(":%d", port))
}
