package file_upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"gopkg.in/mgo.v2/bson"
	"image/jpeg"
	"io"
	"os"
	"path"
)

func Router(r *gin.Engine) {
	r.POST("upload", uploadFile)
	r.POST("uploadIcon", uploadHeaderIcon)
}

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("upload file err :", err)
		c.JSON(200, gin.H{"code": 30001, "msg": "systerm file err", "data": err})
		return
	}

	fullFilename := file.Filename
	filenameWithSuffix := path.Base(fullFilename)
	fileSuffix := path.Ext(filenameWithSuffix)

	fileId := bson.NewObjectId().Hex()

	relFileName := fileId + fileSuffix

	f, err := os.OpenFile("static/f/"+relFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.JSON(200, gin.H{"code": 30001, "msg": "systerm file err", "data": err})
		return
	}
	defer f.Close()

	httpFile, _ := file.Open()
	io.Copy(f, httpFile)

	c.JSON(200, gin.H{"code": 0, "msg": "success", "data": "http://" + c.Request.Host + "/file/" + relFileName})
}

func uploadHeaderIcon(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(200, gin.H{"code": 30001, "msg": "failed", "data": err})
		return
	}

	fileId := bson.NewObjectId().Hex()

	f, err := os.OpenFile("static/images/icon_"+fileId+"_o.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.JSON(200, gin.H{"code": 30001, "msg": "systerm file err", "data": err})
		return
	}

	defer f.Close()

	httpFile, err := file.Open()
	if err != nil {
		c.JSON(200, gin.H{"code": 30001, "msg": "systerm file err", "data": err})
		return
	}

	io.Copy(f, httpFile)

	imageFile, err := os.Open("static/images/icon_" + fileId + "_o.jpg")
	if err != nil {
		c.JSON(200, gin.H{"code": 30001, "msg": "failed", "data": err})
		return
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(imageFile)
	if err != nil {
		c.JSON(200, gin.H{"code": 30001, "msg": "failed", "data": err})
		return
	}

	s := resize.Resize(200, 0, img, resize.Lanczos3)

	out, err := os.Create("static/images/icon_" + fileId + "_s.jpg")
	if err != nil {
		c.JSON(200, gin.H{"code": 30001, "msg": "failed", "data": err})
		return
	}
	defer out.Close()

	jpeg.Encode(out, s, nil)

	b := resize.Resize(500, 0, img, resize.Lanczos3)
	out1, err := os.Create("static/images/icon_" + fileId + "_l.jpg")
	if err != nil {
		c.JSON(200, gin.H{"code": 30001, "msg": "failed", "data": err})
		return
	}
	defer out1.Close()

	jpeg.Encode(out1, b, nil)

	urlPre := "http://" + c.Request.Host + "/image/"
	c.JSON(200, gin.H{"code": 0, "msg": "success", "data": gin.H{"o": urlPre + "icon_" + fileId + "_o.jpg", "s": urlPre + "icon_" + fileId + "_s.jpg", "l": urlPre + "icon_" + fileId + "_l.jpg"}})
}
