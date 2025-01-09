package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.coindesk.com/v1/bpi/currentprice.json"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Some words %T\n", resp)

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	fmt.Print(content)
}
