package controllers

import (
	"SberDS/internal/models"
	"github.com/gin-gonic/gin"
	"log"
)

type MiddlewareController struct {
	*BaseController
}

func NewMiddlewareController(base *BaseController) *MiddlewareController {
	return &MiddlewareController{base}
}

func (c *MiddlewareController) SaveRequestInfoMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto models.RequestSnippet
		user := ctx.MustGet(gin.AuthUserKey).(string)
		dto.UserName = user
		dto.Endpoint = ctx.Request.RequestURI

		err := c.store.AddSnippetToRequestHistory(ctx, dto)
		if err != nil {
			log.Printf("can't add request snippet to request histort: %v", err)
		}
	}
}
