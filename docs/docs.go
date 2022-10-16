// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/datail": {
            "get": {
                "description": "create account by account and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/api/auth"
                ],
                "summary": "Resiter an account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "442": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    }
                }
            }
        },
        "/api/auth/login": {
            "post": {
                "description": "login account by account and password return AccessToken",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/api/auth"
                ],
                "summary": "Login an account",
                "parameters": [
                    {
                        "description": "Account Telephone",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Account password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "442": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    }
                }
            }
        },
        "/api/auth/resiter": {
            "post": {
                "description": "create account by account and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/api/auth"
                ],
                "summary": "Resiter an account",
                "parameters": [
                    {
                        "description": "Account Name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Account Telephone",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Account password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "442": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    }
                }
            }
        },
        "/api/sensor/all/{type}": {
            "get": {
                "description": "get all sensor data by list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/api/sensor/"
                ],
                "summary": "Get all Sensor Data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "temperature",
                        "name": "type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "442": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    }
                }
            }
        },
        "/api/sensor/recent/{type}/{limit}": {
            "get": {
                "description": "get recent sensor data by list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "/api/sensor/"
                ],
                "summary": "Get recent Sensor Data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "temperature",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "Number",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "442": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseJson"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.ResponseJson": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "petstore.swagger.io",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Libary",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
