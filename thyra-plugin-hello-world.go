package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-openapi/loads"
	"github.com/massalabs/thyra-plugin-hello-world/api"
	"github.com/massalabs/thyra-plugin-hello-world/api/server/restapi"
	"github.com/massalabs/thyra-plugin-hello-world/api/server/restapi/operations"
	"github.com/massalabs/thyra-plugin-hello-world/pkg/plugin"
	"github.com/massalabs/thyra-plugin-hello-world/web"
)

func register(pluginID int64, socket net.Addr, spec string) {
	logo, err := web.Content("logo_massa.webp")
	if err != nil {
		panic(err)
	}

	err = plugin.Register(
		pluginID,
		"hello world", "massalabs",
		"A simple hello world plugin.",
		socket,
		string(restapi.SwaggerJSON),
		logo,
	)
	if err != nil {
		panic(fmt.Errorf("while registring plugin: %w", err))
	}

}

func killTime(quit chan bool) {
	ticker := time.NewTicker(5 * time.Second)

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

	if len(os.Args) != 2 {
		panic("this program must be run with correlation id argument!")
	}

	firstArg := os.Args[1]

	correlationID, err := strconv.ParseInt(firstArg, 10, 64)
	if err != nil {
		panic(err)
	}

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
