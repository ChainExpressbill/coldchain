// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
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
        "/dashboard/summary/charts/{type}": {
            "get": {
                "description": "최근 30일 간 데이터",
                "consumes": [
                    "application/json;charset=UTF-8"
                ],
                "produces": [
                    "application/json;charset=UTF-8"
                ],
                "tags": [
                    "Dashboard"
                ],
                "summary": "최근 30일 간 데이터",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chart Type",
                        "name": "type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object",
                                "properties": {
                                    "count": {
                                        "type": "integer"
                                    },
                                    "orderer": {
                                        "type": "string"
                                    },
                                    "timestamp": {
                                        "type": "string"
                                    }
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/dashboard/summary/last-month": {
            "get": {
                "description": "지난 달 주문 요약",
                "consumes": [
                    "application/json;charset=UTF-8"
                ],
                "produces": [
                    "application/json;charset=UTF-8"
                ],
                "tags": [
                    "Dashboard"
                ],
                "summary": "지난 달 주문 요약",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dashboard.SummaryCountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/dashboard/summary/today": {
            "get": {
                "description": "금일 주문 현황",
                "consumes": [
                    "application/json;charset=UTF-8"
                ],
                "produces": [
                    "application/json;charset=UTF-8"
                ],
                "tags": [
                    "Dashboard"
                ],
                "summary": "금일 주문 현황",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dashboard.SummaryCountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/join": {
            "post": {
                "description": "Chain expressbill 관리자 회원가입",
                "consumes": [
                    "application/json;charset=UTF-8"
                ],
                "produces": [
                    "application/json;charset=UTF-8"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Chain expressbill 관리자 회원가입",
                "parameters": [
                    {
                        "description": "id, password, name, emailAddress",
                        "name": "JoinParams",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.JoinParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Chain expressbill 관리자 로그인",
                "consumes": [
                    "application/json;charset=UTF-8"
                ],
                "produces": [
                    "application/json;charset=UTF-8"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Chain expressbill 관리자 로그인",
                "parameters": [
                    {
                        "description": "id, password",
                        "name": "loginParams",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.LoginParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Account"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Chain expressbill 관리자 로그아웃",
                "consumes": [
                    "application/json;charset=UTF-8"
                ],
                "produces": [
                    "application/json;charset=UTF-8"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Chain expressbill 관리자 로그아웃",
                "parameters": [
                    {
                        "description": "id",
                        "name": "logoutParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.LoginParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "account.JoinParams": {
            "type": "object",
            "properties": {
                "emailAddress": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "account.LoginParams": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dashboard.SummaryCountResponse": {
            "type": "object",
            "properties": {
                "orderCount": {
                    "type": "integer"
                },
                "ordererCount": {
                    "type": "integer"
                }
            }
        },
        "ent.Account": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "CreatedAt holds the value of the \"createdAt\" field.\n생성 시간",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AccountQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.AccountEdges"
                },
                "emailAddress": {
                    "description": "EmailAddress holds the value of the \"email_address\" field.\n이메일 주소",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.\n계정 id",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.\n관리자 이름",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "UpdatedAt holds the value of the \"updatedAt\" field.\n수정 시간",
                    "type": "string"
                }
            }
        },
        "ent.AccountEdges": {
            "type": "object",
            "properties": {
                "orders": {
                    "description": "Orders holds the value of the orders edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Order"
                    }
                }
            }
        },
        "ent.Order": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "CreatedAt holds the value of the \"createdAt\" field.\n생성 시간",
                    "type": "string"
                },
                "drugName": {
                    "description": "DrugName holds the value of the \"drug_name\" field.\n약품명",
                    "type": "string"
                },
                "drugStandard": {
                    "description": "DrugStandard holds the value of the \"drug_standard\" field.\n규격",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the OrderQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.OrderEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "oid": {
                    "description": "Oid holds the value of the \"oid\" field.\n주문 ID",
                    "type": "string"
                },
                "orderer": {
                    "description": "Orderer holds the value of the \"orderer\" field.\n주문자(제약업체. 위탁배송을 요청한 업체)",
                    "type": "string"
                },
                "quantity": {
                    "description": "Quantity holds the value of the \"quantity\" field.\n수량",
                    "type": "integer"
                },
                "receiver": {
                    "description": "Receiver holds the value of the \"receiver\" field.\n수령자 (주문 요청 업체(약국, 병원 등))",
                    "type": "string"
                },
                "registerName": {
                    "description": "RegisterName holds the value of the \"register_name\" field.\n주문 등록자",
                    "type": "string"
                },
                "storageCondition": {
                    "description": "StorageCondition holds the value of the \"storage_condition\" field.\n보관조건",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "UpdatedAt holds the value of the \"updatedAt\" field.\n수정 시간",
                    "type": "string"
                }
            }
        },
        "ent.OrderEdges": {
            "type": "object",
            "properties": {
                "manager": {
                    "description": "Manager holds the value of the manager edge.",
                    "$ref": "#/definitions/ent.Account"
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
	Title:            "Coldchain API",
	Description:      "This is coldchain app by chain-expressbill",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
