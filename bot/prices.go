package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getPrices(m string) string {
	type Coin struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
		Code   int    `json:"code"`
	}

	var coin Coin

	crypto := strings.ToUpper(m[1:])
	fmt.Printf("Getting price for %v\n", crypto)
	url := "https://api.binance.com/api/v3/ticker/price?symbol=" + crypto + "USDT"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request to server!")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal([]byte(body), &coin)
	if err != nil {
		log.Println(err)
	}

	r := coin.Price
	if s, err := strconv.ParseFloat(r, 64); err == nil {
		r = "$ " + fmt.Sprintf("%.2f", s)
	}

	if coin.Code == -1121 {
		r = "Invalid cryptocurrency!"
	}

	return r
}
