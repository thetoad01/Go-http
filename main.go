package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// https://mholt.github.io/json-to-go/
type Response struct {
	ID     int    `json:"id,omitempty"`
	Method string `json:"method,omitempty"`
	Code   int    `json:"code,omitempty"`
	Result Result `json:"result,omitempty"`
}
type Data struct {
	I  string `json:"i,omitempty"`
	H  string `json:"h,omitempty"`
	L  string `json:"l,omitempty"`
	A  string `json:"a,omitempty"`
	V  string `json:"v,omitempty"`
	Vv string `json:"vv,omitempty"`
	C  string `json:"c,omitempty"`
	B  string `json:"b,omitempty"`
	K  string `json:"k,omitempty"`
	T  int64  `json:"t,omitempty"`
}
type Result struct {
	Data []Data `json:"data,omitempty"`
}

// https://www.youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE
func main() {
	resp, err := http.Get("https://api.crypto.com/v2/public/get-ticker?instrument_name=BTC_USDT")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	var apiRes Response

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	json.Unmarshal(bytes, &apiRes)

	fmt.Println(apiRes.Result.Data)

	for _, val := range apiRes.Result.Data {
		fmt.Println(val.I)
	}
}
