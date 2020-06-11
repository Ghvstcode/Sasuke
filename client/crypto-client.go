package client

import (
	"net/http"
	"encoding/json"

	"cryptocli/model"
)
const (
	BaseURL string = "https://www.worldcoinindex.com/apiservice/ticker?key=z4Qdvwrkex1Viug7jwgsHvWKv1Af5R"
)

type Currency string

type CryptoCurrency string

type CryptoClient struct {
	client  *http.Client
	baseURL string
}

func NewCryptoClient() *CryptoClient {
	return &CryptoClient{
		client: &http.Client{
			Transport: nil,
		},
		baseURL: BaseURL,
	}
}

func (hc *CryptoClient) Fetch(c string , cc string) (model.Result, error) {
	Url := "https://www.worldcoinindex.com/apiservice/ticker?key=z4Qdvwrkex1Viug7jwgsHvWKv1Af5R&label="+cc+"btc&fiat="+c
	resp, err := hc.client.Get(Url)
	if err != nil {
		return model.Result{}, err
	}
	defer resp.Body.Close()

	var cResp model.Markets
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return model.Result{}, err
	}

	return cResp.Result(), nil
}