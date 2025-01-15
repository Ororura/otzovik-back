package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *UserHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/users", handler.GetUsers)
	return router
}
