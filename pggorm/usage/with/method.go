package with

import (
	"gorm.io/gorm"
)

func Apply(tx *gorm.DB, ww ...With) {
	for _, w := range ww {
		w.Apply(tx)
	}
}

func Preload(query string, args ...any) *WPreload {
	return new(WPreload).Set(query, args...)
}

func Order(values ...string) *WOrder {
	return new(WOrder).Set(values...)
}

func Where(query any, args ...any) *WWhere {
	return new(WWhere).Set(query, args...)
}

func Limit(limit int, offset int) *WLimit {
	return new(WLimit).Set(limit, offset)
}

func Joins(query string, args ...any) *WJoins {
	return new(WJoins).Set(query, args...)
}

func GroupBy(query string) *WGroupBy {
	return new(WGroupBy).Set(query)
}

func Having(query string, args ...any) *WHaving {
	return new(WHaving).Set(query, args...)
}

func Distinct(values ...string) *WDistinct {
	return new(WDistinct).Set(values...)
}

func Select(query any, args ...any) *WSelect {
	return new(WSelect).Set(query, args...)
}

func Omit(columns ...string) *WOmit {
	return new(WOmit).Set(columns...)
}
