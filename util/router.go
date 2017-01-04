package util

import (
	"github.com/gin-gonic/gin"
	"github.com/tudurom/pasta/controllers"
)

func MakeRoutes(c *gin.Engine) {
	var pc controllers.PasteController

	c.POST(GlobalConfig.PasteEndpoint, pc.New)
	c.GET(GlobalConfig.PasteEndpoint+":id", pc.Show)
	c.PATCH(GlobalConfig.PasteEndpoint+":id", pc.Update)
	c.DELETE(GlobalConfig.PasteEndpoint+":id", pc.Delete)
}
