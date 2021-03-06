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
    "title": "Lost and Found",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api",
  "paths": {
    "/items": {
      "get": {
        "tags": [
          "Read items"
        ],
        "operationId": "ListItem",
        "parameters": [
          {
            "type": "string",
            "name": "kind",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Items list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Item"
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "Item"
        ],
        "operationId": "CreateItem",
        "parameters": [
          {
            "name": "item",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Item Created",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          },
          "400": {
            "description": "Bad Request"
          }
        }
      }
    },
    "/items/{itemId}": {
      "get": {
        "tags": [
          "Read Items"
        ],
        "operationId": "GetItemById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "itemId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Item get",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item Not Found"
          }
        }
      }
    }
  },
  "definitions": {
    "Item": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer"
        },
        "kind": {
          "type": "string"
        },
        "name": {
          "type": "string"
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
    "title": "Lost and Found",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api",
  "paths": {
    "/items": {
      "get": {
        "tags": [
          "Read items"
        ],
        "operationId": "ListItem",
        "parameters": [
          {
            "type": "string",
            "name": "kind",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Items list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Item"
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "Item"
        ],
        "operationId": "CreateItem",
        "parameters": [
          {
            "name": "item",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Item Created",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          },
          "400": {
            "description": "Bad Request"
          }
        }
      }
    },
    "/items/{itemId}": {
      "get": {
        "tags": [
          "Read Items"
        ],
        "operationId": "GetItemById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "itemId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Item get",
            "schema": {
              "$ref": "#/definitions/Item"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item Not Found"
          }
        }
      }
    }
  },
  "definitions": {
    "Item": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer"
        },
        "kind": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
}
