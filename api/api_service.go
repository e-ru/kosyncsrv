package api

import "github.com/gin-gonic/gin"

type ApiService interface {
	registerUser(c *gin.Context)
}
