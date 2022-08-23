package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutchanonc/try-k8s/app/usecases"
)

type Handler struct {
	redisUsecase usecases.RedisUsecase
}

func NewHandler(redisUsecase usecases.RedisUsecase) Handler {
	return Handler{
		redisUsecase: redisUsecase,
	}
}

func (h Handler) HandleRedisGet(c *gin.Context) {
	var q Query
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusInternalServerError, CreateErrorResponse(err))
		return 
		 
	}
	value, err := h.redisUsecase.Get(q.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateErrorResponse(err))
		return 
	}
	c.JSON(http.StatusOK, CreateSuccessResponse(value))
}

func (h Handler) HandleRedisSet(c *gin.Context) {
	var q Query
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusInternalServerError, CreateErrorResponse(err))
		return 
		 
	}
	err := h.redisUsecase.Set(q.Key, q.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateErrorResponse(err))
		return 
	}
	c.JSON(http.StatusOK, CreateSuccessResponse(nil))
}