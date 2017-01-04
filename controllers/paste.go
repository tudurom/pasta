package controllers

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type PasteController struct {
	Controller
}

type PasteForm struct {
	Vanity string `form:"vanity"`
	Token  string `form:"token"`
}

func (pc *PasteController) New(ct *gin.Context) {
	var js PasteForm
	err := ct.Bind(&js)
	if err != nil {
		logrus.Debug(err)
		return
	}
	_, _, err = ct.Request.FormFile("data")
	if err != nil {
		logrus.Debug(err)
		http.Error(ct.Writer, "", http.StatusBadRequest)
		return
	}
	ct.JSON(http.StatusOK, js)
}

func (pc *PasteController) Show(ct *gin.Context) {
	ct.JSON(http.StatusNotImplemented, gin.H{"error": "501 not implemented"})
}

func (pc *PasteController) Update(ct *gin.Context) {
	ct.JSON(http.StatusNotImplemented, gin.H{"error": "501 not implemented"})
}

func (pc *PasteController) Delete(ct *gin.Context) {
	ct.JSON(http.StatusNotImplemented, gin.H{"error": "501 not implemented"})
}
