package db

import (
	"context"

	"gorm.io/gorm"
)

type txKeyType int

const txKey txKeyType = 0

func EnsureTx[T any](ctx context.Context, db *gorm.DB, f func(ctx context.Context, tx *gorm.DB) (T, error)) (T, error) {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if !ok {
		var res T
		err := db.Transaction(func(tx *gorm.DB) error {
			var err error
			res, err = f(context.WithValue(ctx, txKey, tx), tx)
			return err
		})
		return res, err
	}
	return f(ctx, tx)
}
