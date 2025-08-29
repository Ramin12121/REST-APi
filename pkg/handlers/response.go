package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c echo.Context, statusCode int, message string) error {
	logrus.Error(message)
	return c.JSON(statusCode, errorResponse{Message: message})
}
