package entity

import (
	"context"
)

// TransactionContext use data context
type TransactionContext struct{}

// TransactionContextKey used data context key
var TransactionContextKey = TransactionContext{}

// GetTransactionContextValue get context value
func GetTransactionContextValue(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	v := ctx.Value(TransactionContextKey)
	if v == nil {
		return ""
	}

	return v.(string)
}

// SetTransactionContextValue set context value
func SetTransactionContextValue(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, TransactionContextKey, value)
}
