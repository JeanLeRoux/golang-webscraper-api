package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"

	"github.com/gin-gonic/gin"
)

type article struct {
	ArticleURL string
	Title      string
	ImageUrl   string
	Body       string
}

type crypto struct {
	Name              string
	Symbol            string
	PriceUSD          string
	VolumeUSD         string
	MarketCapacityUSD string
	Change1h          string
	Change24h         string
	Change7d          string
}

func main() {
	router := gin.Default()
	router.GET("/Tech", getLatestTech)
	router.GET("/Crypto", getLatestCrypto)
	router.Run("localhost:8000")
}

func getLatestCrypto(ginReturn *gin.Context) {
	cryptos := []crypto{}
	c := colly.NewCollector()
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println(err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})
	c.OnHTML("tbody>tr", func(e *colly.HTMLElement) {
		if len(e.ChildText(".cmc-table__column-name--name")) != 0 {
			temp := crypto{
				Name:              e.ChildText(".cmc-table__column-name--name"),
				Symbol:            e.ChildText(".cmc-table__cell--sort-by__symbol"),
				PriceUSD:          e.ChildText(".cmc-table__cell--sort-by__price"),
				VolumeUSD:         e.ChildText(".cmc-table__cell--sort-by__volume-24-h"),
				MarketCapacityUSD: e.ChildText(".cmc-table__cell--sort-by__market-cap>p>.sc-1ow4cwt-1"),
				Change1h:          e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"),
				Change24h:         e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"),
				Change7d:          e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"),
			}
			cryptos = append(cryptos, temp)
		}
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")
	ginReturn.IndentedJSON(http.StatusOK, cryptos)
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
