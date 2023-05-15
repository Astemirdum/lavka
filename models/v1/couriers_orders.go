// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CouriersOrder is an object representing the database table.
type CouriersOrder struct {
	ID        int64 `boil:"id" json:"ID" toml:"ID" yaml:"ID"`
	OrderID   int64 `boil:"order_id" json:"OrderID" toml:"OrderID" yaml:"OrderID"`
	CourierID int64 `boil:"courier_id" json:"CourierID" toml:"CourierID" yaml:"CourierID"`

	R *couriersOrderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L couriersOrderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CouriersOrderColumns = struct {
	ID        string
	OrderID   string
	CourierID string
}{
	ID:        "id",
	OrderID:   "order_id",
	CourierID: "courier_id",
}

var CouriersOrderTableColumns = struct {
	ID        string
	OrderID   string
	CourierID string
}{
	ID:        "couriers_orders.id",
	OrderID:   "couriers_orders.order_id",
	CourierID: "couriers_orders.courier_id",
}

// Generated where

var CouriersOrderWhere = struct {
	ID        whereHelperint64
	OrderID   whereHelperint64
	CourierID whereHelperint64
}{
	ID:        whereHelperint64{field: "\"lk\".\"couriers_orders\".\"id\""},
	OrderID:   whereHelperint64{field: "\"lk\".\"couriers_orders\".\"order_id\""},
	CourierID: whereHelperint64{field: "\"lk\".\"couriers_orders\".\"courier_id\""},
}

// CouriersOrderRels is where relationship names are stored.
var CouriersOrderRels = struct {
	Courier string
	Order   string
}{
	Courier: "Courier",
	Order:   "Order",
}

// couriersOrderR is where relationships are stored.
type couriersOrderR struct {
	Courier *Courier `boil:"Courier" json:"Courier" toml:"Courier" yaml:"Courier"`
	Order   *Order   `boil:"Order" json:"Order" toml:"Order" yaml:"Order"`
}

// NewStruct creates a new relationship struct
func (*couriersOrderR) NewStruct() *couriersOrderR {
	return &couriersOrderR{}
}

func (r *couriersOrderR) GetCourier() *Courier {
	if r == nil {
		return nil
	}
	return r.Courier
}

func (r *couriersOrderR) GetOrder() *Order {
	if r == nil {
		return nil
	}
	return r.Order
}

// couriersOrderL is where Load methods for each relationship are stored.
type couriersOrderL struct{}

var (
	couriersOrderAllColumns            = []string{"id", "order_id", "courier_id"}
	couriersOrderColumnsWithoutDefault = []string{"order_id", "courier_id"}
	couriersOrderColumnsWithDefault    = []string{"id"}
	couriersOrderPrimaryKeyColumns     = []string{"id"}
	couriersOrderGeneratedColumns      = []string{"id"}
)

type (
	// CouriersOrderSlice is an alias for a slice of pointers to CouriersOrder.
	// This should almost always be used instead of []CouriersOrder.
	CouriersOrderSlice []*CouriersOrder

	couriersOrderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	couriersOrderType                 = reflect.TypeOf(&CouriersOrder{})
	couriersOrderMapping              = queries.MakeStructMapping(couriersOrderType)
	couriersOrderPrimaryKeyMapping, _ = queries.BindMapping(couriersOrderType, couriersOrderMapping, couriersOrderPrimaryKeyColumns)
	couriersOrderInsertCacheMut       sync.RWMutex
	couriersOrderInsertCache          = make(map[string]insertCache)
	couriersOrderUpdateCacheMut       sync.RWMutex
	couriersOrderUpdateCache          = make(map[string]updateCache)
	couriersOrderUpsertCacheMut       sync.RWMutex
	couriersOrderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single couriersOrder record from the query.
func (q couriersOrderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CouriersOrder, error) {
	o := &CouriersOrder{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to execute a one query for couriers_orders")
	}

	return o, nil
}

// All returns all CouriersOrder records from the query.
func (q couriersOrderQuery) All(ctx context.Context, exec boil.ContextExecutor) (CouriersOrderSlice, error) {
	var o []*CouriersOrder

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CouriersOrder slice")
	}

	return o, nil
}

// Count returns the count of all CouriersOrder records in the query.
func (q couriersOrderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count couriers_orders rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q couriersOrderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if couriers_orders exists")
	}

	return count > 0, nil
}

// Courier pointed to by the foreign key.
func (o *CouriersOrder) Courier(mods ...qm.QueryMod) courierQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CourierID),
	}

	queryMods = append(queryMods, mods...)

	return Couriers(queryMods...)
}

// Order pointed to by the foreign key.
func (o *CouriersOrder) Order(mods ...qm.QueryMod) orderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OrderID),
	}

	queryMods = append(queryMods, mods...)

	return Orders(queryMods...)
}

// LoadCourier allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (couriersOrderL) LoadCourier(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCouriersOrder interface{}, mods queries.Applicator) error {
	var slice []*CouriersOrder
	var object *CouriersOrder

	if singular {
		var ok bool
		object, ok = maybeCouriersOrder.(*CouriersOrder)
		if !ok {
			object = new(CouriersOrder)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCouriersOrder)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCouriersOrder))
			}
		}
	} else {
		s, ok := maybeCouriersOrder.(*[]*CouriersOrder)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCouriersOrder)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCouriersOrder))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &couriersOrderR{}
		}
		args = append(args, object.CourierID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &couriersOrderR{}
			}

			for _, a := range args {
				if a == obj.CourierID {
					continue Outer
				}
			}

			args = append(args, obj.CourierID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`lk.couriers`),
		qm.WhereIn(`lk.couriers.id in ?`, args...),
		qmhelper.WhereIsNull(`lk.couriers.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Courier")
	}

	var resultSlice []*Courier
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Courier")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for couriers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for couriers")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Courier = foreign
		if foreign.R == nil {
			foreign.R = &courierR{}
		}
		foreign.R.CouriersOrders = append(foreign.R.CouriersOrders, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CourierID == foreign.ID {
				local.R.Courier = foreign
				if foreign.R == nil {
					foreign.R = &courierR{}
				}
				foreign.R.CouriersOrders = append(foreign.R.CouriersOrders, local)
				break
			}
		}
	}

	return nil
}

// LoadOrder allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (couriersOrderL) LoadOrder(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCouriersOrder interface{}, mods queries.Applicator) error {
	var slice []*CouriersOrder
	var object *CouriersOrder

	if singular {
		var ok bool
		object, ok = maybeCouriersOrder.(*CouriersOrder)
		if !ok {
			object = new(CouriersOrder)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCouriersOrder)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCouriersOrder))
			}
		}
	} else {
		s, ok := maybeCouriersOrder.(*[]*CouriersOrder)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCouriersOrder)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCouriersOrder))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &couriersOrderR{}
		}
		args = append(args, object.OrderID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &couriersOrderR{}
			}

			for _, a := range args {
				if a == obj.OrderID {
					continue Outer
				}
			}

			args = append(args, obj.OrderID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`lk.orders`),
		qm.WhereIn(`lk.orders.id in ?`, args...),
		qmhelper.WhereIsNull(`lk.orders.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Order")
	}

	var resultSlice []*Order
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Order")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for orders")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for orders")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Order = foreign
		if foreign.R == nil {
			foreign.R = &orderR{}
		}
		foreign.R.CouriersOrders = append(foreign.R.CouriersOrders, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrderID == foreign.ID {
				local.R.Order = foreign
				if foreign.R == nil {
					foreign.R = &orderR{}
				}
				foreign.R.CouriersOrders = append(foreign.R.CouriersOrders, local)
				break
			}
		}
	}

	return nil
}

// SetCourier of the couriersOrder to the related item.
// Sets o.R.Courier to related.
// Adds o to related.R.CouriersOrders.
func (o *CouriersOrder) SetCourier(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Courier) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"lk\".\"couriers_orders\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"courier_id"}),
		strmangle.WhereClause("\"", "\"", 2, couriersOrderPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CourierID = related.ID
	if o.R == nil {
		o.R = &couriersOrderR{
			Courier: related,
		}
	} else {
		o.R.Courier = related
	}

	if related.R == nil {
		related.R = &courierR{
			CouriersOrders: CouriersOrderSlice{o},
		}
	} else {
		related.R.CouriersOrders = append(related.R.CouriersOrders, o)
	}

	return nil
}

// SetOrder of the couriersOrder to the related item.
// Sets o.R.Order to related.
// Adds o to related.R.CouriersOrders.
func (o *CouriersOrder) SetOrder(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Order) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"lk\".\"couriers_orders\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"order_id"}),
		strmangle.WhereClause("\"", "\"", 2, couriersOrderPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrderID = related.ID
	if o.R == nil {
		o.R = &couriersOrderR{
			Order: related,
		}
	} else {
		o.R.Order = related
	}

	if related.R == nil {
		related.R = &orderR{
			CouriersOrders: CouriersOrderSlice{o},
		}
	} else {
		related.R.CouriersOrders = append(related.R.CouriersOrders, o)
	}

	return nil
}

// CouriersOrders retrieves all the records using an executor.
func CouriersOrders(mods ...qm.QueryMod) couriersOrderQuery {
	mods = append(mods, qm.From("\"lk\".\"couriers_orders\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"lk\".\"couriers_orders\".*"})
	}

	return couriersOrderQuery{q}
}

// FindCouriersOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCouriersOrder(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CouriersOrder, error) {
	couriersOrderObj := &CouriersOrder{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"lk\".\"couriers_orders\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, couriersOrderObj)
	if err != nil {
		return nil, errors.Wrap(err, "models: unable to select from couriers_orders")
	}

	return couriersOrderObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CouriersOrder) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no couriers_orders provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(couriersOrderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	couriersOrderInsertCacheMut.RLock()
	cache, cached := couriersOrderInsertCache[key]
	couriersOrderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			couriersOrderAllColumns,
			couriersOrderColumnsWithDefault,
			couriersOrderColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, couriersOrderGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(couriersOrderType, couriersOrderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(couriersOrderType, couriersOrderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"lk\".\"couriers_orders\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"lk\".\"couriers_orders\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into couriers_orders")
	}

	if !cached {
		couriersOrderInsertCacheMut.Lock()
		couriersOrderInsertCache[key] = cache
		couriersOrderInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CouriersOrder.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CouriersOrder) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	couriersOrderUpdateCacheMut.RLock()
	cache, cached := couriersOrderUpdateCache[key]
	couriersOrderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			couriersOrderAllColumns,
			couriersOrderPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, couriersOrderGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update couriers_orders, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"lk\".\"couriers_orders\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, couriersOrderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(couriersOrderType, couriersOrderMapping, append(wl, couriersOrderPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update couriers_orders row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for couriers_orders")
	}

	if !cached {
		couriersOrderUpdateCacheMut.Lock()
		couriersOrderUpdateCache[key] = cache
		couriersOrderUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q couriersOrderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for couriers_orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for couriers_orders")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CouriersOrderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), couriersOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"lk\".\"couriers_orders\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, couriersOrderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in couriersOrder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all couriersOrder")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CouriersOrder) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no couriers_orders provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(couriersOrderColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	couriersOrderUpsertCacheMut.RLock()
	cache, cached := couriersOrderUpsertCache[key]
	couriersOrderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			couriersOrderAllColumns,
			couriersOrderColumnsWithDefault,
			couriersOrderColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			couriersOrderAllColumns,
			couriersOrderPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, couriersOrderGeneratedColumns)
		update = strmangle.SetComplement(update, couriersOrderGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert couriers_orders, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(couriersOrderPrimaryKeyColumns))
			copy(conflict, couriersOrderPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"lk\".\"couriers_orders\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(couriersOrderType, couriersOrderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(couriersOrderType, couriersOrderMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert couriers_orders")
	}

	if !cached {
		couriersOrderUpsertCacheMut.Lock()
		couriersOrderUpsertCache[key] = cache
		couriersOrderUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CouriersOrder record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CouriersOrder) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CouriersOrder provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), couriersOrderPrimaryKeyMapping)
	sql := "DELETE FROM \"lk\".\"couriers_orders\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from couriers_orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for couriers_orders")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q couriersOrderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no couriersOrderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from couriers_orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for couriers_orders")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CouriersOrderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), couriersOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"lk\".\"couriers_orders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, couriersOrderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from couriersOrder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for couriers_orders")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CouriersOrder) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCouriersOrder(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CouriersOrderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CouriersOrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), couriersOrderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"lk\".\"couriers_orders\".* FROM \"lk\".\"couriers_orders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, couriersOrderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CouriersOrderSlice")
	}

	*o = slice

	return nil
}

// CouriersOrderExists checks if the CouriersOrder row exists.
func CouriersOrderExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"lk\".\"couriers_orders\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if couriers_orders exists")
	}

	return exists, nil
}

// Exists checks if the CouriersOrder row exists.
func (o *CouriersOrder) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CouriersOrderExists(ctx, exec, o.ID)
}