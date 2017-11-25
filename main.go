package main

import (
	// "fmt"
	// "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	// "log"
	"zhihu_search/service/file_upload"
	"zhihu_search/service/user"
	// "zhihu_search/spider"
)

func main() {

	r := gin.Default()

	file_upload.Router(r)
	user.Router(r)

	r.Run(":8080")
}
