package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Println("the HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Json Response")
		fmt.Println(string(data))
	}

	jsonData := map[string]string{"firstname": "Miguel", "lastname": "Silva"}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	//response, _ = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}

	response, err = client.Do(request)
	if err != nil {
		fmt.Println("the HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
