package with

import (
	"gorm.io/gorm"
)

type With interface {
	Apply(tx *gorm.DB)
}

type condQueryArgs struct {
	query any
	args  []any
}

// WPreload - preload associations with given conditions
// [docs]: https://gorm.io/docs/preload.html#Preload
type WPreload struct {
	cond []condQueryArgs
}

func (w *WPreload) Set(query string, args ...any) *WPreload {
	w.cond = append(w.cond, condQueryArgs{query: query, args: args})
	return w
}

func (w *WPreload) Apply(tx *gorm.DB) {
	for _, cond := range w.cond {
		tx.Preload(cond.query.(string), cond.args)
	}
}

// WOrder - specify order when retrieving records from database
// [docs]: https://gorm.io/docs/query.html#Order
type WOrder struct {
	values []string
}

func (w *WOrder) Set(values ...string) *WOrder {
	w.values = append(w.values, values...)
	return w
}

func (w *WOrder) Apply(tx *gorm.DB) {
	for _, value := range w.values {
		tx.Order(value)
	}
}

// WWhere - add where conditions
// [docs]: https://gorm.io/docs/query.html#Conditions
type WWhere struct {
	cond []condQueryArgs
}

func (w *WWhere) Set(query any, args ...any) *WWhere {
	w.cond = append(w.cond, condQueryArgs{query: query, args: args})
	return w
}

func (w *WWhere) Apply(tx *gorm.DB) {
	for _, c := range w.cond {
		tx.Where(c.query, c.args...)
	}
}

// WLimit - specify the [limit] number of records to be retrieved and skip [offset] before starting
// [docs]: https://gorm.io/docs/query.html#Limit-amp-Offset
type WLimit struct {
	limit  int
	offset int
}

func (w *WLimit) Set(limit int, offset int) *WLimit {
	w.limit = limit
	w.offset = offset
	return w
}

func (w *WLimit) Apply(tx *gorm.DB) {
	tx.Limit(w.limit).Offset(w.offset)
}

// WJoins - specify Joins conditions
// [docs]: https://gorm.io/docs/query.html#Joins
// [docs]: https://gorm.io/docs/query.html#Joins-Preloading
type WJoins struct {
	cond []condQueryArgs
}

func (w *WJoins) Set(query string, args ...any) *WJoins {
	w.cond = append(w.cond, condQueryArgs{query: query, args: args})
	return w
}

func (w *WJoins) Apply(tx *gorm.DB) {
	for _, cond := range w.cond {
		tx.Joins(cond.query.(string), cond.args...)
	}
}

// WGroupBy - specify the group method on the find
// [docs]: https://gorm.io/docs/query.html#Group-By-amp-Having
type WGroupBy struct {
	cond []string
}

func (w *WGroupBy) Set(query string) *WGroupBy {
	w.cond = append(w.cond, query)
	return w
}

func (w *WGroupBy) Apply(tx *gorm.DB) {
	for _, query := range w.cond {
		tx.Group(query)
	}
}

// WHaving - specify HAVING conditions for GROUP BY
// [docs]: https://gorm.io/docs/query.html#Group-By-amp-Having
type WHaving struct {
	cond map[string][]any
}

func (w *WHaving) Set(query string, args ...any) *WHaving {
	if w.cond == nil {
		w.cond = make(map[string][]any)
	}
	w.cond[query] = append(w.cond[query], args...)
	return w
}

func (w *WHaving) Apply(tx *gorm.DB) {
	if w.cond == nil {
		return
	}
	for query, args := range w.cond {
		tx.Having(query, args)
	}
}

// WDistinct - specify distinct fields that you want querying
// [docs]: https://gorm.io/docs/query.html#Distinct
type WDistinct struct {
	values []string
}

func (w *WDistinct) Set(values ...string) *WDistinct {
	w.values = append(w.values, values...)
	return w
}

func (w *WDistinct) Apply(tx *gorm.DB) {
	tx.Distinct(w.values)
}

// WSelect - specify fields that you want when querying, creating, updating
// [docs]: https://gorm.io/docs/query.html#Selecting-Specific-Fields
type WSelect struct {
	condition []condQueryArgs
}

func (w *WSelect) Set(query any, args ...any) *WSelect {
	w.condition = append(w.condition, condQueryArgs{query: query, args: args})
	return w
}

func (w *WSelect) Apply(tx *gorm.DB) {
	for _, c := range w.condition {
		tx.Select(c.query, c.args...)
	}
}

// WOmit - specify fields that you want to ignore when creating, updating and querying
// [docs]: https://gorm.io/docs/associations.html#Select-x2F-Omit-Association-fields
type WOmit struct {
	columns []string
}

func (w *WOmit) Set(columns ...string) *WOmit {
	w.columns = append(w.columns, columns...)
	return w
}

func (w *WOmit) Apply(tx *gorm.DB) {
	tx.Omit(w.columns...)
}
