package main

import (
	"github.com/gin-gonic/gin"
	"zhihu_search/service/file_upload"
)

func main() {

	r := gin.Default()

	r.Static("/image", "static/images")
	r.Static("/file", "static/f")

	file_upload.Router(r)

	r.Run(":8080")
}
