package search

import (
	"context"
	"fmt"

	"github.com/lemuelZara/server/quotation"
)

type (
	Reader interface {
		GetByCurrencies(ctx context.Context, currencies []string) ([]quotation.Quotation, error)
	}

	Writer interface {
		Write(ctx context.Context, q quotation.Quotation) error
	}

	SearchService struct {
		reader Reader
		writer Writer
	}
)

func NewSearchService(reader Reader, writer Writer) SearchService {
	return SearchService{reader, writer}
}

func (s SearchService) Search(ctx context.Context, currencies []string) ([]quotation.Quotation, error) {
	quotations, err := s.reader.GetByCurrencies(ctx, currencies)
	if err != nil {
		return []quotation.Quotation{}, fmt.Errorf("failed on get currencies: %w", err)
	}

	for _, q := range quotations {
		err := s.writer.Write(ctx, q)
		if err != nil {
			return []quotation.Quotation{}, fmt.Errorf("failed on save currency %s-%s: %w", q.From, q.To, err)
		}
	}

	return quotations, nil
}
