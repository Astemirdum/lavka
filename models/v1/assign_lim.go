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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// AssignLim is an object representing the database table.
type AssignLim struct {
	ID          int         `boil:"id" json:"ID" toml:"ID" yaml:"ID"`
	CourierType CourierType `boil:"courier_type" json:"CourierType" toml:"CourierType" yaml:"CourierType"`
	Lim         types.JSON  `boil:"lim" json:"Lim" toml:"Lim" yaml:"Lim"`

	R *assignLimR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L assignLimL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AssignLimColumns = struct {
	ID          string
	CourierType string
	Lim         string
}{
	ID:          "id",
	CourierType: "courier_type",
	Lim:         "lim",
}

var AssignLimTableColumns = struct {
	ID          string
	CourierType string
	Lim         string
}{
	ID:          "assign_lim.id",
	CourierType: "assign_lim.courier_type",
	Lim:         "assign_lim.lim",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperCourierType struct{ field string }

func (w whereHelperCourierType) EQ(x CourierType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperCourierType) NEQ(x CourierType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperCourierType) LT(x CourierType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperCourierType) LTE(x CourierType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperCourierType) GT(x CourierType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperCourierType) GTE(x CourierType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperCourierType) IN(slice []CourierType) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperCourierType) NIN(slice []CourierType) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_JSON struct{ field string }

func (w whereHelpertypes_JSON) EQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_JSON) NEQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_JSON) LT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSON) LTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSON) GT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSON) GTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AssignLimWhere = struct {
	ID          whereHelperint
	CourierType whereHelperCourierType
	Lim         whereHelpertypes_JSON
}{
	ID:          whereHelperint{field: "\"lk\".\"assign_lim\".\"id\""},
	CourierType: whereHelperCourierType{field: "\"lk\".\"assign_lim\".\"courier_type\""},
	Lim:         whereHelpertypes_JSON{field: "\"lk\".\"assign_lim\".\"lim\""},
}

// AssignLimRels is where relationship names are stored.
var AssignLimRels = struct {
}{}

// assignLimR is where relationships are stored.
type assignLimR struct {
}

// NewStruct creates a new relationship struct
func (*assignLimR) NewStruct() *assignLimR {
	return &assignLimR{}
}

// assignLimL is where Load methods for each relationship are stored.
type assignLimL struct{}

var (
	assignLimAllColumns            = []string{"id", "courier_type", "lim"}
	assignLimColumnsWithoutDefault = []string{"courier_type", "lim"}
	assignLimColumnsWithDefault    = []string{"id"}
	assignLimPrimaryKeyColumns     = []string{"id"}
	assignLimGeneratedColumns      = []string{"id"}
)

type (
	// AssignLimSlice is an alias for a slice of pointers to AssignLim.
	// This should almost always be used instead of []AssignLim.
	AssignLimSlice []*AssignLim

	assignLimQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	assignLimType                 = reflect.TypeOf(&AssignLim{})
	assignLimMapping              = queries.MakeStructMapping(assignLimType)
	assignLimPrimaryKeyMapping, _ = queries.BindMapping(assignLimType, assignLimMapping, assignLimPrimaryKeyColumns)
	assignLimInsertCacheMut       sync.RWMutex
	assignLimInsertCache          = make(map[string]insertCache)
	assignLimUpdateCacheMut       sync.RWMutex
	assignLimUpdateCache          = make(map[string]updateCache)
	assignLimUpsertCacheMut       sync.RWMutex
	assignLimUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single assignLim record from the query.
func (q assignLimQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AssignLim, error) {
	o := &AssignLim{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to execute a one query for assign_lim")
	}

	return o, nil
}

// All returns all AssignLim records from the query.
func (q assignLimQuery) All(ctx context.Context, exec boil.ContextExecutor) (AssignLimSlice, error) {
	var o []*AssignLim

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AssignLim slice")
	}

	return o, nil
}

// Count returns the count of all AssignLim records in the query.
func (q assignLimQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count assign_lim rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q assignLimQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if assign_lim exists")
	}

	return count > 0, nil
}

// AssignLims retrieves all the records using an executor.
func AssignLims(mods ...qm.QueryMod) assignLimQuery {
	mods = append(mods, qm.From("\"lk\".\"assign_lim\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"lk\".\"assign_lim\".*"})
	}

	return assignLimQuery{q}
}

// FindAssignLim retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAssignLim(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*AssignLim, error) {
	assignLimObj := &AssignLim{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"lk\".\"assign_lim\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, assignLimObj)
	if err != nil {
		return nil, errors.Wrap(err, "models: unable to select from assign_lim")
	}

	return assignLimObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AssignLim) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no assign_lim provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(assignLimColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	assignLimInsertCacheMut.RLock()
	cache, cached := assignLimInsertCache[key]
	assignLimInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			assignLimAllColumns,
			assignLimColumnsWithDefault,
			assignLimColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, assignLimGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(assignLimType, assignLimMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(assignLimType, assignLimMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"lk\".\"assign_lim\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"lk\".\"assign_lim\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into assign_lim")
	}

	if !cached {
		assignLimInsertCacheMut.Lock()
		assignLimInsertCache[key] = cache
		assignLimInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the AssignLim.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AssignLim) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	assignLimUpdateCacheMut.RLock()
	cache, cached := assignLimUpdateCache[key]
	assignLimUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			assignLimAllColumns,
			assignLimPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, assignLimGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update assign_lim, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"lk\".\"assign_lim\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, assignLimPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(assignLimType, assignLimMapping, append(wl, assignLimPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update assign_lim row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for assign_lim")
	}

	if !cached {
		assignLimUpdateCacheMut.Lock()
		assignLimUpdateCache[key] = cache
		assignLimUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q assignLimQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for assign_lim")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for assign_lim")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AssignLimSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), assignLimPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"lk\".\"assign_lim\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, assignLimPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in assignLim slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all assignLim")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AssignLim) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no assign_lim provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(assignLimColumnsWithDefault, o)

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

	assignLimUpsertCacheMut.RLock()
	cache, cached := assignLimUpsertCache[key]
	assignLimUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			assignLimAllColumns,
			assignLimColumnsWithDefault,
			assignLimColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			assignLimAllColumns,
			assignLimPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, assignLimGeneratedColumns)
		update = strmangle.SetComplement(update, assignLimGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert assign_lim, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(assignLimPrimaryKeyColumns))
			copy(conflict, assignLimPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"lk\".\"assign_lim\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(assignLimType, assignLimMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(assignLimType, assignLimMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert assign_lim")
	}

	if !cached {
		assignLimUpsertCacheMut.Lock()
		assignLimUpsertCache[key] = cache
		assignLimUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single AssignLim record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AssignLim) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AssignLim provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), assignLimPrimaryKeyMapping)
	sql := "DELETE FROM \"lk\".\"assign_lim\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from assign_lim")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for assign_lim")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q assignLimQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no assignLimQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from assign_lim")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for assign_lim")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AssignLimSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), assignLimPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"lk\".\"assign_lim\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, assignLimPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from assignLim slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for assign_lim")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AssignLim) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAssignLim(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AssignLimSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AssignLimSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), assignLimPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"lk\".\"assign_lim\".* FROM \"lk\".\"assign_lim\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, assignLimPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AssignLimSlice")
	}

	*o = slice

	return nil
}

// AssignLimExists checks if the AssignLim row exists.
func AssignLimExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"lk\".\"assign_lim\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if assign_lim exists")
	}

	return exists, nil
}

// Exists checks if the AssignLim row exists.
func (o *AssignLim) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AssignLimExists(ctx, exec, o.ID)
}
