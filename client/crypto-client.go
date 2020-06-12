package client

import (
	//"io"
	//"os"
	"time"
	//"fmt"
	"net/http"
	//"net"
	"encoding/json"

	"cryptocli/model"
)
const (
	baseURL string = "https://www.worldcoinindex.com/apiservice/ticker?key=z4Qdvwrkex1Viug7jwgsHvWKv1Af5R"
)

// // Currency
// type Currency string

// type CryptoCurrency string

//CryptoClient is exported...
type CryptoClient struct {
	client  *http.Client
	baseURL string
}

//NewCryptoClient is exported ...
func NewCryptoClient() *CryptoClient {
	return &CryptoClient{
		client: &http.Client{
			Transport: nil,
			Timeout:   30000 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (hc *CryptoClient) Fetch(c string , cc string) (model.Result, error) {
	//Url := "https://www.worldcoinindex.com/apiservice/ticker?key=z4Qdvwrkex1Viug7jwgsHvWKv1Af5R&label="+cc+"btc&fiat="+c
	URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids="+cc+"&convert="+c
	resp, err := hc.client.Get(Url)
	if err != nil {
		return model.Result{}, err
	}

	defer resp.Body.Close()

	var cResp model.Cryptoresponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return model.Result{}, err
	}
	return cResp.Result(), nil
}