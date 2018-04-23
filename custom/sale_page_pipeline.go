package custom

import (
	"fmt"
	"os"

	"github.com/essyding/go_spider/core/common/com_interfaces"
	"github.com/essyding/go_spider/core/common/page_items"
)

type SalePagePipelineFile struct {
	pFile *os.File

	path string

	listed map[string]bool
}

func NewSalePagePipelineFile(path string) *SalePagePipelineFile {
	pFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic("File '" + path + "' in PipelineFile open failed.")
	}
	return &SalePagePipelineFile{path: path, pFile: pFile}
}

func (this *SalePagePipelineFile) Process(items *page_items.PageItems, t com_interfaces.Task) {
	println("----------------------------------------------------------------------------------------------\n")
	println("Crawled url :\t" + items.GetRequest().GetUrl() + "\n")
	for _, value := range items.GetAll() {
		switch v := value.(type) {
		case []string:
			println("HHH????Len = ", len(v))
			counter := 0
			for i := 1; i < len(v); i++ {
				println(v[i])
				if !this.listed[v[i]] {
					// this.pFile.WriteString(v[i] + "\n")
					this.listed[v[i]] = true
					counter += 1
					println("written")
				}
			}
			println("Crawed %v properties, %v are newly added to db\n", len(v), counter)
			if err := this.pFile.Sync(); err != nil {
				println("flush error.\n")
			}
		default:
			fmt.Printf("WTF?")
		}
	}
}
