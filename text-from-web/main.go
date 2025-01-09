package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.coindesk.com/v1/bpi/currentprice.json"

type Response struct {
	Bpi map[string]BitcoinPrice `json:"bpi"`
}

type BitcoinPrice struct {
	Code string `json:"code"`
	Rate string `json:"rate"`
}

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		panic(err)
	}

	for _, bitPrice := range response.Bpi {
		fmt.Printf("The price for %v: %v\n", bitPrice.Code, bitPrice.Rate)
	}
}
