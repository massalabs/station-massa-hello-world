// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HelloHandlerFunc turns a function with the right signature into a hello handler
type HelloHandlerFunc func(HelloParams) middleware.Responder

// Handle executing the request and returning a response
func (fn HelloHandlerFunc) Handle(params HelloParams) middleware.Responder {
	return fn(params)
}

// HelloHandler interface for that can handle valid hello params
type HelloHandler interface {
	Handle(HelloParams) middleware.Responder
}

// NewHello creates a new http.Handler for the hello operation
func NewHello(ctx *middleware.Context, handler HelloHandler) *Hello {
	return &Hello{Context: ctx, Handler: handler}
}

/* Hello swagger:route PUT /api/hello hello

Hello hello API

*/
type Hello struct {
	Context *middleware.Context
	Handler HelloHandler
}

func (o *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewHelloParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// HelloOKBody hello o k body
//
// swagger:model HelloOKBody
type HelloOKBody struct {

	// Greeting message.
	// Required: true
	Message string `json:"message"`
}

// Validate validates this hello o k body
func (o *HelloOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HelloOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("helloOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hello o k body based on context it is used
func (o *HelloOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *HelloOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HelloOKBody) UnmarshalBinary(b []byte) error {
	var res HelloOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
