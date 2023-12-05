package repository

import (
	"context"

	"github.com/alvinfebriando/gin-gorm-skeleton/valueobject"
	"gorm.io/gorm"
)

const defaultLimit = 10

func getPagination(q *valueobject.Query) (int, int) {
	limit := q.Limit
	if limit <= 0 {
		limit = defaultLimit
	}
	page := q.Page
	if page < 0 {
		page = 1
	}
	offset := (page - 1) * limit
	return limit, offset
}

func extractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return nil
}
