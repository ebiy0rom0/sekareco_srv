// Package api GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2022-07-24 17:51:01.163626862 +0000 UTC m=+1.518556266
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT License",
            "url": "https://github.com/ebiy0rom0/sekareco_srv/blob/develop/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/prsk/person": {
            "put": {
                "description": "wip",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "proseka"
                ],
                "summary": "wip",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0-beta",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "sekareco_srv",
	Description:      "sekareco REST API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
