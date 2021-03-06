// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "JJ",
            "url": "http://jianjun.fun",
            "email": "jianjun.xie@aliyun.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/category": {
            "post": {
                "description": "分类列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "获取分类列表",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PageListParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.CategoryRes"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.SignupParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    }
                }
            }
        },
        "/posts": {
            "post": {
                "description": "添加文章",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "添加文章",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddPostParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Res"
                        }
                    }
                }
            }
        },
        "/posts/:id": {
            "get": {
                "description": "文章详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "获取文章详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "参数",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Posts"
                        }
                    }
                }
            },
            "put": {
                "description": "更新文章",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "更新文章",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddPostParam"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "pid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Res"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除文章",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "删除文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "参数",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Res"
                        }
                    }
                }
            }
        },
        "/posts/list": {
            "post": {
                "description": "文章列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "获取文章列表",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PageListParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.PostsRes"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "注册接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.SignupParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":123,\"msg\":\"\"}",
                        "schema": {
                            "$ref": "#/definitions/controllers.Res"
                        }
                    }
                }
            }
        },
        "/user/:id": {
            "get": {
                "description": "用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "根据用户id获取详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "参数",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserModal"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "description": "用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "根据用户token获取详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserModal"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CategoryRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Category"
                    }
                },
                "msg": {
                    "type": "string"
                },
                "pageIndex": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "controllers.PostsRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Posts"
                    }
                },
                "msg": {
                    "type": "string"
                },
                "pageIndex": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "controllers.Res": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "models.AddPostParam": {
            "type": "object",
            "required": [
                "cid",
                "content",
                "title"
            ],
            "properties": {
                "cid": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "cid": {
                    "type": "integer"
                },
                "cname": {
                    "type": "string"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "createAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                },
                "updateAt": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "models.PageListParams": {
            "type": "object",
            "properties": {
                "asc": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "condition": {
                    "type": "object",
                    "additionalProperties": true
                },
                "desc": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pageIndex": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "models.Posts": {
            "type": "object",
            "properties": {
                "authId": {
                    "type": "integer"
                },
                "authName": {
                    "type": "string"
                },
                "cid": {
                    "type": "integer"
                },
                "cname": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "createAt": {
                    "type": "string"
                },
                "pid": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updateAt": {
                    "type": "string"
                }
            }
        },
        "models.SignupParams": {
            "type": "object",
            "required": [
                "passwd",
                "userName"
            ],
            "properties": {
                "passwd": {
                    "type": "string"
                },
                "repasswd": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "models.UserModal": {
            "type": "object",
            "properties": {
                "createAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                },
                "updateAt": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "new-bell",
	Description: "go web 练习",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
