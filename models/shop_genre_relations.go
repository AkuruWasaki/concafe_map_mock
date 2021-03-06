// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// ShopGenreRelation is an object representing the database table.
type ShopGenreRelation struct {
	ID          int `boil:"id" json:"id" toml:"id" yaml:"id"`
	ShopGenreID int `boil:"shop_genre_id" json:"shop_genre_id" toml:"shop_genre_id" yaml:"shop_genre_id"`
	ShopID      int `boil:"shop_id" json:"shop_id" toml:"shop_id" yaml:"shop_id"`

	R *shopGenreRelationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L shopGenreRelationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ShopGenreRelationColumns = struct {
	ID          string
	ShopGenreID string
	ShopID      string
}{
	ID:          "id",
	ShopGenreID: "shop_genre_id",
	ShopID:      "shop_id",
}

var ShopGenreRelationTableColumns = struct {
	ID          string
	ShopGenreID string
	ShopID      string
}{
	ID:          "shop_genre_relations.id",
	ShopGenreID: "shop_genre_relations.shop_genre_id",
	ShopID:      "shop_genre_relations.shop_id",
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

var ShopGenreRelationWhere = struct {
	ID          whereHelperint
	ShopGenreID whereHelperint
	ShopID      whereHelperint
}{
	ID:          whereHelperint{field: "`shop_genre_relations`.`id`"},
	ShopGenreID: whereHelperint{field: "`shop_genre_relations`.`shop_genre_id`"},
	ShopID:      whereHelperint{field: "`shop_genre_relations`.`shop_id`"},
}

// ShopGenreRelationRels is where relationship names are stored.
var ShopGenreRelationRels = struct {
	ShopGenre string
	Shop      string
}{
	ShopGenre: "ShopGenre",
	Shop:      "Shop",
}

// shopGenreRelationR is where relationships are stored.
type shopGenreRelationR struct {
	ShopGenre *ShopGenre `boil:"ShopGenre" json:"ShopGenre" toml:"ShopGenre" yaml:"ShopGenre"`
	Shop      *Shop      `boil:"Shop" json:"Shop" toml:"Shop" yaml:"Shop"`
}

// NewStruct creates a new relationship struct
func (*shopGenreRelationR) NewStruct() *shopGenreRelationR {
	return &shopGenreRelationR{}
}

// shopGenreRelationL is where Load methods for each relationship are stored.
type shopGenreRelationL struct{}

var (
	shopGenreRelationAllColumns            = []string{"id", "shop_genre_id", "shop_id"}
	shopGenreRelationColumnsWithoutDefault = []string{"shop_genre_id", "shop_id"}
	shopGenreRelationColumnsWithDefault    = []string{"id"}
	shopGenreRelationPrimaryKeyColumns     = []string{"id"}
)

type (
	// ShopGenreRelationSlice is an alias for a slice of pointers to ShopGenreRelation.
	// This should almost always be used instead of []ShopGenreRelation.
	ShopGenreRelationSlice []*ShopGenreRelation
	// ShopGenreRelationHook is the signature for custom ShopGenreRelation hook methods
	ShopGenreRelationHook func(context.Context, boil.ContextExecutor, *ShopGenreRelation) error

	shopGenreRelationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	shopGenreRelationType                 = reflect.TypeOf(&ShopGenreRelation{})
	shopGenreRelationMapping              = queries.MakeStructMapping(shopGenreRelationType)
	shopGenreRelationPrimaryKeyMapping, _ = queries.BindMapping(shopGenreRelationType, shopGenreRelationMapping, shopGenreRelationPrimaryKeyColumns)
	shopGenreRelationInsertCacheMut       sync.RWMutex
	shopGenreRelationInsertCache          = make(map[string]insertCache)
	shopGenreRelationUpdateCacheMut       sync.RWMutex
	shopGenreRelationUpdateCache          = make(map[string]updateCache)
	shopGenreRelationUpsertCacheMut       sync.RWMutex
	shopGenreRelationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var shopGenreRelationBeforeInsertHooks []ShopGenreRelationHook
var shopGenreRelationBeforeUpdateHooks []ShopGenreRelationHook
var shopGenreRelationBeforeDeleteHooks []ShopGenreRelationHook
var shopGenreRelationBeforeUpsertHooks []ShopGenreRelationHook

var shopGenreRelationAfterInsertHooks []ShopGenreRelationHook
var shopGenreRelationAfterSelectHooks []ShopGenreRelationHook
var shopGenreRelationAfterUpdateHooks []ShopGenreRelationHook
var shopGenreRelationAfterDeleteHooks []ShopGenreRelationHook
var shopGenreRelationAfterUpsertHooks []ShopGenreRelationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ShopGenreRelation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopGenreRelationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ShopGenreRelation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopGenreRelationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ShopGenreRelation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopGenreRelationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ShopGenreRelation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopGenreRelationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ShopGenreRelation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopGenreRelationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ShopGenreRelation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopGenreRelationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ShopGenreRelation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopGenreRelationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ShopGenreRelation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopGenreRelationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ShopGenreRelation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopGenreRelationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddShopGenreRelationHook registers your hook function for all future operations.
func AddShopGenreRelationHook(hookPoint boil.HookPoint, shopGenreRelationHook ShopGenreRelationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		shopGenreRelationBeforeInsertHooks = append(shopGenreRelationBeforeInsertHooks, shopGenreRelationHook)
	case boil.BeforeUpdateHook:
		shopGenreRelationBeforeUpdateHooks = append(shopGenreRelationBeforeUpdateHooks, shopGenreRelationHook)
	case boil.BeforeDeleteHook:
		shopGenreRelationBeforeDeleteHooks = append(shopGenreRelationBeforeDeleteHooks, shopGenreRelationHook)
	case boil.BeforeUpsertHook:
		shopGenreRelationBeforeUpsertHooks = append(shopGenreRelationBeforeUpsertHooks, shopGenreRelationHook)
	case boil.AfterInsertHook:
		shopGenreRelationAfterInsertHooks = append(shopGenreRelationAfterInsertHooks, shopGenreRelationHook)
	case boil.AfterSelectHook:
		shopGenreRelationAfterSelectHooks = append(shopGenreRelationAfterSelectHooks, shopGenreRelationHook)
	case boil.AfterUpdateHook:
		shopGenreRelationAfterUpdateHooks = append(shopGenreRelationAfterUpdateHooks, shopGenreRelationHook)
	case boil.AfterDeleteHook:
		shopGenreRelationAfterDeleteHooks = append(shopGenreRelationAfterDeleteHooks, shopGenreRelationHook)
	case boil.AfterUpsertHook:
		shopGenreRelationAfterUpsertHooks = append(shopGenreRelationAfterUpsertHooks, shopGenreRelationHook)
	}
}

// One returns a single shopGenreRelation record from the query.
func (q shopGenreRelationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ShopGenreRelation, error) {
	o := &ShopGenreRelation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for shop_genre_relations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ShopGenreRelation records from the query.
func (q shopGenreRelationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ShopGenreRelationSlice, error) {
	var o []*ShopGenreRelation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ShopGenreRelation slice")
	}

	if len(shopGenreRelationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ShopGenreRelation records in the query.
func (q shopGenreRelationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count shop_genre_relations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q shopGenreRelationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if shop_genre_relations exists")
	}

	return count > 0, nil
}

// ShopGenre pointed to by the foreign key.
func (o *ShopGenreRelation) ShopGenre(mods ...qm.QueryMod) shopGenreQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ShopGenreID),
	}

	queryMods = append(queryMods, mods...)

	query := ShopGenres(queryMods...)
	queries.SetFrom(query.Query, "`shop_genres`")

	return query
}

// Shop pointed to by the foreign key.
func (o *ShopGenreRelation) Shop(mods ...qm.QueryMod) shopQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ShopID),
	}

	queryMods = append(queryMods, mods...)

	query := Shops(queryMods...)
	queries.SetFrom(query.Query, "`shops`")

	return query
}

// LoadShopGenre allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (shopGenreRelationL) LoadShopGenre(ctx context.Context, e boil.ContextExecutor, singular bool, maybeShopGenreRelation interface{}, mods queries.Applicator) error {
	var slice []*ShopGenreRelation
	var object *ShopGenreRelation

	if singular {
		object = maybeShopGenreRelation.(*ShopGenreRelation)
	} else {
		slice = *maybeShopGenreRelation.(*[]*ShopGenreRelation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &shopGenreRelationR{}
		}
		args = append(args, object.ShopGenreID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &shopGenreRelationR{}
			}

			for _, a := range args {
				if a == obj.ShopGenreID {
					continue Outer
				}
			}

			args = append(args, obj.ShopGenreID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`shop_genres`),
		qm.WhereIn(`shop_genres.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ShopGenre")
	}

	var resultSlice []*ShopGenre
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ShopGenre")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for shop_genres")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for shop_genres")
	}

	if len(shopGenreRelationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ShopGenre = foreign
		if foreign.R == nil {
			foreign.R = &shopGenreR{}
		}
		foreign.R.ShopGenreRelations = append(foreign.R.ShopGenreRelations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ShopGenreID == foreign.ID {
				local.R.ShopGenre = foreign
				if foreign.R == nil {
					foreign.R = &shopGenreR{}
				}
				foreign.R.ShopGenreRelations = append(foreign.R.ShopGenreRelations, local)
				break
			}
		}
	}

	return nil
}

// LoadShop allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (shopGenreRelationL) LoadShop(ctx context.Context, e boil.ContextExecutor, singular bool, maybeShopGenreRelation interface{}, mods queries.Applicator) error {
	var slice []*ShopGenreRelation
	var object *ShopGenreRelation

	if singular {
		object = maybeShopGenreRelation.(*ShopGenreRelation)
	} else {
		slice = *maybeShopGenreRelation.(*[]*ShopGenreRelation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &shopGenreRelationR{}
		}
		args = append(args, object.ShopID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &shopGenreRelationR{}
			}

			for _, a := range args {
				if a == obj.ShopID {
					continue Outer
				}
			}

			args = append(args, obj.ShopID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`shops`),
		qm.WhereIn(`shops.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Shop")
	}

	var resultSlice []*Shop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Shop")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for shops")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for shops")
	}

	if len(shopGenreRelationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Shop = foreign
		if foreign.R == nil {
			foreign.R = &shopR{}
		}
		foreign.R.ShopGenreRelations = append(foreign.R.ShopGenreRelations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ShopID == foreign.ID {
				local.R.Shop = foreign
				if foreign.R == nil {
					foreign.R = &shopR{}
				}
				foreign.R.ShopGenreRelations = append(foreign.R.ShopGenreRelations, local)
				break
			}
		}
	}

	return nil
}

// SetShopGenre of the shopGenreRelation to the related item.
// Sets o.R.ShopGenre to related.
// Adds o to related.R.ShopGenreRelations.
func (o *ShopGenreRelation) SetShopGenre(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ShopGenre) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `shop_genre_relations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"shop_genre_id"}),
		strmangle.WhereClause("`", "`", 0, shopGenreRelationPrimaryKeyColumns),
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

	o.ShopGenreID = related.ID
	if o.R == nil {
		o.R = &shopGenreRelationR{
			ShopGenre: related,
		}
	} else {
		o.R.ShopGenre = related
	}

	if related.R == nil {
		related.R = &shopGenreR{
			ShopGenreRelations: ShopGenreRelationSlice{o},
		}
	} else {
		related.R.ShopGenreRelations = append(related.R.ShopGenreRelations, o)
	}

	return nil
}

// SetShop of the shopGenreRelation to the related item.
// Sets o.R.Shop to related.
// Adds o to related.R.ShopGenreRelations.
func (o *ShopGenreRelation) SetShop(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Shop) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `shop_genre_relations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"shop_id"}),
		strmangle.WhereClause("`", "`", 0, shopGenreRelationPrimaryKeyColumns),
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

	o.ShopID = related.ID
	if o.R == nil {
		o.R = &shopGenreRelationR{
			Shop: related,
		}
	} else {
		o.R.Shop = related
	}

	if related.R == nil {
		related.R = &shopR{
			ShopGenreRelations: ShopGenreRelationSlice{o},
		}
	} else {
		related.R.ShopGenreRelations = append(related.R.ShopGenreRelations, o)
	}

	return nil
}

// ShopGenreRelations retrieves all the records using an executor.
func ShopGenreRelations(mods ...qm.QueryMod) shopGenreRelationQuery {
	mods = append(mods, qm.From("`shop_genre_relations`"))
	return shopGenreRelationQuery{NewQuery(mods...)}
}

// FindShopGenreRelation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindShopGenreRelation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ShopGenreRelation, error) {
	shopGenreRelationObj := &ShopGenreRelation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `shop_genre_relations` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, shopGenreRelationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from shop_genre_relations")
	}

	if err = shopGenreRelationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return shopGenreRelationObj, err
	}

	return shopGenreRelationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ShopGenreRelation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no shop_genre_relations provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shopGenreRelationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	shopGenreRelationInsertCacheMut.RLock()
	cache, cached := shopGenreRelationInsertCache[key]
	shopGenreRelationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			shopGenreRelationAllColumns,
			shopGenreRelationColumnsWithDefault,
			shopGenreRelationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(shopGenreRelationType, shopGenreRelationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(shopGenreRelationType, shopGenreRelationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `shop_genre_relations` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `shop_genre_relations` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `shop_genre_relations` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, shopGenreRelationPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into shop_genre_relations")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == shopGenreRelationMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for shop_genre_relations")
	}

CacheNoHooks:
	if !cached {
		shopGenreRelationInsertCacheMut.Lock()
		shopGenreRelationInsertCache[key] = cache
		shopGenreRelationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ShopGenreRelation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ShopGenreRelation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	shopGenreRelationUpdateCacheMut.RLock()
	cache, cached := shopGenreRelationUpdateCache[key]
	shopGenreRelationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			shopGenreRelationAllColumns,
			shopGenreRelationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update shop_genre_relations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `shop_genre_relations` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, shopGenreRelationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(shopGenreRelationType, shopGenreRelationMapping, append(wl, shopGenreRelationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update shop_genre_relations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for shop_genre_relations")
	}

	if !cached {
		shopGenreRelationUpdateCacheMut.Lock()
		shopGenreRelationUpdateCache[key] = cache
		shopGenreRelationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q shopGenreRelationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for shop_genre_relations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for shop_genre_relations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ShopGenreRelationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shopGenreRelationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `shop_genre_relations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, shopGenreRelationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in shopGenreRelation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all shopGenreRelation")
	}
	return rowsAff, nil
}

var mySQLShopGenreRelationUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ShopGenreRelation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no shop_genre_relations provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shopGenreRelationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLShopGenreRelationUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	shopGenreRelationUpsertCacheMut.RLock()
	cache, cached := shopGenreRelationUpsertCache[key]
	shopGenreRelationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			shopGenreRelationAllColumns,
			shopGenreRelationColumnsWithDefault,
			shopGenreRelationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			shopGenreRelationAllColumns,
			shopGenreRelationPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert shop_genre_relations, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`shop_genre_relations`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `shop_genre_relations` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(shopGenreRelationType, shopGenreRelationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(shopGenreRelationType, shopGenreRelationMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for shop_genre_relations")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == shopGenreRelationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(shopGenreRelationType, shopGenreRelationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for shop_genre_relations")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for shop_genre_relations")
	}

CacheNoHooks:
	if !cached {
		shopGenreRelationUpsertCacheMut.Lock()
		shopGenreRelationUpsertCache[key] = cache
		shopGenreRelationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ShopGenreRelation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ShopGenreRelation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ShopGenreRelation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), shopGenreRelationPrimaryKeyMapping)
	sql := "DELETE FROM `shop_genre_relations` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from shop_genre_relations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for shop_genre_relations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q shopGenreRelationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no shopGenreRelationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shop_genre_relations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shop_genre_relations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ShopGenreRelationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(shopGenreRelationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shopGenreRelationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `shop_genre_relations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, shopGenreRelationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shopGenreRelation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shop_genre_relations")
	}

	if len(shopGenreRelationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ShopGenreRelation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindShopGenreRelation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ShopGenreRelationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ShopGenreRelationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shopGenreRelationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `shop_genre_relations`.* FROM `shop_genre_relations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, shopGenreRelationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ShopGenreRelationSlice")
	}

	*o = slice

	return nil
}

// ShopGenreRelationExists checks if the ShopGenreRelation row exists.
func ShopGenreRelationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `shop_genre_relations` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if shop_genre_relations exists")
	}

	return exists, nil
}
