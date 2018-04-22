package custom

import (
	"fmt"
	"net/url"
	"strings"

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

func ResolveCaptcha(p *page.Page) bool {
	query := p.GetHtmlParser()
	if sel := query.Has("div#verify_page"); sel != nil {
		println("getting captcha...")
		fmt.Printf("req: %v", p.GetHeader())
		return true
	}
	return false
}

// Parse html dom here and record the parse result that we want to Page.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this *SalePageProcessor) Process(p *page.Page) {
	// resolve captcha
	ResolveCaptcha(p)
	query := p.GetHtmlParser()
	var urls []string
	query.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if u := NormalizeSalePath(href); u != "" {
			urls = append(urls, u)
		}
	})
	// these urls will be saved and crawed by other coroutines.
	p.AddTargetRequests(urls, "html")

	var props []string
	query.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if p := NormalizePropertyPath(href); p != "" {
			props = append(props, p)
		}
	})
	if len(props) == 0 {
		println("Seems getting capchaed------------------------\n")
		p.SetSkip(true)
	}
	// the entity we want to save by Pipeline
	p.AddField("props", props)
}

func (this *SalePageProcessor) Finish() {
	fmt.Println("finish")
}
