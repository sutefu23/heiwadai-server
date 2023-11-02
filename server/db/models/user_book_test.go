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

func testUserBooks(t *testing.T) {
	t.Parallel()

	query := UserBooks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserBooksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
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

	count, err := UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserBooksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserBooks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserBooksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserBookSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserBooksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserBookExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UserBook exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserBookExists to return true, but got false.")
	}
}

func testUserBooksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userBookFound, err := FindUserBook(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userBookFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserBooksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserBooks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserBooksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserBooks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserBooksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userBookOne := &UserBook{}
	userBookTwo := &UserBook{}
	if err = randomize.Struct(seed, userBookOne, userBookDBTypes, false, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}
	if err = randomize.Struct(seed, userBookTwo, userBookDBTypes, false, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userBookOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userBookTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserBooks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserBooksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userBookOne := &UserBook{}
	userBookTwo := &UserBook{}
	if err = randomize.Struct(seed, userBookOne, userBookDBTypes, false, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}
	if err = randomize.Struct(seed, userBookTwo, userBookDBTypes, false, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userBookOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userBookTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userBookBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserBook) error {
	*o = UserBook{}
	return nil
}

func userBookAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserBook) error {
	*o = UserBook{}
	return nil
}

func userBookAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserBook) error {
	*o = UserBook{}
	return nil
}

func userBookBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserBook) error {
	*o = UserBook{}
	return nil
}

func userBookAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserBook) error {
	*o = UserBook{}
	return nil
}

func userBookBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserBook) error {
	*o = UserBook{}
	return nil
}

func userBookAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserBook) error {
	*o = UserBook{}
	return nil
}

func userBookBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserBook) error {
	*o = UserBook{}
	return nil
}

func userBookAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserBook) error {
	*o = UserBook{}
	return nil
}

func testUserBooksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserBook{}
	o := &UserBook{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userBookDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserBook object: %s", err)
	}

	AddUserBookHook(boil.BeforeInsertHook, userBookBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userBookBeforeInsertHooks = []UserBookHook{}

	AddUserBookHook(boil.AfterInsertHook, userBookAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userBookAfterInsertHooks = []UserBookHook{}

	AddUserBookHook(boil.AfterSelectHook, userBookAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userBookAfterSelectHooks = []UserBookHook{}

	AddUserBookHook(boil.BeforeUpdateHook, userBookBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userBookBeforeUpdateHooks = []UserBookHook{}

	AddUserBookHook(boil.AfterUpdateHook, userBookAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userBookAfterUpdateHooks = []UserBookHook{}

	AddUserBookHook(boil.BeforeDeleteHook, userBookBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userBookBeforeDeleteHooks = []UserBookHook{}

	AddUserBookHook(boil.AfterDeleteHook, userBookAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userBookAfterDeleteHooks = []UserBookHook{}

	AddUserBookHook(boil.BeforeUpsertHook, userBookBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userBookBeforeUpsertHooks = []UserBookHook{}

	AddUserBookHook(boil.AfterUpsertHook, userBookAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userBookAfterUpsertHooks = []UserBookHook{}
}

func testUserBooksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserBooksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userBookColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserBookToOneBookPlanUsingBookPlan(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserBook
	var foreign BookPlan

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userBookDBTypes, false, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, bookPlanDBTypes, false, bookPlanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookPlan struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BookPlanID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.BookPlan().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddBookPlanHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *BookPlan) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := UserBookSlice{&local}
	if err = local.L.LoadBookPlan(ctx, tx, false, (*[]*UserBook)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BookPlan == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.BookPlan = nil
	if err = local.L.LoadBookPlan(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BookPlan == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testUserBookToOneUserDatumUsingBookUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserBook
	var foreign UserDatum

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userBookDBTypes, false, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDatumDBTypes, false, userDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserDatum struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BookUserID = foreign.UserID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.BookUser().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UserID != foreign.UserID {
		t.Errorf("want: %v, got %v", foreign.UserID, check.UserID)
	}

	ranAfterSelectHook := false
	AddUserDatumHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *UserDatum) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := UserBookSlice{&local}
	if err = local.L.LoadBookUser(ctx, tx, false, (*[]*UserBook)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BookUser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.BookUser = nil
	if err = local.L.LoadBookUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BookUser == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testUserBookToOneBookGuestDatumUsingGuestDatum(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserBook
	var foreign BookGuestDatum

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userBookDBTypes, false, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, bookGuestDatumDBTypes, false, bookGuestDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookGuestDatum struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.GuestDataID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.GuestDatum().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddBookGuestDatumHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *BookGuestDatum) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := UserBookSlice{&local}
	if err = local.L.LoadGuestDatum(ctx, tx, false, (*[]*UserBook)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.GuestDatum == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.GuestDatum = nil
	if err = local.L.LoadGuestDatum(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.GuestDatum == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testUserBookToOneSetOpBookPlanUsingBookPlan(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserBook
	var b, c BookPlan

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userBookDBTypes, false, strmangle.SetComplement(userBookPrimaryKeyColumns, userBookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookPlanDBTypes, false, strmangle.SetComplement(bookPlanPrimaryKeyColumns, bookPlanColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookPlanDBTypes, false, strmangle.SetComplement(bookPlanPrimaryKeyColumns, bookPlanColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BookPlan{&b, &c} {
		err = a.SetBookPlan(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.BookPlan != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserBooks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BookPlanID != x.ID {
			t.Error("foreign key was wrong value", a.BookPlanID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BookPlanID))
		reflect.Indirect(reflect.ValueOf(&a.BookPlanID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BookPlanID != x.ID {
			t.Error("foreign key was wrong value", a.BookPlanID, x.ID)
		}
	}
}
func testUserBookToOneSetOpUserDatumUsingBookUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserBook
	var b, c UserDatum

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userBookDBTypes, false, strmangle.SetComplement(userBookPrimaryKeyColumns, userBookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDatumDBTypes, false, strmangle.SetComplement(userDatumPrimaryKeyColumns, userDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDatumDBTypes, false, strmangle.SetComplement(userDatumPrimaryKeyColumns, userDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*UserDatum{&b, &c} {
		err = a.SetBookUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.BookUser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BookUserUserBooks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BookUserID != x.UserID {
			t.Error("foreign key was wrong value", a.BookUserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BookUserID))
		reflect.Indirect(reflect.ValueOf(&a.BookUserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BookUserID != x.UserID {
			t.Error("foreign key was wrong value", a.BookUserID, x.UserID)
		}
	}
}
func testUserBookToOneSetOpBookGuestDatumUsingGuestDatum(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserBook
	var b, c BookGuestDatum

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userBookDBTypes, false, strmangle.SetComplement(userBookPrimaryKeyColumns, userBookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookGuestDatumDBTypes, false, strmangle.SetComplement(bookGuestDatumPrimaryKeyColumns, bookGuestDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookGuestDatumDBTypes, false, strmangle.SetComplement(bookGuestDatumPrimaryKeyColumns, bookGuestDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BookGuestDatum{&b, &c} {
		err = a.SetGuestDatum(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.GuestDatum != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.GuestDatumUserBooks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GuestDataID != x.ID {
			t.Error("foreign key was wrong value", a.GuestDataID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GuestDataID))
		reflect.Indirect(reflect.ValueOf(&a.GuestDataID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GuestDataID != x.ID {
			t.Error("foreign key was wrong value", a.GuestDataID, x.ID)
		}
	}
}

func testUserBooksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
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

func testUserBooksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserBookSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserBooksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserBooks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userBookDBTypes = map[string]string{`ID`: `uuid`, `TLBookingNumber`: `character varying`, `StayFrom`: `timestamp without time zone`, `StayTo`: `timestamp without time zone`, `Adult`: `integer`, `Child`: `integer`, `RoomCount`: `integer`, `CheckInTime`: `character varying`, `TotalCost`: `integer`, `GuestDataID`: `uuid`, `BookPlanID`: `uuid`, `BookUserID`: `uuid`, `Note`: `text`, `CreateAt`: `timestamp with time zone`, `UpdateAt`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testUserBooksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userBookPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userBookAllColumns) == len(userBookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserBooksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userBookAllColumns) == len(userBookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserBook{}
	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userBookDBTypes, true, userBookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userBookAllColumns, userBookPrimaryKeyColumns) {
		fields = userBookAllColumns
	} else {
		fields = strmangle.SetComplement(
			userBookAllColumns,
			userBookPrimaryKeyColumns,
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

	slice := UserBookSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserBooksUpsert(t *testing.T) {
	t.Parallel()

	if len(userBookAllColumns) == len(userBookPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserBook{}
	if err = randomize.Struct(seed, &o, userBookDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserBook: %s", err)
	}

	count, err := UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userBookDBTypes, false, userBookPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserBook struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserBook: %s", err)
	}

	count, err = UserBooks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
