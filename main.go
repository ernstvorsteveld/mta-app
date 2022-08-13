package main

import (
	"github.com/ernstvorsteveld/go-mta-app/routes"
	"github.com/ernstvorsteveld/mta-common/common"
	"github.com/ernstvorsteveld/mta-parser-service/mta"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ControlData struct {
	r *gin.Engine
	c chan common.FilenameMessage
}

func main() {
	controlData := ControlData{
		r: gin.Default(),
		c: make(chan common.FilenameMessage),
	}
	controlData.initializeServer()
	controlData.registerRoutes()

	controlData.startListener()
	controlData.run()
}

func (c *ControlData) initializeServer() {
	c.r.Use(cors.Default())
}

func (c *ControlData) registerRoutes() {
	apiV1 := c.r.Group("/api/v1")
	routes.Register(apiV1.Group("/product"))
	routes.RegisterUpload(apiV1.Group("/sta"), c.c)
}

func (c *ControlData) run() {
	c.r.Run(":5000")
}

func (c *ControlData) startListener() {
	mta.Start(c.c)
}
