package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/loads"
	"github.com/massalabs/station-massa-hello-world/api"
	"github.com/massalabs/station-massa-hello-world/api/server/restapi"
	"github.com/massalabs/station-massa-hello-world/api/server/restapi/operations"
	"github.com/massalabs/station-massa-hello-world/pkg/plugin"
	"github.com/massalabs/station-massa-hello-world/web"
)

func killTime(quit chan bool) {
	ticker := time.NewTicker(5 * time.Second) //nolint:gomnd

	fmt.Fprintf(os.Stdout, "Plugin is initializing.\n")

	for {
		select {
		case <-ticker.C:
			fmt.Fprintf(os.Stdout, "Tic!\n")
		case <-quit:
			fmt.Fprintf(os.Stdout, "Plugin is shutting down.\nBye!\n")

			return
		}
	}
}

func initializeAPI() *restapi.Server {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		panic(err)
	}

	pluginAPI := operations.NewHelloWorldAPI(swaggerSpec)
	server := restapi.NewServer(pluginAPI)

	pluginAPI.WebHandler = operations.WebHandlerFunc(web.Handle)
	pluginAPI.DefaultPageHandler = operations.DefaultPageHandlerFunc(web.DefaultRedirectHandler)
	pluginAPI.HelloHandler = operations.HelloHandlerFunc(api.Hello)

	server.ConfigureAPI()

	return server
}

func main() {
	quit := make(chan bool)
	intSig := make(chan os.Signal, 1)
	signal.Notify(intSig, syscall.SIGINT, syscall.SIGTERM)

	go killTime(quit)

	server := initializeAPI()

	listener, err := server.HTTPListener()
	if err != nil {
		panic(err)
	}

	err = plugin.RegisterPlugin(listener)
	if err != nil {
		panic(err)
	}

	if err := server.Serve(); err != nil {
		panic(err)
	}

	<-intSig
	quit <- true
}
