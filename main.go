package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"log"
	"zhihu_search/service/file_upload"
	"zhihu_search/spider"
)

func main() {

	doc, err := spider.GetHtmlDoc("https://www.zhihu.com/people/ge-zhi-xin-49/logs") //("https://www.zhihu.com/people/ge-zhi-xin-49/activities")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".ProfileHeader-info").Each(func(i int, s *goquery.Selection) {
		band := s.Find(".ProfileHeader-detailLabel").Text()
		title := s.Find(".ProfileHeader-detailValue").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
		fmt.Println(s.Text())
	})

	// spider.Start()

	r := gin.Default()

	file_upload.Router(r)

	r.Run(":8080")
}
