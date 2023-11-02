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

func testBookGuestData(t *testing.T) {
	t.Parallel()

	query := BookGuestData()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBookGuestDataDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
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

	count, err := BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookGuestDataQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BookGuestData().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookGuestDataSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookGuestDatumSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookGuestDataExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookGuestDatumExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BookGuestDatum exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookGuestDatumExists to return true, but got false.")
	}
}

func testBookGuestDataFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookGuestDatumFound, err := FindBookGuestDatum(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if bookGuestDatumFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBookGuestDataBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BookGuestData().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBookGuestDataOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BookGuestData().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBookGuestDataAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bookGuestDatumOne := &BookGuestDatum{}
	bookGuestDatumTwo := &BookGuestDatum{}
	if err = randomize.Struct(seed, bookGuestDatumOne, bookGuestDatumDBTypes, false, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, bookGuestDatumTwo, bookGuestDatumDBTypes, false, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookGuestDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookGuestDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookGuestData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBookGuestDataCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bookGuestDatumOne := &BookGuestDatum{}
	bookGuestDatumTwo := &BookGuestDatum{}
	if err = randomize.Struct(seed, bookGuestDatumOne, bookGuestDatumDBTypes, false, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, bookGuestDatumTwo, bookGuestDatumDBTypes, false, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookGuestDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookGuestDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookGuestDatumBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
	*o = BookGuestDatum{}
	return nil
}

func bookGuestDatumAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
	*o = BookGuestDatum{}
	return nil
}

func bookGuestDatumAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
	*o = BookGuestDatum{}
	return nil
}

func bookGuestDatumBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
	*o = BookGuestDatum{}
	return nil
}

func bookGuestDatumAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
	*o = BookGuestDatum{}
	return nil
}

func bookGuestDatumBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
	*o = BookGuestDatum{}
	return nil
}

func bookGuestDatumAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
	*o = BookGuestDatum{}
	return nil
}

func bookGuestDatumBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
	*o = BookGuestDatum{}
	return nil
}

func bookGuestDatumAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
	*o = BookGuestDatum{}
	return nil
}

func testBookGuestDataHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BookGuestDatum{}
	o := &BookGuestDatum{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum object: %s", err)
	}

	AddBookGuestDatumHook(boil.BeforeInsertHook, bookGuestDatumBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookGuestDatumBeforeInsertHooks = []BookGuestDatumHook{}

	AddBookGuestDatumHook(boil.AfterInsertHook, bookGuestDatumAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookGuestDatumAfterInsertHooks = []BookGuestDatumHook{}

	AddBookGuestDatumHook(boil.AfterSelectHook, bookGuestDatumAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookGuestDatumAfterSelectHooks = []BookGuestDatumHook{}

	AddBookGuestDatumHook(boil.BeforeUpdateHook, bookGuestDatumBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookGuestDatumBeforeUpdateHooks = []BookGuestDatumHook{}

	AddBookGuestDatumHook(boil.AfterUpdateHook, bookGuestDatumAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookGuestDatumAfterUpdateHooks = []BookGuestDatumHook{}

	AddBookGuestDatumHook(boil.BeforeDeleteHook, bookGuestDatumBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookGuestDatumBeforeDeleteHooks = []BookGuestDatumHook{}

	AddBookGuestDatumHook(boil.AfterDeleteHook, bookGuestDatumAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookGuestDatumAfterDeleteHooks = []BookGuestDatumHook{}

	AddBookGuestDatumHook(boil.BeforeUpsertHook, bookGuestDatumBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookGuestDatumBeforeUpsertHooks = []BookGuestDatumHook{}

	AddBookGuestDatumHook(boil.AfterUpsertHook, bookGuestDatumAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookGuestDatumAfterUpsertHooks = []BookGuestDatumHook{}
}

func testBookGuestDataInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookGuestDataInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookGuestDatumColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookGuestDatumToManyGuestDatumUserBooks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookGuestDatum
	var b, c UserBook

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userBookDBTypes, false, userBookColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userBookDBTypes, false, userBookColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.GuestDataID = a.ID
	c.GuestDataID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.GuestDatumUserBooks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.GuestDataID == b.GuestDataID {
			bFound = true
		}
		if v.GuestDataID == c.GuestDataID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BookGuestDatumSlice{&a}
	if err = a.L.LoadGuestDatumUserBooks(ctx, tx, false, (*[]*BookGuestDatum)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuestDatumUserBooks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GuestDatumUserBooks = nil
	if err = a.L.LoadGuestDatumUserBooks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuestDatumUserBooks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBookGuestDatumToManyAddOpGuestDatumUserBooks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookGuestDatum
	var b, c, d, e UserBook

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookGuestDatumDBTypes, false, strmangle.SetComplement(bookGuestDatumPrimaryKeyColumns, bookGuestDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserBook{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userBookDBTypes, false, strmangle.SetComplement(userBookPrimaryKeyColumns, userBookColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*UserBook{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGuestDatumUserBooks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuestDataID {
			t.Error("foreign key was wrong value", a.ID, first.GuestDataID)
		}
		if a.ID != second.GuestDataID {
			t.Error("foreign key was wrong value", a.ID, second.GuestDataID)
		}

		if first.R.GuestDatum != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.GuestDatum != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GuestDatumUserBooks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GuestDatumUserBooks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GuestDatumUserBooks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBookGuestDataReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
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

func testBookGuestDataReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookGuestDatumSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookGuestDataSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookGuestData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookGuestDatumDBTypes = map[string]string{`ID`: `uuid`, `FirstName`: `character varying`, `LastName`: `character varying`, `FirstNameKana`: `character varying`, `LastNameKana`: `character varying`, `CompanyName`: `character varying`, `BirthDate`: `date`, `ZipCode`: `character varying`, `Prefecture`: `integer`, `City`: `character varying`, `Address`: `text`, `Tel`: `character varying`, `Mail`: `character varying`, `CreateAt`: `timestamp with time zone`, `UpdateAt`: `timestamp with time zone`}
	_                     = bytes.MinRead
)

func testBookGuestDataUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bookGuestDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookGuestDatumAllColumns) == len(bookGuestDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBookGuestDataSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bookGuestDatumAllColumns) == len(bookGuestDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookGuestDatum{}
	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookGuestDatumDBTypes, true, bookGuestDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookGuestDatumAllColumns, bookGuestDatumPrimaryKeyColumns) {
		fields = bookGuestDatumAllColumns
	} else {
		fields = strmangle.SetComplement(
			bookGuestDatumAllColumns,
			bookGuestDatumPrimaryKeyColumns,
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

	slice := BookGuestDatumSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBookGuestDataUpsert(t *testing.T) {
	t.Parallel()

	if len(bookGuestDatumAllColumns) == len(bookGuestDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BookGuestDatum{}
	if err = randomize.Struct(seed, &o, bookGuestDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookGuestDatum: %s", err)
	}

	count, err := BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookGuestDatumDBTypes, false, bookGuestDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookGuestDatum: %s", err)
	}

	count, err = BookGuestData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}