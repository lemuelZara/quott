package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/lemuelZara/server/quotation"
)

type Client struct {
	url     string
	timeout time.Duration
}

func NewClient() Client {
	return Client{
		url:     "https://economia.awesomeapi.com.br/last/",
		timeout: 200 * time.Millisecond,
	}
}

func (c Client) GetByCurrencies(ctx context.Context, currencies []string) ([]quotation.Quotation, error) {
	params := strings.Join(currencies, ",")
	url := fmt.Sprintf(c.url + params)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []quotation.Quotation{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: c.timeout}
	res, err := client.Do(req)
	if err != nil {
		return []quotation.Quotation{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return []quotation.Quotation{}, fmt.Errorf("%w %s", quotation.ErrCoinNotExists, params)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []quotation.Quotation{}, err
	}

	var r result
	err = json.Unmarshal(data, &r)
	if err != nil {
		return []quotation.Quotation{}, err
	}

	return toQuotations(r)
}
