//Key:z4Qdvwrkex1Viug7jwgsHvWKv1Af5R
package main

import (
	"flag"
	"fmt"
	"cryptocli/client"
	"log"
)

func main () {
	fiatCurrency := flag.String(
		"c", "USD", "The name of the fiat currency you would like to know the price of your crypto in",
	)

	nameOfCrypto := flag.String(
		"cc", "BTC", "Input the name of the CryptoCurrency you would like to know the price of",
	)

	outputType := flag.String(
		"o", "text", "Print output in format: text/json",
	)

	flag.Parse()

	cryptoClient := client.NewCryptoClient()

	crypto, err := cryptoClient.Fetch(*fiatCurrency, *nameOfCrypto)
	if err != nil {
		log.Println(err)
	}

	if *outputType == "json" {
		fmt.Println(crypto.JSON())
	} else {
		fmt.Println(crypto.TextOutput())
	}
}

