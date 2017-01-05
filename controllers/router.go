package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tudurom/pasta/util"
)

func MakeRoutes(c *gin.Engine) {
	var pc PasteController

	c.POST(util.GlobalConfig.PasteEndpoint, pc.New)
	c.GET(util.GlobalConfig.PasteEndpoint+":id", pc.Show)
	c.PATCH(util.GlobalConfig.PasteEndpoint+":id", pc.Update)
	c.DELETE(util.GlobalConfig.PasteEndpoint+":id", pc.Delete)
}
