package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Routers interface {
	GetReport(ctx *gin.Context)
	SetReport(ctx *gin.Context)
	GetMaxObservationTime(ctx *gin.Context)
	GetRequestsHistory(ctx *gin.Context)
	SaveRequestInfoMiddleware() gin.HandlerFunc
}

func NewRouters(routers Routers) http.Handler {
	r := gin.Default()

	api := r.Group("/api", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
		"Sber":  "123",
		"Test":  "456",
	}))
	v1 := api.Group("/v1")
	v1.GET("/requests_history", routers.GetRequestsHistory)
	v1.Use(routers.SaveRequestInfoMiddleware())
	{
		v1.GET("/get_report/:report_id", routers.GetReport)
		v1.GET("/get_observation_time/:model_id", routers.GetMaxObservationTime)
		v1.POST("/set_report", routers.SetReport)
	}

	return r
}
