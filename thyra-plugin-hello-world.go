package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/loads"
	"github.com/massalabs/thyra-plugin-hello-world/api"
	"github.com/massalabs/thyra-plugin-hello-world/api/server/restapi"
	"github.com/massalabs/thyra-plugin-hello-world/api/server/restapi/operations"
	"github.com/massalabs/thyra-plugin-hello-world/pkg/plugin"
	"github.com/massalabs/thyra-plugin-hello-world/web"
)

const logoFile = "logo_massa.webp"

func register(pluginID string, socket net.Addr, spec string) {
	err := plugin.Register(
		pluginID,
		"hello world", "massalabs",
		"A simple hello world plugin.",
		socket,
		spec,
		fmt.Sprintf("web/%s", logoFile),
	)
	if err != nil {
		panic(fmt.Errorf("while registering plugin: %w", err))
	}
}

func killTime(quit chan bool) {
	ticker := time.NewTicker(5 * time.Second) //nolint:gomnd

	fmt.Fprintf(os.Stderr, "Plugin is initializing.\n")

	for {
		select {
		case <-ticker.C:
			fmt.Fprintf(os.Stdout, "Tic!\n")
		case <-quit:
			fmt.Fprintf(os.Stderr, "Plugin is shutting down.\nBye!\n")

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
	//nolint:gomnd
	if len(os.Args) != 2 {
		panic("this program must be run with correlation id argument!")
	}

	correlationID := os.Args[1]

	quit := make(chan bool)
	intSig := make(chan os.Signal, 1)
	signal.Notify(intSig, syscall.SIGINT, syscall.SIGTERM)

	go killTime(quit)

	server := initializeAPI()

	l, err := server.HTTPListener()
	if err != nil {
		panic(err)
	}

	register(correlationID, l.Addr(), string(restapi.SwaggerJSON))

	if err := server.Serve(); err != nil {
		panic(err)
	}

	<-intSig
	quit <- true
}
