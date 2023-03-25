package View

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC Structure",
    "version": "1.0.11"
  },
  "tags": [
    {
      "name": "ChatCSIfElse",
      "description": "Testing of MVC's structure"
    }
  ],
  "paths": {
    "/visitChatCSIfElse": {
      "post": {
        "tags": [
          "ChatCSIfElse"
        ],
        "summary": "visit ChatCSIfElse",
        "requestBody": {
          "description": "visit ChatCSIfElse",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ChatCSIfElseReq"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "visit ChatCSIfElse",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ChatCSIfElseReq"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "ChatCSIfElseReq": {
        "type": "object",
        "properties": {
          "username": {
            "type": "string",
            "example": "Mercedes"
          },
          "visit_date_time": {
            "type": "string",
            "example": "Benz"
          },
          "message_in": {
            "type": "string",
            "format": "date-time",
            "example": "2023-03-25 10:08:00"
          }
        }
      }
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Golang MVC",
	Description:      "63050096_2565_1",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
