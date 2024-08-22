package main

import (
	"net/http"

	"github.com/lemuelZara/server/internal/database"
	"github.com/lemuelZara/server/internal/server"
	entrypoint "github.com/lemuelZara/server/quotation/entrypoint/search"
	quotationHttp "github.com/lemuelZara/server/quotation/repository/http"
	"github.com/lemuelZara/server/quotation/repository/sql"
	"github.com/lemuelZara/server/quotation/usecase/search"
)

func main() {
	// Internal
	router := http.NewServeMux()
	app := server.NewWebApplication(router)

	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	// Quotation
	reader := quotationHttp.NewClient()
	writer := sql.NewWriteSQLite(db)
	searchService := search.NewSearchService(reader, writer)
	handler := entrypoint.NewSearchHandler(searchService)
	entrypoint.RegisterEndpoints(handler, router)

	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
