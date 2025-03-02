// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ping": {
            "get": {
                "description": "ping",
                "tags": [
                    "docs"
                ],
                "responses": {
                    "200": {
                        "description": "pong"
                    }
                }
            }
        },
        "/trip": {
            "post": {
                "description": "create empty trip",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "docs"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TripResponse"
                        }
                    }
                }
            }
        },
        "/trip-book": {
            "post": {
                "description": "Book trip",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "docs"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Offer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/trip/{tripId}/offer": {
            "post": {
                "description": "push new offer to kafka topic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "docs"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "param",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Offer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/trips": {
            "get": {
                "description": "get all trips",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "analytics"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.OfferPrice"
                            }
                        }
                    }
                }
            }
        },
        "/trips/{tripId}": {
            "get": {
                "description": "get trip by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "analytics"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "tripId",
                        "name": "tripId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.OfferPrice"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Offer": {
            "type": "object",
            "properties": {
                "arrival": {
                    "type": "string"
                },
                "departure": {
                    "type": "string"
                },
                "pnr": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "tripId": {
                    "type": "integer"
                }
            }
        },
        "model.OfferPrice": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number"
                },
                "tripId": {
                    "type": "integer"
                }
            }
        },
        "model.TripResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:80",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger of API",
	Description:      "Swagger for SOA",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
