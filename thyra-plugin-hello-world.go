package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/loads"
	"github.com/massalabs/thyra-plugin-hello-world/api"
	"github.com/massalabs/thyra-plugin-hello-world/api/server/restapi"
	"github.com/massalabs/thyra-plugin-hello-world/api/server/restapi/operations"
	"github.com/massalabs/thyra-plugin-hello-world/web"
)

const ThyraRegisterEndpoint = "https://my.massa/plugin-manager/register"

type Register struct {
	ID          int64
	Name        string
	Description string
	Logo        string
	Authority   string
	APISpec     string
}

func register(pluginID int64, authority net.Addr, spec string) {
	logo, err := web.Content("logo_massa.webp")
	if err != nil {
		panic(err)
	}

	reg := Register{
		ID:          pluginID,
		Name:        "hello world",
		Description: "A simple hello world plugin.",
		Authority:   "http://" + authority.String(),
		APISpec:     spec,
		Logo:        base64.StdEncoding.EncodeToString(logo),
	}

	body, err := json.Marshal(reg)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", ThyraRegisterEndpoint, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}

	if response.StatusCode == http.StatusNoContent {
		fmt.Fprintf(os.Stdout, "Plugin successfuly registered!\n")
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("response Body:", string(body))
		fmt.Fprintf(os.Stdout, "Registry failed. HTTP status: %d, HTTP body: %v.\n", response.StatusCode, body)
		response.Body.Close()
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
	quit := make(chan bool)
	intSig := make(chan os.Signal, 1)
	signal.Notify(intSig, syscall.SIGINT, syscall.SIGTERM)

	go killTime(quit)

	server := initializeAPI()

	l, err := server.HTTPListener()
	if err != nil {
		panic(err)
	}

	register(1, l.Addr(), string(restapi.SwaggerJSON))

	if err := server.Serve(); err != nil {
		panic(err)
	}

	<-intSig
	quit <- true
}
