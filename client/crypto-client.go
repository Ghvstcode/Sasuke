package client

import (
	"encoding/json"
	"fmt"
	l "github.com/Ghvstcode/Sasuke/logger"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ghvstcode/Sasuke/model"
)

const (
	baseURL string = "https://api.nomics.com/v1/currencies/ticker?"
)



//CryptoClient a struct containing a HTTP client & a base URL
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
	Key, exists := os.LookupEnv("API_KEY")

	if !exists {
		log.Print("Env Variable not found")
	}

	URL := baseURL+"key="+Key+"&interval=1d&ids="+cc+"&convert="+c
	resp, err := hc.client.Get(URL)

	if err != nil ||resp.StatusCode != 200 {
		l.ErrorLogger.Fatalln("An Error occurred, please try again")
	}

	defer resp.Body.Close()

	var cResp model.Cryptoresponse

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		fmt.Println("Error", err)
		l.ErrorLogger.Println("An Internal error occurred, unable to retrieve cryptocurrency")
	}


	return cResp.Result(), nil
}