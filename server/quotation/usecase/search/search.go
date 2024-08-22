package search

import (
	"context"

	"github.com/lemuelZara/server/quotation"
)

type (
	Reader interface {
		GetByCurrencies(ctx context.Context, currencies []string) ([]quotation.Quotation, error)
	}

	SearchService struct {
		reader Reader
	}
)

func NewSearchService(reader Reader) SearchService {
	return SearchService{reader}
}

func (s SearchService) Search(ctx context.Context, currencies []string) ([]quotation.Quotation, error) {
	quotations, err := s.reader.GetByCurrencies(ctx, currencies)
	if err != nil {
		return []quotation.Quotation{}, err
	}

	return quotations, nil
}
