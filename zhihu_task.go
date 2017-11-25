package main

import (
	"time"
	"zhihu_search/spider"
)

func main() {
	spider.Start()

	for {
		time.Sleep(time.Second * time.Duration(3000000))
	}
}
