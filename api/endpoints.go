package api

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/massalabs/station-massa-hello-world/api/server/restapi/operations"
)

func Hello(params operations.HelloParams) middleware.Responder {
	target := "World"
	if params.Name != nil {
		target = *params.Name
	}

	return operations.NewHelloOK().WithPayload(
		&operations.HelloOKBody{Message: fmt.Sprintf("Hello, %s!", target)},
	)
}

func Goodbye(params operations.GoodbyeParams) middleware.Responder {
	target := "World"
	if params.Name != nil {
		target = *params.Name
	}

	return operations.NewGoodbyeOK().WithPayload(
		&operations.GoodbyeOKBody{Message: fmt.Sprintf("Goodbye, %s!", target)},
	)
}
