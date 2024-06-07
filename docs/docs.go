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
        "/api/admin/users": {
            "post": {
                "description": "Get all users from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Data_users"
                            }
                        }
                    },
                    "400": {
                        "description": "You not admin",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "No data found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to query data: [Error Message]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/jwt/test": {
            "post": {
                "description": "Создание JWT токена для пользователя \"test_user\"",
                "produces": [
                    "application/json"
                ],
                "summary": "Получение JWT токена",
                "operationId": "createJwtToken",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Token"
                        }
                    }
                }
            }
        },
        "/api/jwt/verify": {
            "post": {
                "description": "Проверка валидности переданного тестового JWT токена",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Проверка валидности JWT токена",
                "operationId": "verifyJwtToken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer your_token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.jwtValidationResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.jwtValidationResult"
                        }
                    }
                }
            }
        },
        "/api/sockets/thermalmapdata": {
            "get": {
                "description": "This endpoint retrieves thermal data based on query parameters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves thermal data based on query parameters",
                "operationId": "socketThermalOut",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID",
                        "name": "uuid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start Time in RFC3339 format (example: 2024-05-18T18:15:00.000+03:00)",
                        "name": "start_time",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End Time in RFC3339 format (example: 2024-05-18T18:16:00.000+03:00)",
                        "name": "end_time",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Latitude",
                        "name": "latitude",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Longitude",
                        "name": "longitude",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "RSRP",
                        "name": "rsrp",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "RSSI",
                        "name": "rssi",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "RSRQ",
                        "name": "rsrq",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "RSSNR",
                        "name": "rssnr",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "CQI",
                        "name": "cqi",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Bandwidth",
                        "name": "bandwidth",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Cell ID",
                        "name": "cell_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of thermal data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Data"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "No data found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to query or encode data",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/auth": {
            "post": {
                "description": "Производит аутентификацию пользователя на основе предоставленных данных\nПри авторизации без пароля можно произвести авторизацию через Token\nДля этого по аналогии с /api/jwt/verify в поле Authorization нужно разместить ваш значение Bearer: \u003cyour_token\u003e\nEmail всё равно нужно указать для избежания \"призрачных\" аккаунтов, сопоставляется с текущими email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Аутентификация пользователя и выдача токена доступа",
                "operationId": "authenticateUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer your_token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Данные пользователя для аутентификации",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/api.authRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешная аутентификация",
                        "schema": {
                            "$ref": "#/definitions/api.authResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Ошибка аутентификации",
                        "schema": {
                            "$ref": "#/definitions/api.authResponseError"
                        }
                    }
                }
            }
        },
        "/api/user/logout": {
            "get": {
                "description": "This endpoint logs out a user by invalidating their session.",
                "produces": [
                    "application/json"
                ],
                "summary": "Logs out a user",
                "operationId": "userLogout",
                "responses": {
                    "200": {
                        "description": "Successful logout",
                        "schema": {
                            "$ref": "#/definitions/net.Msg"
                        }
                    }
                }
            }
        },
        "/api/user/register": {
            "post": {
                "description": "Производит аутентификацию пользователя на основе предоставленных данных\nПри авторизации без пароля можно произвести авторизацию через Token\nДля этого по аналогии с /api/jwt/verify в поле Authorization нужно разместить ваш значение Bearer: \u003cyour_token\u003e\nEmail всё равно нужно указать для избежания \"призрачных\" аккаунтов, сопоставляется с текущими email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Регистрирует пользователя и выдача токена доступа",
                "operationId": "registrationUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer your_token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "Данные пользователя для аутентификации",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/api.registerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешная аутентификация",
                        "schema": {
                            "$ref": "#/definitions/api.registerResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Ошибка аутентификации",
                        "schema": {
                            "$ref": "#/definitions/api.registerResponseError"
                        }
                    }
                }
            }
        },
        "/api/user/verify": {
            "get": {
                "description": "Проверка валидности переданного ключа верификации пользователя\nДля выдачи такой ссылки в поле key ничего не пишите, поставьте поле \"тестовый режим\" на значение 1 и сделайте отправку\nДля конечной проверки перейдите по тестовой ссылке (НЕ ЯВЛЯЕТСЯ ЧАСТЬЮ API ради забавы разместил)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Верификация почты путём перехода по ссылке",
                "operationId": "verifyUserKey",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ключ верификации пользователя",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 0,
                        "description": "Тестовый режим",
                        "name": "test",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешная проверка ключа",
                        "schema": {
                            "$ref": "#/definitions/net.Msg"
                        }
                    },
                    "400": {
                        "description": "Ошибка при проверке ключа",
                        "schema": {
                            "$ref": "#/definitions/api.linkResponseErr"
                        }
                    },
                    "426": {
                        "description": "Ключ",
                        "schema": {
                            "$ref": "#/definitions/api.linkResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Data": {
            "type": "object",
            "properties": {
                "bandwidth": {
                    "type": "string"
                },
                "cellID": {
                    "type": "string"
                },
                "cqi": {
                    "type": "string"
                },
                "latitude": {
                    "type": "string"
                },
                "longitude": {
                    "type": "string"
                },
                "rsrp": {
                    "type": "string"
                },
                "rsrq": {
                    "type": "string"
                },
                "rssi": {
                    "type": "string"
                },
                "rssnr": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "api.Data_users": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "group": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                },
                "verifed": {
                    "type": "boolean"
                }
            }
        },
        "api.authRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "api.authResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.authResponseSuccess": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "jwt": {
                    "type": "string"
                }
            }
        },
        "api.jwtValidationResult": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                }
            }
        },
        "api.linkResponse": {
            "type": "object",
            "properties": {
                "test_link": {
                    "type": "string"
                }
            }
        },
        "api.linkResponseErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.registerRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "api.registerResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.registerResponseSuccess": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "jwt": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "net.Msg": {
            "type": "object",
            "additionalProperties": true
        },
        "types.Token": {
            "type": "object",
            "properties": {
                "jwt": {
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