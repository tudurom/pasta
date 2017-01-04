package controllers

import "github.com/gin-gonic/gin"

type Controller struct {
}

type ControllerInterface interface {
	New(ct *gin.Context)
	Show(ct *gin.Context)
	Update(ct *gin.Context)
	Delete(ct *gin.Context)
}
