package controller

import (
	"auto_traveler/helper/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct{
	Message 	string 		`json:"message"`
	Messages	[]string 	`json:"messages,omitempty"`
	Data 		interface{}	`json:"data"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Message = messages.BaseResponseMessageSuccess
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Message = messages.BaseResponseMessageFailed
	response.Messages = []string{err.Error()}

	return c.JSON(status, response)
}
