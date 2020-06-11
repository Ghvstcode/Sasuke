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

func (hc *CryptoClient) Fetch(c Currency , cc CryptoCurrency) (model.Comic, error) {
	Url := "https://www.worldcoinindex.com/apiservice/ticker?key=z4Qdvwrkex1Viug7jwgsHvWKv1Af5R&label=",cc,"btc&fiat=",c
	resp, err := hc.client.Get(hc.buildURL(n))
	if err != nil {
		return model.Comic{}, err
	}
	defer resp.Body.Close()

	var comicResp model.ComicResponse
	if err := json.NewDecoder(resp.Body).Decode(&comicResp); err != nil {
		return model.Comic{}, err
	}

	if save {
		if err := hc.SaveToDisk(comicResp.Img, "."); err != nil {
			fmt.Println("Failed to save image!")
		}
	}
	return comicResp.Comic(), nil
}