package main

import (
	"wiki-server/service"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/wiki", service.ProcessGetWikis)
	r.GET("/wiki/:uuid", service.ProcessGetWiki)
	r.POST("/wiki", service.ProcessCreateWiki)
	r.PUT("/wiki/:uuid", service.ProcessUpdateWiki)
	r.DELETE("/wiki/:uuid", service.ProcessDeleteWiki)
	return r
}
