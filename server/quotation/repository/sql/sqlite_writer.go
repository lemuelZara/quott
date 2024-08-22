package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/lemuelZara/server/quotation"
)

const (
	fields       = "code, code_in, bid"
	upsertParams = ":code, :code_in, :bid"
	upsertSQL    = "REPLACE INTO quotation (" + fields + ") VALUES (" + upsertParams + ")"
)

type WriteSQLite struct {
	db *sql.DB
}

func NewWriteSQLite(db *sql.DB) WriteSQLite {
	return WriteSQLite{db}
}

func (w WriteSQLite) Write(ctx context.Context, q quotation.Quotation) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	_, err := w.db.ExecContext(ctx, upsertSQL, addUpsertParams(q)...)
	if err != nil {
		return err
	}

	return nil
}

func addUpsertParams(q quotation.Quotation) []any {
	return []any{
		sql.Named("code", q.From),
		sql.Named("code_in", q.To),
		sql.Named("bid", q.BID),
	}
}
