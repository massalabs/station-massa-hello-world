package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/loads"
	"github.com/massalabs/station-massa-hello-world/api"
	"github.com/massalabs/station-massa-hello-world/api/server/restapi"
	"github.com/massalabs/station-massa-hello-world/api/server/restapi/operations"
	"github.com/massalabs/station-massa-hello-world/web"
	pluginKit "github.com/massalabs/station/plugin-kit"
)

const (
	StandaloneEnvVar = "STANDALONE"
)

func killTime(quit chan bool) {
	ticker := time.NewTicker(5 * time.Second) //nolint:gomnd

	_, err := fmt.Fprintf(os.Stdout, "Plugin is initializing.\n")
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ticker.C:
			_, err := fmt.Fprintf(os.Stdout, "Tic!\n")
			if err != nil {
				log.Println(err)
			}
		case <-quit:
			_, err := fmt.Fprintf(os.Stdout, "Plugin is shutting down.\nBye!\n")
			if err != nil {
				log.Println(err)
			}

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

	if os.Getenv(StandaloneEnvVar) != "1" { // plugin registration is skipped in standalone mode
		listener, err := server.HTTPListener()
		if err != nil {
			panic(err)
		}

		err = pluginKit.RegisterPlugin(listener) // register plugin to Massa Station
		if err != nil {
			panic(err)
		}
	}

	if err := server.Serve(); err != nil {
		panic(err)
	}

	<-intSig
	quit <- true
}
