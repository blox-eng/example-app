// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "support@blox.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/work-reports": {
            "get": {
                "description": "Retrieve a list of all work reports",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Work Report"
                ],
                "summary": "Get all work reports",
                "responses": {
                    "200": {
                        "description": "A list of work reports",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_blox-eng_backend_internal_model.WorkReport"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new work report",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Work Report"
                ],
                "summary": "Create a new work report",
                "parameters": [
                    {
                        "description": "Add Work Report",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_internal_model.CreateWorkReport"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_blox-eng_backend_internal_model.WorkReport"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    }
                }
            }
        },
        "/api/work-reports/{id}": {
            "get": {
                "description": "Retrieve details of a specific work report by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Work Report"
                ],
                "summary": "Get a specific work report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Work report ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_blox-eng_backend_internal_model.WorkReport"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing work report by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Work Report"
                ],
                "summary": "Update an existing work report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Work report ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_blox-eng_backend_internal_model.WorkReport"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a work report by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Work Report"
                ],
                "summary": "Delete a work report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Work report ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_blox-eng_backend_pkg_httputil.HTTPErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_blox-eng_backend_internal_model.CreateWorkReport": {
            "type": "object",
            "properties": {
                "quantity_unit": {
                    "type": "string",
                    "example": "Kg"
                },
                "site_id": {
                    "type": "integer",
                    "example": 1
                },
                "work": {
                    "type": "string",
                    "example": "Hammering"
                },
                "work_quantity": {
                    "type": "number",
                    "example": 10
                },
                "worker": {
                    "type": "string",
                    "example": "John Doe"
                }
            }
        },
        "github_com_blox-eng_backend_internal_model.WorkReport": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "quantity_unit": {
                    "type": "string",
                    "example": "Kg"
                },
                "site_id": {
                    "type": "integer",
                    "example": 1
                },
                "updated_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "work": {
                    "type": "string",
                    "example": "Hammering"
                },
                "work_quantity": {
                    "type": "number",
                    "example": 10
                },
                "worker": {
                    "type": "string",
                    "example": "John Doe"
                }
            }
        },
        "github_com_blox-eng_backend_pkg_httputil.HTTPErr": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "err": {}
            }
        },
        "github_com_blox-eng_backend_pkg_httputil.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "status": {}
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
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
	Title:            "BLOX",
	Description:      "This is the backend of Blox App service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
