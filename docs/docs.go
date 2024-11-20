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
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authenticate",
                "parameters": [
                    {
                        "description": "Authentication data",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.InputLoginDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Token",
                        "schema": {
                            "$ref": "#/definitions/dto.OutputToken"
                        }
                    }
                }
            }
        },
        "/quiz": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quiz"
                ],
                "summary": "Get all quizzes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag ID",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "em progresso",
                        "name": "in_progress",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Quiz"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quiz"
                ],
                "summary": "Create quiz",
                "parameters": [
                    {
                        "description": "Quiz data",
                        "name": "quiz",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateQuizDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Quiz",
                        "schema": {
                            "$ref": "#/definitions/models.Quiz"
                        }
                    }
                }
            }
        },
        "/quiz/answer_check": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quiz"
                ],
                "summary": "Check",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Answer ID",
                        "name": "answer_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Question ID",
                        "name": "question_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/quizController.answerCheckResponse"
                        }
                    }
                }
            }
        },
        "/quiz/finish/{quiz_id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Quiz"
                ],
                "summary": "Finish quiz",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Quiz ID",
                        "name": "quiz_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Success"
                    }
                }
            }
        },
        "/quiz/history": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quiz"
                ],
                "summary": "Get quizzes history",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserCompletedQuiz"
                            }
                        }
                    }
                }
            }
        },
        "/quiz/questions/{quiz_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quiz"
                ],
                "summary": "Get question by Quiz_ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "QUIZ ID",
                        "name": "quiz_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.OutputQuestionDto"
                            }
                        }
                    }
                }
            }
        },
        "/quiz/start/{quiz_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Quiz"
                ],
                "summary": "Start quiz",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Quiz ID",
                        "name": "quiz_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/dto.QuizSession"
                        }
                    }
                }
            }
        },
        "/rank": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rank"
                ],
                "summary": "Get rank",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Update Rank",
                        "name": "update_rank",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Rank"
                            }
                        }
                    }
                }
            }
        },
        "/tag": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    },
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "Get all tags",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/models.Tag"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "Create tag",
                "parameters": [
                    {
                        "description": "Tag data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateTagDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Tag",
                        "schema": {
                            "$ref": "#/definitions/models.Tag"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/user/": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user by ID",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserDto"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Success"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/user/icon": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Upload de ícone de usuário",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Imagem do ícone do usuário (PNG, JPEG, etc.)",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Imagem processada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/icon/{user_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Obter ícone de usuário",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "USER ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ícone do usuário em PNG",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/user/kpi": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user kpi",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/dto.UserKpiDto"
                        }
                    }
                }
            }
        },
        "/user/user_infos": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user info with Token",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateQuizDto": {
            "type": "object",
            "required": [
                "tagId",
                "title"
            ],
            "properties": {
                "tagId": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.CreateTagDto": {
            "type": "object",
            "required": [
                "description"
            ],
            "properties": {
                "description": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUserDto": {
            "type": "object",
            "required": [
                "name",
                "password",
                "username"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "dto.InputLoginDto": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.OutputAnswerDto": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.OutputQuestionDto": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.OutputAnswerDto"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.OutputToken": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.QuizRank": {
            "type": "object",
            "properties": {
                "quizId": {
                    "type": "integer"
                },
                "quizTitle": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                },
                "tagDescription": {
                    "type": "string"
                },
                "tagId": {
                    "type": "integer"
                }
            }
        },
        "dto.QuizSession": {
            "type": "object",
            "properties": {
                "correct": {
                    "type": "integer"
                },
                "questions": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/dto.Status"
                    }
                },
                "quizTitle": {
                    "type": "string"
                },
                "quizzId": {
                    "type": "integer"
                },
                "score": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "dto.Status": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "Unanswered",
                "InCorrect",
                "Correct"
            ]
        },
        "dto.TagRank": {
            "type": "object",
            "properties": {
                "tagDescription": {
                    "type": "string"
                },
                "tagId": {
                    "type": "integer"
                },
                "totalscore": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateUserDto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UserKpiDto": {
            "type": "object",
            "properties": {
                "quizzesRank": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.QuizRank"
                    }
                },
                "tagsRank": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.TagRank"
                    }
                },
                "totaScore": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Quiz": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "tag": {
                    "$ref": "#/definitions/models.Tag"
                },
                "tagId": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Rank": {
            "type": "object",
            "properties": {
                "bestScoreQuiz": {
                    "$ref": "#/definitions/models.Quiz"
                },
                "bestScoreQuizId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "totalScore": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Tag": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "description": "User object",
            "type": "object",
            "required": [
                "name",
                "password",
                "username"
            ],
            "properties": {
                "iconPath": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "models.UserCompletedQuiz": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "quiz": {
                    "$ref": "#/definitions/models.Quiz"
                },
                "quizId": {
                    "type": "integer"
                },
                "score": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "quizController.answerCheckResponse": {
            "type": "object",
            "properties": {
                "is_correct": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Insira o token de autenticação no formato: \"Bearer {token}\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Mensina",
	Description:      "API desenvolvida para projeto academico",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
