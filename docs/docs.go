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
        "/article/add": {
            "post": {
                "description": "Upload an article along with an image",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Article Title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Article Description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Article Summary",
                        "name": "summary",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Article Category",
                        "name": "category",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Article Image",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/article/all": {
            "get": {
                "description": "get all articles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "responses": {}
            }
        },
        "/article/byCategory/{categoryId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Get Articles By category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "categoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/article/edit": {
            "put": {
                "description": "Edit an article",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Article Title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Article Description",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Article Summary",
                        "name": "summary",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Article Category",
                        "name": "category",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Article Status",
                        "name": "status",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Article id",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Article Image",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/article/property/add": {
            "post": {
                "tags": [
                    "articles"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Property description",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Article id",
                        "name": "article_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Property image",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/article/property/edit": {
            "put": {
                "tags": [
                    "articles"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Property description",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Article id",
                        "name": "property_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Property image",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/article/property/{propertyId}": {
            "delete": {
                "tags": [
                    "articles"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Property ID",
                        "name": "propertyId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/article/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Get Article By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Delete Article By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "User registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginCredential"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/auth.AuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/register/send-otp": {
            "post": {
                "description": "User registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RequiredOtpRegistering"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/auth/register/validate-otp": {
            "post": {
                "description": "User otp validation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ValidateOtpRegistering"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/category/add": {
            "post": {
                "description": "Adding new category by admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Add Category",
                "parameters": [
                    {
                        "description": "Adding new category",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/category.NewCategory"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/category/all": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get Categories",
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/category/tree": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get Categorys with tree format",
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/category/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get Category By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Delete Category By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/category/{id}/{name}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Edit Category By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "New Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.AuthResponse": {
            "type": "object",
            "required": [
                "first_name",
                "last_name",
                "phone"
            ],
            "properties": {
                "first_name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                },
                "id": {
                    "type": "integer"
                },
                "jwt": {
                    "$ref": "#/definitions/types.JwtTokens"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 40,
                    "minLength": 3
                },
                "phone": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "auth.LoginCredential": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                },
                "phone": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                }
            }
        },
        "auth.RegisterResponse": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "auth.RequiredOtpRegistering": {
            "type": "object",
            "required": [
                "phone"
            ],
            "properties": {
                "phone": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                }
            }
        },
        "auth.ValidateOtpRegistering": {
            "type": "object",
            "required": [
                "otp",
                "phone"
            ],
            "properties": {
                "otp": {
                    "type": "string",
                    "maxLength": 5,
                    "minLength": 5
                },
                "phone": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                }
            }
        },
        "category.NewCategory": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 50
                },
                "parent_id": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "types.JwtTokens": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                },
                "token": {
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
