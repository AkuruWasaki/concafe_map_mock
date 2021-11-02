// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testShopGenres(t *testing.T) {
	t.Parallel()

	query := ShopGenres()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testShopGenresDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testShopGenresQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ShopGenres().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testShopGenresSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ShopGenreSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testShopGenresExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ShopGenreExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ShopGenre exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ShopGenreExists to return true, but got false.")
	}
}

func testShopGenresFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	shopGenreFound, err := FindShopGenre(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if shopGenreFound == nil {
		t.Error("want a record, got nil")
	}
}

func testShopGenresBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ShopGenres().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testShopGenresOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ShopGenres().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testShopGenresAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	shopGenreOne := &ShopGenre{}
	shopGenreTwo := &ShopGenre{}
	if err = randomize.Struct(seed, shopGenreOne, shopGenreDBTypes, false, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}
	if err = randomize.Struct(seed, shopGenreTwo, shopGenreDBTypes, false, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = shopGenreOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = shopGenreTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ShopGenres().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testShopGenresCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	shopGenreOne := &ShopGenre{}
	shopGenreTwo := &ShopGenre{}
	if err = randomize.Struct(seed, shopGenreOne, shopGenreDBTypes, false, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}
	if err = randomize.Struct(seed, shopGenreTwo, shopGenreDBTypes, false, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = shopGenreOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = shopGenreTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func shopGenreBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ShopGenre) error {
	*o = ShopGenre{}
	return nil
}

func shopGenreAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ShopGenre) error {
	*o = ShopGenre{}
	return nil
}

func shopGenreAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ShopGenre) error {
	*o = ShopGenre{}
	return nil
}

func shopGenreBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ShopGenre) error {
	*o = ShopGenre{}
	return nil
}

func shopGenreAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ShopGenre) error {
	*o = ShopGenre{}
	return nil
}

func shopGenreBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ShopGenre) error {
	*o = ShopGenre{}
	return nil
}

func shopGenreAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ShopGenre) error {
	*o = ShopGenre{}
	return nil
}

func shopGenreBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ShopGenre) error {
	*o = ShopGenre{}
	return nil
}

func shopGenreAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ShopGenre) error {
	*o = ShopGenre{}
	return nil
}

func testShopGenresHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ShopGenre{}
	o := &ShopGenre{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, shopGenreDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ShopGenre object: %s", err)
	}

	AddShopGenreHook(boil.BeforeInsertHook, shopGenreBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	shopGenreBeforeInsertHooks = []ShopGenreHook{}

	AddShopGenreHook(boil.AfterInsertHook, shopGenreAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	shopGenreAfterInsertHooks = []ShopGenreHook{}

	AddShopGenreHook(boil.AfterSelectHook, shopGenreAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	shopGenreAfterSelectHooks = []ShopGenreHook{}

	AddShopGenreHook(boil.BeforeUpdateHook, shopGenreBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	shopGenreBeforeUpdateHooks = []ShopGenreHook{}

	AddShopGenreHook(boil.AfterUpdateHook, shopGenreAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	shopGenreAfterUpdateHooks = []ShopGenreHook{}

	AddShopGenreHook(boil.BeforeDeleteHook, shopGenreBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	shopGenreBeforeDeleteHooks = []ShopGenreHook{}

	AddShopGenreHook(boil.AfterDeleteHook, shopGenreAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	shopGenreAfterDeleteHooks = []ShopGenreHook{}

	AddShopGenreHook(boil.BeforeUpsertHook, shopGenreBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	shopGenreBeforeUpsertHooks = []ShopGenreHook{}

	AddShopGenreHook(boil.AfterUpsertHook, shopGenreAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	shopGenreAfterUpsertHooks = []ShopGenreHook{}
}

func testShopGenresInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testShopGenresInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(shopGenreColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testShopGenreToManyShopGenreRelations(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ShopGenre
	var b, c ShopGenreRelation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, shopGenreRelationDBTypes, false, shopGenreRelationColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, shopGenreRelationDBTypes, false, shopGenreRelationColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ShopGenreID = a.ID
	c.ShopGenreID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ShopGenreRelations().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ShopGenreID == b.ShopGenreID {
			bFound = true
		}
		if v.ShopGenreID == c.ShopGenreID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ShopGenreSlice{&a}
	if err = a.L.LoadShopGenreRelations(ctx, tx, false, (*[]*ShopGenre)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ShopGenreRelations); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ShopGenreRelations = nil
	if err = a.L.LoadShopGenreRelations(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ShopGenreRelations); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testShopGenreToManyAddOpShopGenreRelations(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ShopGenre
	var b, c, d, e ShopGenreRelation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, shopGenreDBTypes, false, strmangle.SetComplement(shopGenrePrimaryKeyColumns, shopGenreColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ShopGenreRelation{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, shopGenreRelationDBTypes, false, strmangle.SetComplement(shopGenreRelationPrimaryKeyColumns, shopGenreRelationColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ShopGenreRelation{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddShopGenreRelations(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ShopGenreID {
			t.Error("foreign key was wrong value", a.ID, first.ShopGenreID)
		}
		if a.ID != second.ShopGenreID {
			t.Error("foreign key was wrong value", a.ID, second.ShopGenreID)
		}

		if first.R.ShopGenre != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ShopGenre != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ShopGenreRelations[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ShopGenreRelations[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ShopGenreRelations().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testShopGenresReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testShopGenresReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ShopGenreSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testShopGenresSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ShopGenres().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	shopGenreDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`}
	_                = bytes.MinRead
)

func testShopGenresUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(shopGenrePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(shopGenreAllColumns) == len(shopGenrePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenrePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testShopGenresSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(shopGenreAllColumns) == len(shopGenrePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ShopGenre{}
	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, shopGenreDBTypes, true, shopGenrePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(shopGenreAllColumns, shopGenrePrimaryKeyColumns) {
		fields = shopGenreAllColumns
	} else {
		fields = strmangle.SetComplement(
			shopGenreAllColumns,
			shopGenrePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ShopGenreSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testShopGenresUpsert(t *testing.T) {
	t.Parallel()

	if len(shopGenreAllColumns) == len(shopGenrePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLShopGenreUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ShopGenre{}
	if err = randomize.Struct(seed, &o, shopGenreDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ShopGenre: %s", err)
	}

	count, err := ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, shopGenreDBTypes, false, shopGenrePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ShopGenre struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ShopGenre: %s", err)
	}

	count, err = ShopGenres().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
