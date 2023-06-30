// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Hello world plugin API.",
    "title": "station-massa - Hello world",
    "version": "0.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "operationId": "defaultPage",
        "responses": {
          "301": {
            "description": "Move to the default endpoint"
          }
        }
      }
    },
    "/api/goodbye": {
      "put": {
        "produces": [
          "application/json"
        ],
        "operationId": "goodbye",
        "parameters": [
          {
            "type": "string",
            "description": "the name of the person to be greeted.",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Goodbye content.",
            "schema": {
              "type": "object",
              "required": [
                "message"
              ],
              "properties": {
                "message": {
                  "description": "Goodbye message.",
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle."
          }
        }
      }
    },
    "/api/hello": {
      "put": {
        "produces": [
          "application/json"
        ],
        "operationId": "hello",
        "parameters": [
          {
            "type": "string",
            "description": "the name of the person to be greeted.",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Greeting content.",
            "schema": {
              "type": "object",
              "required": [
                "message"
              ],
              "properties": {
                "message": {
                  "description": "Greeting message.",
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle."
          }
        }
      }
    },
    "/web/{resource}": {
      "get": {
        "produces": [
          "application/json",
          "text/javascript",
          "text/html",
          "text/css",
          "text/webp",
          "image/png"
        ],
        "operationId": "web",
        "parameters": [
          {
            "type": "string",
            "description": "Website resource.",
            "name": "resource",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Page found"
          },
          "404": {
            "description": "Resource not found."
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Hello world plugin API.",
    "title": "station-massa - Hello world",
    "version": "0.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "operationId": "defaultPage",
        "responses": {
          "301": {
            "description": "Move to the default endpoint"
          }
        }
      }
    },
    "/api/goodbye": {
      "put": {
        "produces": [
          "application/json"
        ],
        "operationId": "goodbye",
        "parameters": [
          {
            "type": "string",
            "description": "the name of the person to be greeted.",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Goodbye content.",
            "schema": {
              "type": "object",
              "required": [
                "message"
              ],
              "properties": {
                "message": {
                  "description": "Goodbye message.",
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle."
          }
        }
      }
    },
    "/api/hello": {
      "put": {
        "produces": [
          "application/json"
        ],
        "operationId": "hello",
        "parameters": [
          {
            "type": "string",
            "description": "the name of the person to be greeted.",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Greeting content.",
            "schema": {
              "type": "object",
              "required": [
                "message"
              ],
              "properties": {
                "message": {
                  "description": "Greeting message.",
                  "type": "string",
                  "x-nullable": false
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle."
          }
        }
      }
    },
    "/web/{resource}": {
      "get": {
        "produces": [
          "application/json",
          "image/png",
          "text/css",
          "text/html",
          "text/javascript",
          "text/webp"
        ],
        "operationId": "web",
        "parameters": [
          {
            "type": "string",
            "description": "Website resource.",
            "name": "resource",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Page found"
          },
          "404": {
            "description": "Resource not found."
          }
        }
      }
    }
  }
}`))
}
