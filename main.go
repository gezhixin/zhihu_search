package main

import (
	// "fmt"
	// "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	// "log"
	"zhihu_search/service/file_upload"
	"zhihu_search/spider"
)

func main() {

	spider.Start()

	r := gin.Default()

	file_upload.Router(r)

	r.Run(":8080")
}
