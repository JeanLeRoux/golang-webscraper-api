package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type cryptoListResponse struct {
	Data cryptoList `json:"data"`
}

type cryptoList struct {
	CryptoCurrencyList []cryptoDetails `json:"cryptoCurrencyList"`
	TotalCount         string          `json:"totalCount"`
}

type cryptoDetails struct {
	Id                int            `json:"id"`
	Name              string         `json:"name"`
	Symbol            string         `json:"symbol"`
	Slug              string         `json:"slug"`
	CmcRank           int            `json:"cmcRank"`
	MarketPairCount   float64        `json:"marketPairCount"`
	CirculatingSupply float64        `json:"circulatingSupply"`
	TotalSupply       float64        `json:"totalSupply"`
	MaxSupply         float64        `json:"maxSupply"`
	Ath               float64        `json:"ath"`
	Atl               float64        `json:"atl"`
	High24h           float64        `json:"high24h"`
	Low24h            float64        `json:"low24h"`
	IsActive          int            `json:"isActive"`
	LastUpdated       string         `json:"lastUpdated"`
	DateAdded         string         `json:"dateAdded"`
	Quotes            []cryptoQuotes `json:"quotes"`
}

type cryptoQuotes struct {
	Name                     string  `json:"name"`
	Price                    float64 `json:"price"`
	Volume24h                float64 `json:"volume24h"`
	Volume7d                 float64 `json:"volume7d"`
	Volume30d                float64 `json:"volume30d"`
	MarketCap                float64 `json:"marketCap"`
	PercentChange1h          float64 `json:"percentChange1h"`
	PercentChange24h         float64 `json:"percentChange24h"`
	PercentChange7d          float64 `json:"percentChange7d"`
	LastUpdated              float64 `json:"lastUpdated"`
	PercentChange30d         float64 `json:"percentChange30d"`
	PercentChange60d         float64 `json:"percentChange60d"`
	PercentChange90d         float64 `json:"percentChange90d"`
	FullyDilluttedMarketCap  float64 `json:"fullyDilluttedMarketCap"`
	MarketCapByTotalSupply   float64 `json:"marketCapByTotalSupply"`
	Dominance                float64 `json:"dominance"`
	Turnover                 float64 `json:"turnover"`
	YtdPriceChangePercentage float64 `json:"ytdPriceChangePercentage"`
}

type cryptoMetadata struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	CmcRank int    `json:"cmcRank"`
}

func main() {
	router := gin.Default()
	router.GET("/Tech", getLatestTech)
	router.GET("/CryptoMeta", getCryptoMetadata)
	router.Run("localhost:8000")
}

func getCryptoMetadata(ginReturn *gin.Context) {
	resp, err := http.Get("https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing?start=1&limit=100&sortBy=market_cap&sortType=desc&convert=USD&cryptoType=all&tagType=all&audited=false&aux=ath,atl,high24h,low24h,num_market_pairs,cmc_rank,date_added,max_supply,circulating_supply,total_supply,volume_7d,volume_30d")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var cryptoListResp cryptoListResponse
	json.Unmarshal(body, &cryptoListResp)
	var cryptoMeta []cryptoMetadata
	for _, value := range cryptoListResp.Data.CryptoCurrencyList {
		temp := cryptoMetadata{
			Id:      value.Id,
			Name:    value.Name,
			Symbol:  value.Symbol,
			CmcRank: value.CmcRank,
		}
		cryptoMeta = append(cryptoMeta, temp)
	}
	ginReturn.IndentedJSON(http.StatusOK, cryptoMeta)

}

// func getCryptoHistoricData(ginReturn *gin.Context) {
// 	news := []article{}
// 	c := colly.NewCollector()
// 	detailCollector := c.Clone()
// 	c.OnError(func(_ *colly.Response, err error) {
// 		fmt.Println(err)
// 	})
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("visiting", r.URL.String())
// 	})
// 	c.OnHTML(".article-listing", func(e *colly.HTMLElement) {
// 		e.ForEach("article", func(i int, h *colly.HTMLElement) {
// 			url := h.ChildAttr(".tf-article-url", "href")
// 			URL := h.Request.AbsoluteURL(url)
// 			detailCollector.Visit(URL)
// 		})
// 	})
// 	c.Visit("https://www.businessinsider.co.za/tech")
// 	ginReturn.IndentedJSON(http.StatusOK, news)
// }

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
