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
            "email": "cardoso.business.ctt@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "List accounts.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "List accounts",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Page ID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maximum": 10,
                        "minimum": 5,
                        "type": "integer",
                        "description": "Page Size",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Account"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create an account. The client must create and log in a user before.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Create an account",
                "parameters": [
                    {
                        "type": "string",
                        "name": "currency",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/db.Account"
                        }
                    },
                    "400": {
                        "description": "Account already exists!"
                    },
                    "401": {
                        "description": "Unauthorized user!"
                    }
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get an account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Get an account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Account"
                        }
                    }
                }
            }
        },
        "/loans": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a loan.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loans"
                ],
                "summary": "Create a loan",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "name": "account_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/db.LoanTransferTransactionResult"
                        }
                    }
                }
            }
        },
        "/transfers": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "List transfers.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transfers"
                ],
                "summary": "List transfers",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Page ID",
                        "name": "page_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maximum": 10,
                        "minimum": 5,
                        "type": "integer",
                        "description": "Page Size",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Transfer"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a transfer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transfers"
                ],
                "summary": "Create a transfer",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "currency",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "name": "from_account_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "from_account_owner",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "name": "to_account_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "to_account_owner",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/db.TransferTransactionResult"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create an user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create an user",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "full_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 6,
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/api.userResponse"
                        }
                    },
                    "400": {
                        "description": "User already exists!"
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login an user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login an user",
                "parameters": [
                    {
                        "minLength": 6,
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.loginUserResponse"
                        }
                    },
                    "400": {
                        "description": "Username or password incorrect!"
                    }
                }
            }
        }
    },
    "definitions": {
        "api.loginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.userResponse"
                }
            }
        },
        "api.userResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password_changed_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "db.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "db.Entry": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "amount": {
                    "description": "can be negative or positive",
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "db.LoanTransfer": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "amount": {
                    "description": "must be positive",
                    "type": "integer"
                },
                "end_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "interest_rate": {
                    "description": "must be positive",
                    "type": "string"
                },
                "open": {
                    "type": "boolean"
                },
                "start_at": {
                    "type": "string"
                }
            }
        },
        "db.LoanTransferTransactionResult": {
            "type": "object",
            "properties": {
                "loan": {
                    "$ref": "#/definitions/db.LoanTransfer"
                },
                "to_account": {
                    "$ref": "#/definitions/db.Account"
                },
                "to_entry": {
                    "$ref": "#/definitions/db.Entry"
                }
            }
        },
        "db.Transfer": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "must be positive",
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "from_account_id": {
                    "type": "integer"
                },
                "from_account_owner": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "to_account_id": {
                    "type": "integer"
                },
                "to_account_owner": {
                    "type": "string"
                }
            }
        },
        "db.TransferTransactionResult": {
            "type": "object",
            "properties": {
                "from_account": {
                    "$ref": "#/definitions/db.Account"
                },
                "from_entry": {
                    "$ref": "#/definitions/db.Entry"
                },
                "to_account": {
                    "$ref": "#/definitions/db.Account"
                },
                "to_entry": {
                    "$ref": "#/definitions/db.Entry"
                },
                "transfer": {
                    "$ref": "#/definitions/db.Transfer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "United Atomic Bank API Documentation",
	Description:      "This is the United Atomic Bank API. All features available in this application are documented below.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
