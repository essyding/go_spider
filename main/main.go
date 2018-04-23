// Package main is the crawler entry.
package main

import (
	"github.com/essyding/go_spider/core/pipeline"
	"github.com/essyding/go_spider/core/scheduler"
	"github.com/essyding/go_spider/core/spider"
	"github.com/essyding/go_spider/custom"
)

func main() {
	spider.NewSpider(custom.NewSalePageProcessor(), "InitalProperty").
		SetSleepTime("rand", 1000, 3000).
		SetScheduler(scheduler.NewQueueScheduler(true)).                                       // remove duplicate url.
		AddUrlWithHeaderFile("http://shanghai.anjuke.com/sale/pudong", "html", "headers.txt"). // start url, html is the responce type ("html" or "json")
		AddPipeline(pipeline.NewPipelineConsole()).                                            // print result on screen
		SetThreadnum(3).                                                                       // crawl request by three Coroutines
		Run()
}
