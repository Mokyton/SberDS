package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Routers interface {
	GetReport(ctx *gin.Context)
	SetReport(ctx *gin.Context)
}

func NewRouters(routers Routers) http.Handler {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	v1 := r.Group("v1")
	v1.GET("/get_report/:report_id", routers.GetReport)

	v1.POST("/set_report", routers.SetReport)

	return r
}
