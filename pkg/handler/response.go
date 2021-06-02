package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatalln(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
