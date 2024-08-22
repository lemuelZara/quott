package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"
)

type quotation struct {
	BID float32 `json:"bid"`
}

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{
		Timeout: 300 * time.Millisecond,
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if res.StatusCode == http.StatusInternalServerError {
		panic(fmt.Errorf("not possible get quotations: %s", string(data)))
	}

	var result []quotation
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tmpl, err := template.New("Quotation").Parse("Dólar: {{.BID}}\n")
	if err != nil {
		panic(err)
	}

	for _, v := range result {
		err := tmpl.Execute(file, v)
		if err != nil {
			panic(err)
		}
	}
}
