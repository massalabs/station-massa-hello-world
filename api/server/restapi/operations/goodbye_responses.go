// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GoodbyeOKCode is the HTTP code returned for type GoodbyeOK
const GoodbyeOKCode int = 200

/*GoodbyeOK Goodbye content.

swagger:response goodbyeOK
*/
type GoodbyeOK struct {

	/*
	  In: Body
	*/
	Payload *GoodbyeOKBody `json:"body,omitempty"`
}

// NewGoodbyeOK creates GoodbyeOK with default headers values
func NewGoodbyeOK() *GoodbyeOK {

	return &GoodbyeOK{}
}

// WithPayload adds the payload to the goodbye o k response
func (o *GoodbyeOK) WithPayload(payload *GoodbyeOKBody) *GoodbyeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the goodbye o k response
func (o *GoodbyeOK) SetPayload(payload *GoodbyeOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GoodbyeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GoodbyeInternalServerErrorCode is the HTTP code returned for type GoodbyeInternalServerError
const GoodbyeInternalServerErrorCode int = 500

/*GoodbyeInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response goodbyeInternalServerError
*/
type GoodbyeInternalServerError struct {
}

// NewGoodbyeInternalServerError creates GoodbyeInternalServerError with default headers values
func NewGoodbyeInternalServerError() *GoodbyeInternalServerError {

	return &GoodbyeInternalServerError{}
}

// WriteResponse to the client
func (o *GoodbyeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
