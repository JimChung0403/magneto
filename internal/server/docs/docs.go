// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-12-25 14:41:21.51466 +0800 CST m=+0.030352252

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "Swagger test_swag API",
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
        "version": "1.0"
    },
    "basePath": "/",
    "paths": {
        "/v1/example/testget/{name}": {
            "get": {
                "description": "用GET method 測試api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "EX1"
                ],
                "summary": "用GET method 測試api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Jim",
                        "description": "名字1",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 18,
                        "description": "年紀",
                        "name": "old",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/apimodel.Response"
                        }
                    }
                }
            }
        },
        "/v1/example/testpost": {
            "post": {
                "description": "用POST method 測試api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "EX2"
                ],
                "summary": "用POST method 測試api",
                "parameters": [
                    {
                        "description": "body",
                        "name": "sendBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/apimodel.ExamplePost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/apimodel.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apimodel.ExamplePost": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "jim"
                },
                "old": {
                    "type": "integer",
                    "example": 30
                }
            }
        },
        "apimodel.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
