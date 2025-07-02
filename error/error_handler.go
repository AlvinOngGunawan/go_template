package error

import (
	"github.com/labstack/echo/v4"
	"time"
)

type Response struct {
	RequestID string      `json:"request_id"`
	Timestamp string      `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
	Error     interface{} `json:"error,omitempty"`
}

type Error struct {
	Error error
}

func JSON(c echo.Context, status int, data interface{}, err interface{}) error {
	requestID := c.Get("request_id")
	return c.JSON(status, &Response{
		RequestID: requestID.(string),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
		Error:     err,
	})
}
