// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testStayableStoreInfos(t *testing.T) {
	t.Parallel()

	query := StayableStoreInfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStayableStoreInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
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

	count, err := StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStayableStoreInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := StayableStoreInfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStayableStoreInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StayableStoreInfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStayableStoreInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StayableStoreInfoExists(ctx, tx, o.StoreID)
	if err != nil {
		t.Errorf("Unable to check if StayableStoreInfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StayableStoreInfoExists to return true, but got false.")
	}
}

func testStayableStoreInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	stayableStoreInfoFound, err := FindStayableStoreInfo(ctx, tx, o.StoreID)
	if err != nil {
		t.Error(err)
	}

	if stayableStoreInfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStayableStoreInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = StayableStoreInfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStayableStoreInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := StayableStoreInfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStayableStoreInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stayableStoreInfoOne := &StayableStoreInfo{}
	stayableStoreInfoTwo := &StayableStoreInfo{}
	if err = randomize.Struct(seed, stayableStoreInfoOne, stayableStoreInfoDBTypes, false, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, stayableStoreInfoTwo, stayableStoreInfoDBTypes, false, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stayableStoreInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stayableStoreInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StayableStoreInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStayableStoreInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stayableStoreInfoOne := &StayableStoreInfo{}
	stayableStoreInfoTwo := &StayableStoreInfo{}
	if err = randomize.Struct(seed, stayableStoreInfoOne, stayableStoreInfoDBTypes, false, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, stayableStoreInfoTwo, stayableStoreInfoDBTypes, false, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stayableStoreInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stayableStoreInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func stayableStoreInfoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *StayableStoreInfo) error {
	*o = StayableStoreInfo{}
	return nil
}

func stayableStoreInfoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *StayableStoreInfo) error {
	*o = StayableStoreInfo{}
	return nil
}

func stayableStoreInfoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *StayableStoreInfo) error {
	*o = StayableStoreInfo{}
	return nil
}

func stayableStoreInfoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StayableStoreInfo) error {
	*o = StayableStoreInfo{}
	return nil
}

func stayableStoreInfoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StayableStoreInfo) error {
	*o = StayableStoreInfo{}
	return nil
}

func stayableStoreInfoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StayableStoreInfo) error {
	*o = StayableStoreInfo{}
	return nil
}

func stayableStoreInfoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StayableStoreInfo) error {
	*o = StayableStoreInfo{}
	return nil
}

func stayableStoreInfoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StayableStoreInfo) error {
	*o = StayableStoreInfo{}
	return nil
}

func stayableStoreInfoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StayableStoreInfo) error {
	*o = StayableStoreInfo{}
	return nil
}

func testStayableStoreInfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &StayableStoreInfo{}
	o := &StayableStoreInfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo object: %s", err)
	}

	AddStayableStoreInfoHook(boil.BeforeInsertHook, stayableStoreInfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stayableStoreInfoBeforeInsertHooks = []StayableStoreInfoHook{}

	AddStayableStoreInfoHook(boil.AfterInsertHook, stayableStoreInfoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stayableStoreInfoAfterInsertHooks = []StayableStoreInfoHook{}

	AddStayableStoreInfoHook(boil.AfterSelectHook, stayableStoreInfoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stayableStoreInfoAfterSelectHooks = []StayableStoreInfoHook{}

	AddStayableStoreInfoHook(boil.BeforeUpdateHook, stayableStoreInfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stayableStoreInfoBeforeUpdateHooks = []StayableStoreInfoHook{}

	AddStayableStoreInfoHook(boil.AfterUpdateHook, stayableStoreInfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stayableStoreInfoAfterUpdateHooks = []StayableStoreInfoHook{}

	AddStayableStoreInfoHook(boil.BeforeDeleteHook, stayableStoreInfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stayableStoreInfoBeforeDeleteHooks = []StayableStoreInfoHook{}

	AddStayableStoreInfoHook(boil.AfterDeleteHook, stayableStoreInfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stayableStoreInfoAfterDeleteHooks = []StayableStoreInfoHook{}

	AddStayableStoreInfoHook(boil.BeforeUpsertHook, stayableStoreInfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stayableStoreInfoBeforeUpsertHooks = []StayableStoreInfoHook{}

	AddStayableStoreInfoHook(boil.AfterUpsertHook, stayableStoreInfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stayableStoreInfoAfterUpsertHooks = []StayableStoreInfoHook{}
}

func testStayableStoreInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStayableStoreInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(stayableStoreInfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStayableStoreInfoToOneStoreUsingStore(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local StayableStoreInfo
	var foreign Store

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stayableStoreInfoDBTypes, false, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.StoreID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Store().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddStoreHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Store) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := StayableStoreInfoSlice{&local}
	if err = local.L.LoadStore(ctx, tx, false, (*[]*StayableStoreInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Store == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Store = nil
	if err = local.L.LoadStore(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Store == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testStayableStoreInfoToOneSetOpStoreUsingStore(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a StayableStoreInfo
	var b, c Store

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stayableStoreInfoDBTypes, false, strmangle.SetComplement(stayableStoreInfoPrimaryKeyColumns, stayableStoreInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Store{&b, &c} {
		err = a.SetStore(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Store != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StayableStoreInfo != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StoreID != x.ID {
			t.Error("foreign key was wrong value", a.StoreID)
		}

		if exists, err := StayableStoreInfoExists(ctx, tx, a.StoreID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testStayableStoreInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
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

func testStayableStoreInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StayableStoreInfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStayableStoreInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StayableStoreInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stayableStoreInfoDBTypes = map[string]string{`StoreID`: `uuid`, `Parking`: `character varying`, `Latitude`: `double precision`, `Longitude`: `double precision`, `AccessInfo`: `character varying`, `RestAPIURL`: `character varying`, `BookingSystemID`: `character varying`, `CreateAt`: `timestamp with time zone`, `UpdateAt`: `timestamp with time zone`}
	_                        = bytes.MinRead
)

func testStayableStoreInfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(stayableStoreInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(stayableStoreInfoAllColumns) == len(stayableStoreInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStayableStoreInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stayableStoreInfoAllColumns) == len(stayableStoreInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StayableStoreInfo{}
	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stayableStoreInfoDBTypes, true, stayableStoreInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stayableStoreInfoAllColumns, stayableStoreInfoPrimaryKeyColumns) {
		fields = stayableStoreInfoAllColumns
	} else {
		fields = strmangle.SetComplement(
			stayableStoreInfoAllColumns,
			stayableStoreInfoPrimaryKeyColumns,
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

	slice := StayableStoreInfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStayableStoreInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(stayableStoreInfoAllColumns) == len(stayableStoreInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := StayableStoreInfo{}
	if err = randomize.Struct(seed, &o, stayableStoreInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StayableStoreInfo: %s", err)
	}

	count, err := StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, stayableStoreInfoDBTypes, false, stayableStoreInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StayableStoreInfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StayableStoreInfo: %s", err)
	}

	count, err = StayableStoreInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
