// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/file/delete/{id}": {
            "get": {
                "description": "Delete a file",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "File ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/file/mutiple_upload": {
            "post": {
                "description": "Upload mutiple files",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "files",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/file/retrieve/{id}": {
            "get": {
                "description": "Retrieve a file",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "File ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/file/upload": {
            "post": {
                "description": "Upload a single file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "File uploader API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
