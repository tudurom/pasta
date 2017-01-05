package controllers

import (
	"bytes"
	"hash/adler32"
	"io"
	"mime"
	"net/http"
	"path/filepath"

	filetype "gopkg.in/h2non/filetype.v1"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tudurom/pasta/models"
	"github.com/tudurom/pasta/util"
)

type PasteController struct {
	Controller
}

type PasteForm struct {
	Vanity string `form:"vanity"`
}

func createPaste(filename string, file bytes.Buffer) (models.Paste, error) {
	var p models.Paste

	p.Filename = filename
	p.UUID = uuid.New().String()
	p.Hash = adler32.Checksum(file.Bytes())
	filekind, err := filetype.Match(file.Bytes())
	logrus.Debug(filepath.Ext(filename))
	if err != nil || filekind.Extension == "unknown" {
		p.MimeType = mime.TypeByExtension(filepath.Ext(filename))
		if p.MimeType == "" {
			p.MimeType = "application/octet-stream"
		}
	} else {
		p.MimeType = filekind.MIME.Value
	}

	return p, nil
}

func (pc *PasteController) New(ct *gin.Context) {
	var pf PasteForm
	err := ct.Bind(&pf)
	if err != nil {
		logrus.Debug(err)
		return
	}
	file, _, err := ct.Request.FormFile("data")
	if err != nil {
		logrus.Debug(err)
		http.Error(ct.Writer, "", http.StatusBadRequest)
		return
	}
	defer file.Close()

	var p models.Paste
	nf := util.GlobalDB.First(&p, "filename = ?", pf.Vanity).RecordNotFound()
	if nf {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, file)
		if err != nil {
			logrus.Debug(err)
			http.Error(ct.Writer, "", http.StatusInternalServerError)
		}
		p, err = createPaste(pf.Vanity, buf)
		util.GlobalDB.Create(&p)
		logrus.Debug("Created paste", p)
	}
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
