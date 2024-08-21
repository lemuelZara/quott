package http

import (
	"strconv"

	"github.com/lemuelZara/server/quotation"
)

type (
	coin struct {
		Code   string `json:"code"`
		CodeIn string `json:"codein"`
		BID    string `json:"bid"`
	}

	result map[string]coin
)

func toQuotations(data result) ([]quotation.Quotation, error) {
	quotations := make([]quotation.Quotation, 0, len(data))
	for _, c := range data {
		q, err := toQuotation(c)
		if err != nil {
			return []quotation.Quotation{}, err
		}
		quotations = append(quotations, q)
	}

	return quotations, nil
}

func toQuotation(data coin) (quotation.Quotation, error) {
	bid, err := strconv.ParseFloat(data.BID, 32)
	if err != nil {
		return quotation.Quotation{}, err
	}

	return quotation.Quotation{
		From: data.Code,
		To:   data.CodeIn,
		BID:  float32(bid),
	}, nil
}
