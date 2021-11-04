package main

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
	Slug    string `json:"slug"`
	Symbol  string `json:"symbol"`
	CmcRank int    `json:"cmcRank"`
	IconUrl string `json:"iconUrl"`
}

type cryptoNewsResponse struct {
	Data []cryptoNews `json:"data"`
}

type cryptoNews struct {
	Cover string         `json:"cover"`
	Meta  cryptoNewsMeta `json:"meta"`
}

type cryptoNewsMeta struct {
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	SourceUrl string `json:"sourceUrl"`
	CreatedAt string `json:"createdAt"`
}

type cryptoGraphResponse struct {
	Data graphData `json:"data"`
}

type graphData struct {
	Points map[string]interface{} `json:"points"`
}
