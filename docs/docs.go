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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/alipay/": {
            "post": {
                "description": "支付宝支付下单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "支付宝支付"
                ],
                "summary": "支付宝支付",
                "parameters": [
                    {
                        "type": "string",
                        "default": "h5",
                        "description": "支付类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "测试",
                        "description": "商品名称",
                        "name": "subject",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "default": 0.01,
                        "description": "金额",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pay.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllers.AlipayData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pay.ResponseVerificationErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pay.ResponseError"
                        }
                    }
                }
            }
        },
        "/wechat/": {
            "post": {
                "description": "微信支付下单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "微信支付"
                ],
                "summary": "微信支付",
                "parameters": [
                    {
                        "type": "string",
                        "default": "h5",
                        "description": "支付类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pay.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pay.ResponseError"
                        }
                    }
                }
            }
        },
        "/wechat/v2/": {
            "post": {
                "description": "微信支付下单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "微信支付"
                ],
                "summary": "微信支付",
                "parameters": [
                    {
                        "type": "string",
                        "default": "h5",
                        "description": "支付类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "测试",
                        "description": "商品名称",
                        "name": "subject",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "default": 0.01,
                        "description": "金额",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/pay.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllers.WechatV2Data"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/pay.ResponseVerificationErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pay.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.AlipayData": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "controllers.WechatV2Data": {
            "type": "object",
            "properties": {
                "pay_sign": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "pay.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "pay.ResponseError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "pay.ResponseVerificationErr": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
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
	Host:        "127.0.0.1:8022",
	BasePath:    "/pay/",
	Schemes:     []string{},
	Title:       "api",
	Description: "api desc",
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
	swag.Register("swagger", &s{})
}
