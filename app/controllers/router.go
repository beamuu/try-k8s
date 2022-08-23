package controllers

import "github.com/gin-gonic/gin"

func NewAppRouter(h Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/get", h.HandleRedisGet)
	router.POST("/set", h.HandleRedisSet)
	return router
}