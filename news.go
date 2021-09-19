package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type article struct {
	ArticleURL string
	Title      string
	ImageUrl   string
	Body       string
}

func getLatestTech(ginReturn *gin.Context) {
	news := []article{}
	c := colly.NewCollector()
	detailCollector := c.Clone()
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println(err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})
	c.OnHTML(".article-listing", func(e *colly.HTMLElement) {
		e.ForEach("article", func(i int, h *colly.HTMLElement) {
			url := h.ChildAttr(".tf-article-url", "href")
			URL := h.Request.AbsoluteURL(url)
			detailCollector.Visit(URL)
		})
	})

	detailCollector.OnHTML("section[id=lhs_column]", func(f *colly.HTMLElement) {
		body := ""
		articleBody := f.DOM.Find(".articleBody")
		articleBody.Find("p").Each(func(i int, s *goquery.Selection) {
			body += s.Text()
		})
		temp := article{
			Title:      f.ChildText("h1"),
			ImageUrl:   f.ChildAttr("img", "src"),
			ArticleURL: f.Request.URL.String(),
			Body:       body,
		}
		news = append(news, temp)
	})
	c.Visit("https://www.businessinsider.co.za/tech")
	ginReturn.IndentedJSON(http.StatusOK, news)
	// file, _ := json.MarshalIndent(news, "", " ")
	// _ = ioutil.WriteFile("news.json", file, 0644)
}
