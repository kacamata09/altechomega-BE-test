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
        "/author": {
            "get": {
                "description": "Retrieve a list of all authors",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Author"
                ],
                "summary": "Get all authors",
                "operationId": "get-all-authors",
                "responses": {
                    "50": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new author record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Author"
                ],
                "summary": "Create a new author",
                "operationId": "create-author",
                "parameters": [
                    {
                        "description": "Author data",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Author"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Author"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            }
        },
        "/author/{id}": {
            "get": {
                "description": "Retrieve an author by their unique ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Author"
                ],
                "summary": "Get an author by ID",
                "operationId": "get-author-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Author"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing author record by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Author"
                ],
                "summary": "Update an author by ID",
                "operationId": "update-author-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated author data",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Author"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Author"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an author record by their ID",
                "tags": [
                    "Author"
                ],
                "summary": "Delete an author by ID",
                "operationId": "delete-author-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            }
        },
        "/author/{id}/books": {
            "get": {
                "description": "Retrieve an author by their ID along with a list of their books",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Author"
                ],
                "summary": "Get an author with books by ID",
                "operationId": "get-author-with-books-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Author"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            }
        },
        "/book": {
            "get": {
                "description": "Retrieve a list of all books",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get all books",
                "operationId": "get-all-books",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Book"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new book record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Create a new book",
                "operationId": "create-book",
                "parameters": [
                    {
                        "description": "Book data",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Book"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            }
        },
        "/book/{id}": {
            "get": {
                "description": "Retrieve a book by its unique ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get a book by ID",
                "operationId": "get-book-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Book"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing book record by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Update a book by ID",
                "operationId": "update-book-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated book data",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a book record by its ID",
                "tags": [
                    "Book"
                ],
                "summary": "Delete a book by ID",
                "operationId": "delete-book-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper_http.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Author": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "book": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Book"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Book": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "pages": {
                    "type": "integer"
                },
                "publish_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "helper_http.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {},
                "meta": {},
                "status": {
                    "type": "integer"
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