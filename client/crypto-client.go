package client

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/GhvstCode/cryptocli/model"
)

const (
	baseURL string = "https://www.worldcoinindex.com/apiservice/ticker?key=z4Qdvwrkex1Viug7jwgsHvWKv1Af5R"
)



//CryptoClient is exported...
type CryptoClient struct {
	client  *http.Client
	baseURL string
}

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
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

	URL := "https://api.nomics.com/v1/currencies/ticker?key="+Key+"&interval=1d&ids="+cc+"&convert="+c
	resp, err := hc.client.Get(URL)

	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}

	defer resp.Body.Close()

	var cResp model.Cryptoresponse

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	return cResp.Result(), nil
}