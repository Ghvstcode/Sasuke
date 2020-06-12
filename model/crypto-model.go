package model

import (
	"fmt"
	"encoding/json"
	"time"
    //"strconv"
)
// Cryptoresponse is
type Cryptoresponse []struct {
	ID                string    `json:"id"`
	Currency          string    `json:"currency"`
	Symbol            string    `json:"symbol"`
	Name              string    `json:"name"`
	LogoURL           string    `json:"logo_url"`
	Price             string    `json:"price"`
	PriceDate         time.Time `json:"price_date"`
	PriceTimestamp    time.Time `json:"price_timestamp"`
	CirculatingSupply string    `json:"circulating_supply"`
	MarketCap         string    `json:"market_cap"`
	Rank              string    `json:"rank"`
	High              string    `json:"high"`
	HighTimestamp     time.Time `json:"high_timestamp"`
	OneD              struct {
		Volume             string `json:"volume"`
		PriceChange        string `json:"price_change"`
		PriceChangePct     string `json:"price_change_pct"`
		VolumeChange       string `json:"volume_change"`
		VolumeChangePct    string `json:"volume_change_pct"`
		MarketCapChange    string `json:"market_cap_change"`
		MarketCapChangePct string `json:"market_cap_change_pct"`
	} `json:"1d"`
}




// Result is ...
type Result struct {
	Name      string  `json:"Name"`
	Price     string `json:"Price"`
	Rank      string    `json:"rank"`
	High      string    `json:"high"`
	CirculatingSupply string    `json:"circulating_supply"`
	Date      string   `json:"Timestamp"`
}

//FormattedDate returns string
func (c Cryptoresponse) FormattedDate() string {
	return c[0].PriceTimestamp.String()
}

//JSON returns string
func (r Result) JSON() string {
	rJSON, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(rJSON)
}

//Result is exported ...
func (c Cryptoresponse) Result() Result {
	return Result{   
		Name: c[0].Name,      
		Price: string(c[0].Price),      
		Rank: c[0].Rank,
		High: c[0].High,
		CirculatingSupply: c[0].CirculatingSupply,    
		Date: c.FormattedDate(),      
	}
}

//TextOutput is exported ...
func (r Result) TextOutput() string {
	p := fmt.Sprintf(
		"Name: %s\nPrice : %s\nRank: %s\nHigh: %s\nCirculatingSupply: %s\nDate: %s\n",
		 r.Name, r.Price, r.Rank, r.High, r.CirculatingSupply, r.Date)
	return p
}