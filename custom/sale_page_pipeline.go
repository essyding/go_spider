package custom

import (
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
			counter := 0
			for _, vs := range v {
				if !this.listed[vs] {
					this.pFile.WriteString(vs + "\n")
					this.listed[vs] = true
					counter += 1
				}
			}
			println("Crawed %v properties, %v are newly added to db\n", len(v), counter)
			this.pFile.Sync()
		default:
		}
	}
}
