package model

type cryptoresponse struct {
	Markets []struct {
		Label     string  `json:"Label"`
		Name      string  `json:"Name"`
		Price     float64 `json:"Price"`
		Volume24H float64 `json:"Volume_24h"`
		Timestamp int     `json:"Timestamp"`
	} `json:"Markets"`
}

type result struct {
	Label     string  `json:"Label"`
	Name      string  `json:"Name"`
	Price     float64 `json:"Price"`
	Volume    float64 `json:"Volume_24h"`
	Date      string   `json:"Timestamp"`
}