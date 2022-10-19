// Code generated by go-swagger; DO NOT EDIT.

package rest_server_zrok

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
    "application/zrok.v1+json"
  ],
  "produces": [
    "application/zrok.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "zrok client access",
    "title": "zrok",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/account": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "createAccount",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/accountRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "account created"
          },
          "400": {
            "description": "account not created (already exists)",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/disable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "identity"
        ],
        "operationId": "disable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/disableRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "environment disabled"
          },
          "401": {
            "description": "invalid environment"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/enable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "identity"
        ],
        "operationId": "enable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/enableRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "environment enabled",
            "schema": {
              "$ref": "#/definitions/enableResponse"
            }
          },
          "401": {
            "description": "invalid api key"
          },
          "404": {
            "description": "account not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "login successful",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "401": {
            "description": "invalid login"
          }
        }
      }
    },
    "/overview": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "metadata"
        ],
        "operationId": "overview",
        "responses": {
          "200": {
            "description": "overview returned",
            "schema": {
              "$ref": "#/definitions/environmentServicesList"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "register",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/registerRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "account created",
            "schema": {
              "$ref": "#/definitions/registerResponse"
            }
          },
          "404": {
            "description": "request not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/tunnel": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "tunnel"
        ],
        "operationId": "tunnel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/tunnelRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "tunnel created",
            "schema": {
              "$ref": "#/definitions/tunnelResponse"
            }
          },
          "401": {
            "description": "invalid environment identity",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/untunnel": {
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "tunnel"
        ],
        "operationId": "untunnel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/untunnelRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "tunnel removed"
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/verify": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "verify",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/verifyRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "token ready",
            "schema": {
              "$ref": "#/definitions/verifyResponse"
            }
          },
          "404": {
            "description": "token not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "metadata"
        ],
        "operationId": "version",
        "responses": {
          "200": {
            "description": "current server version",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "accountRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        }
      }
    },
    "authUser": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "disableRequest": {
      "type": "object",
      "properties": {
        "identity": {
          "type": "string"
        }
      }
    },
    "enableRequest": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "host": {
          "type": "string"
        }
      }
    },
    "enableResponse": {
      "type": "object",
      "properties": {
        "cfg": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "environment": {
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "address": {
          "type": "string"
        },
        "createdAt": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "host": {
          "type": "string"
        },
        "updatedAt": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "environmentServices": {
      "type": "object",
      "properties": {
        "environment": {
          "$ref": "#/definitions/environment"
        },
        "services": {
          "$ref": "#/definitions/services"
        }
      }
    },
    "environmentServicesList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environmentServices"
      }
    },
    "environments": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environment"
      }
    },
    "errorMessage": {
      "type": "string"
    },
    "loginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "loginResponse": {
      "type": "string"
    },
    "principal": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "registerRequest": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "registerResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "service": {
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "backend": {
          "type": "string"
        },
        "createdAt": {
          "type": "string"
        },
        "frontend": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "updatedAt": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "services": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/service"
      }
    },
    "tunnelRequest": {
      "type": "object",
      "properties": {
        "authScheme": {
          "type": "string"
        },
        "authUsers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/authUser"
          }
        },
        "endpoint": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "tunnelResponse": {
      "type": "object",
      "properties": {
        "proxyEndpoint": {
          "type": "string"
        },
        "svcName": {
          "type": "string"
        }
      }
    },
    "untunnelRequest": {
      "type": "object",
      "properties": {
        "svcName": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "verifyRequest": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "verifyResponse": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        }
      }
    },
    "version": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/zrok.v1+json"
  ],
  "produces": [
    "application/zrok.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "zrok client access",
    "title": "zrok",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/account": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "createAccount",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/accountRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "account created"
          },
          "400": {
            "description": "account not created (already exists)",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/disable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "identity"
        ],
        "operationId": "disable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/disableRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "environment disabled"
          },
          "401": {
            "description": "invalid environment"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/enable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "identity"
        ],
        "operationId": "enable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/enableRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "environment enabled",
            "schema": {
              "$ref": "#/definitions/enableResponse"
            }
          },
          "401": {
            "description": "invalid api key"
          },
          "404": {
            "description": "account not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "login successful",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "401": {
            "description": "invalid login"
          }
        }
      }
    },
    "/overview": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "metadata"
        ],
        "operationId": "overview",
        "responses": {
          "200": {
            "description": "overview returned",
            "schema": {
              "$ref": "#/definitions/environmentServicesList"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "register",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/registerRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "account created",
            "schema": {
              "$ref": "#/definitions/registerResponse"
            }
          },
          "404": {
            "description": "request not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/tunnel": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "tunnel"
        ],
        "operationId": "tunnel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/tunnelRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "tunnel created",
            "schema": {
              "$ref": "#/definitions/tunnelResponse"
            }
          },
          "401": {
            "description": "invalid environment identity",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/untunnel": {
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "tunnel"
        ],
        "operationId": "untunnel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/untunnelRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "tunnel removed"
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/verify": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "verify",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/verifyRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "token ready",
            "schema": {
              "$ref": "#/definitions/verifyResponse"
            }
          },
          "404": {
            "description": "token not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "metadata"
        ],
        "operationId": "version",
        "responses": {
          "200": {
            "description": "current server version",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "accountRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        }
      }
    },
    "authUser": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "disableRequest": {
      "type": "object",
      "properties": {
        "identity": {
          "type": "string"
        }
      }
    },
    "enableRequest": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "host": {
          "type": "string"
        }
      }
    },
    "enableResponse": {
      "type": "object",
      "properties": {
        "cfg": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "environment": {
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "address": {
          "type": "string"
        },
        "createdAt": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "host": {
          "type": "string"
        },
        "updatedAt": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "environmentServices": {
      "type": "object",
      "properties": {
        "environment": {
          "$ref": "#/definitions/environment"
        },
        "services": {
          "$ref": "#/definitions/services"
        }
      }
    },
    "environmentServicesList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environmentServices"
      }
    },
    "environments": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environment"
      }
    },
    "errorMessage": {
      "type": "string"
    },
    "loginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "loginResponse": {
      "type": "string"
    },
    "principal": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "registerRequest": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "registerResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "service": {
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "backend": {
          "type": "string"
        },
        "createdAt": {
          "type": "string"
        },
        "frontend": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "updatedAt": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "services": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/service"
      }
    },
    "tunnelRequest": {
      "type": "object",
      "properties": {
        "authScheme": {
          "type": "string"
        },
        "authUsers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/authUser"
          }
        },
        "endpoint": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "tunnelResponse": {
      "type": "object",
      "properties": {
        "proxyEndpoint": {
          "type": "string"
        },
        "svcName": {
          "type": "string"
        }
      }
    },
    "untunnelRequest": {
      "type": "object",
      "properties": {
        "svcName": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "verifyRequest": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "verifyResponse": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        }
      }
    },
    "version": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  }
}`))
}
