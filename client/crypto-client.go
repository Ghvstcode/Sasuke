package client

import (
	"time"
	"net/http"
	"encoding/json"

	"cryptocli/model"
)
const (
	baseURL string = "https://www.worldcoinindex.com/apiservice/ticker?key=z4Qdvwrkex1Viug7jwgsHvWKv1Af5R"
)



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

//Fetch is exported ...
func (hc *CryptoClient) Fetch(c string , cc string) (model.Result, error) {
	URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids="+cc+"&convert="+c
	resp, err := hc.client.Get(URL)

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