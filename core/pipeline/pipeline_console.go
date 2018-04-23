package pipeline

import (
	"sync"

	"github.com/essyding/go_spider/core/common/com_interfaces"
	"github.com/essyding/go_spider/core/common/page_items"
)

type PipelineConsole struct {
	m sync.Map
}

func NewPipelineConsole() *PipelineConsole {
	return &PipelineConsole{
		m: sync.Map{},
	}
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
				if _, ok := this.m.Load(vs); !ok {
					println(key + "\t:\t" + vs)
					this.m.Store(vs, true)
				}
			}
		default:
			println(key + "\t:\tunsupported type.")
		}
	}
}
