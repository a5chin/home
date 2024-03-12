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
        "/articles": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "記事一覧取得 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.GetArticlesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/viewers": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LP"
                ],
                "summary": "閲覧人数取得 API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.GetTotalViewersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.GetArticlesResponse": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Article"
                    }
                }
            }
        },
        "controller.GetTotalViewersResponse": {
            "type": "object",
            "properties": {
                "viewers": {
                    "type": "integer"
                }
            }
        },
        "entity.Article": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "favorite": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "viewer": {
                    "type": "integer"
                }
            }
        },
        "entity.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "home",
	Description:      "home",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
