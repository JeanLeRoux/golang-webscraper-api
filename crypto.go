package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

var cryptoListUrl = "https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing?start=1&limit=100&sortBy=market_cap&sortType=desc&convert=USD&cryptoType=all&tagType=all&audited=false&aux=ath,atl,high24h,low24h,num_market_pairs,cmc_rank,date_added,max_supply,circulating_supply,total_supply,volume_7d,volume_30d"

func getCryptoMetadata(ginReturn *gin.Context) {
	ginReturn.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	resp, err := http.Get(cryptoListUrl)
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
			Slug:    value.Slug,
			Symbol:  value.Symbol,
			CmcRank: value.CmcRank,
			IconUrl: fmt.Sprintf("https://s2.coinmarketcap.com/static/img/coins/128x128/%s.png", strconv.Itoa(value.Id)),
		}
		cryptoMeta = append(cryptoMeta, temp)
	}
	ginReturn.IndentedJSON(http.StatusOK, cryptoMeta)

}

func getCryptoNews(ginReturn *gin.Context) {
	ginReturn.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	url := fmt.Sprintf("https://api.coinmarketcap.com/data-api/v3/headlines/coinPage/news/slug?slug=%s&size=5&page=1", ginReturn.Query("crypto"))
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var cryptoNewsResp cryptoNewsResponse
	json.Unmarshal(body, &cryptoNewsResp)
	for index, value := range cryptoNewsResp.Data {
		if value.Cover == "" {
			cryptoNewsResp.Data[index].Cover = "https://s2.coinmarketcap.com/static/cloud/img/news/placeholder1.jpg"
		}
	}
	ginReturn.IndentedJSON(http.StatusOK, cryptoNewsResp)

}

func getCryptoChartData(ginReturn *gin.Context) {
	ginReturn.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	url := fmt.Sprintf("https://api.coinmarketcap.com/data-api/v3/cryptocurrency/detail/chart?id=%s&range=ALL", ginReturn.Query("crypto"))
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var cryptoGraphResp cryptoGraphResponse
	json.Unmarshal(body, &cryptoGraphResp)
	graphChartPoint := [][]int{}
	for pointKey, pointValues := range cryptoGraphResp.Data.Points {
		prices := pointValues.(map[string]interface{})["v"].([]interface{})
		newPointKey, _ := strconv.Atoi(pointKey + "000")

		temp := []int{
			newPointKey,
		}

		for k, v := range prices {
			if k == 0 || k == 1 {
				val, _ := v.(float64)
				temp = append(temp, int(val))
			}
		}

		graphChartPoint = append(graphChartPoint, temp)

	}
	sort.Slice(graphChartPoint[:], func(i, j int) bool {
		for x := range graphChartPoint[i] {
			if graphChartPoint[i][x] == graphChartPoint[j][x] {
				continue
			}
			return graphChartPoint[i][x] < graphChartPoint[j][x]
		}
		return false
	})
	ginReturn.IndentedJSON(http.StatusOK, graphChartPoint)

}

func getCryptoDetails(ginReturn *gin.Context) {
	ginReturn.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	resp, err := http.Get(cryptoListUrl)
	cryptoId, _ := strconv.Atoi(ginReturn.Query("crypto"))
	fmt.Println(cryptoId)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var cryptoListResp cryptoListResponse
	json.Unmarshal(body, &cryptoListResp)
	for _, value := range cryptoListResp.Data.CryptoCurrencyList {
		if value.Id == cryptoId {
			ginReturn.IndentedJSON(http.StatusOK, value)
		}
	}

}
