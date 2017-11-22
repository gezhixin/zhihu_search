package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zhihu_search/service/file_upload"
	"zhihu_search/spider"
)

func main() {

	m1 := make(map[string]interface{})
	spider.HttpGetJson("https://www.zhihu.com", &m1)
	fmt.Println(m1)

	r := gin.Default()

	file_upload.Router(r)

	r.Run(":8080")
}
