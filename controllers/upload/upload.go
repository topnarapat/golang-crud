package uploadcontroller

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vincent-petithory/dataurl"
)

func UploadImage(c *gin.Context) {

	var input InputUpload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b64data := input.ImageName[strings.IndexByte(input.ImageName, ',')+1:]
	ext, _ := dataurl.DecodeString(input.ImageName) // find file type
	data, _ := base64.StdEncoding.DecodeString(b64data)
	r := string(data)
	newFilename := uuid.New().String()
	var extFilename string = ""
	if ext.ContentType() == "image/png" {
		extFilename = ".png"
		ioutil.WriteFile("./public/images/"+newFilename+".png", []byte(r), 0644)
	} else if ext.ContentType() == "image/jpeg" {
		extFilename = ".jpg"
		ioutil.WriteFile("./public/images/"+newFilename+".jpg", []byte(r), 0644)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Upload successful",
		"image_url": "http://localhost:3001/api/v1/public/images/" + newFilename + extFilename,
	})

}
