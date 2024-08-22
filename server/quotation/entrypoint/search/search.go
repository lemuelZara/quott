package entrypoint

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lemuelZara/server/quotation"
)

type (
	Reader interface {
		Search(ctx context.Context, currencies []string) ([]quotation.Quotation, error)
	}

	Search struct {
		reader Reader
	}
)

func NewSearchHandler(reader Reader) Search {
	return Search{reader}
}

func RegisterEndpoints(handler Search, router *http.ServeMux) {
	router.HandleFunc("/cotacao", handler.Search)
}

func (s Search) Search(w http.ResponseWriter, r *http.Request) {
	quotations, err := s.reader.Search(r.Context(), []string{"USD-EUR"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data, _ := json.Marshal(quotations)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
