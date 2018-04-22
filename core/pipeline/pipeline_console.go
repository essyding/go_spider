package pipeline

import (
	"github.com/essyding/go_spider/core/common/com_interfaces"
	"github.com/essyding/go_spider/core/common/page_items"
)

type PipelineConsole struct {
}

func NewPipelineConsole() *PipelineConsole {
	return &PipelineConsole{}
}

func (this *PipelineConsole) Process(items *page_items.PageItems, t com_interfaces.Task) {
	println("----------------------------------------------------------------------------------------------")
	println("Crawled url :\t" + items.GetRequest().GetUrl() + "\n")
	println("Crawled result : ")
	for key, value := range items.GetAll() {
		switch v := value.(type) {
		case string:
			println(key + "\t:\t" + v)
		case []string:
			for _, vs := range v {
				println(key + "\t:\t" + vs)
			}
		default:
			println(key + "\t:\tunsupported type.")
		}
	}
}
