package custom

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/essyding/go_spider/core/common/page"
)

type SalePageProcessor struct {
}

func NewSalePageProcessor() *SalePageProcessor {
	return &SalePageProcessor{}
}

func NormalizeSalePath(s string) string {
	if !strings.HasPrefix(s, "https://shanghai.anjuke.com/sale") {
		return ""
	}
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path)
}

func NormalizePropertyPath(s string) string {
	if !strings.HasPrefix(s, "https://shanghai.anjuke.com/prop/view") {
		return ""
	}
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}
	return strings.TrimPrefix(u.Path, "/prop/view/")
}

// Parse html dom here and record the parse result that we want to Page.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this *SalePageProcessor) Process(p *page.Page) {
	query := p.GetHtmlParser()
	var urls []string
	var props []string
	query.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if u := NormalizeSalePath(href); u != "" {
			urls = append(urls, u)
		}
		if p := NormalizePropertyPath(href); p != "" {
			props = append(props, p)
		}
	})
	// these urls will be saved and crawed by other coroutines.
	p.AddTargetRequests(urls, "html")

	if len(props) == 0 {
		fmt.Println("Seems getting captcha??????------------------------")
		query.Find("#verify_page").Each(func(i int, s *goquery.Selection) {
			fmt.Printf("Seeing captcha, sleep 10s")
			time.Sleep(10 * time.Second)
		})
		p.SetSkip(true)
	}
	// the entity we want to save by Pipeline
	p.AddField("props", props)
}

func (this *SalePageProcessor) Finish() {
	fmt.Println("finish")
}
