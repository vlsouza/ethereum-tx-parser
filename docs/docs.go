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
        "/currentBlock": {
            "get": {
                "description": "Get the latest parsed block",
                "produces": [
                    "application/json"
                ],
                "summary": "Get current block",
                "responses": {
                    "200": {
                        "description": "A map where the key is 'currentBlock' and the value is the block number",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "/subscribe/{address}": {
            "post": {
                "description": "Subscribe to notifications for incoming/outgoing transactions for a specific Ethereum address",
                "produces": [
                    "application/json"
                ],
                "summary": "Subscribe given an address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ethereum address to subscribe to",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A map where the key is 'subscribed' and the value is a boolean indicating success",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    },
                    "400": {
                        "description": "Address is required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Already subscribed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/transactions/{address}": {
            "get": {
                "description": "Retrieve inbound and outbound transactions for a subscribed Ethereum address",
                "produces": [
                    "application/json"
                ],
                "summary": "Get transactions given an address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ethereum address to retrieve transactions for",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/parser.Transaction"
                            }
                        }
                    },
                    "400": {
                        "description": "Address is required",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "parser.Transaction": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
