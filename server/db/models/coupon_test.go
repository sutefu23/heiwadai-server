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

func testCoupons(t *testing.T) {
	t.Parallel()

	query := Coupons()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCouponsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
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

	count, err := Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCouponsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Coupons().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCouponsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CouponSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCouponsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CouponExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Coupon exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CouponExists to return true, but got false.")
	}
}

func testCouponsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	couponFound, err := FindCoupon(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if couponFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCouponsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Coupons().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCouponsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Coupons().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCouponsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	couponOne := &Coupon{}
	couponTwo := &Coupon{}
	if err = randomize.Struct(seed, couponOne, couponDBTypes, false, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}
	if err = randomize.Struct(seed, couponTwo, couponDBTypes, false, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = couponOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = couponTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Coupons().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCouponsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	couponOne := &Coupon{}
	couponTwo := &Coupon{}
	if err = randomize.Struct(seed, couponOne, couponDBTypes, false, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}
	if err = randomize.Struct(seed, couponTwo, couponDBTypes, false, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = couponOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = couponTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func couponBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
	*o = Coupon{}
	return nil
}

func couponAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
	*o = Coupon{}
	return nil
}

func couponAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
	*o = Coupon{}
	return nil
}

func couponBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
	*o = Coupon{}
	return nil
}

func couponAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
	*o = Coupon{}
	return nil
}

func couponBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
	*o = Coupon{}
	return nil
}

func couponAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
	*o = Coupon{}
	return nil
}

func couponBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
	*o = Coupon{}
	return nil
}

func couponAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Coupon) error {
	*o = Coupon{}
	return nil
}

func testCouponsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Coupon{}
	o := &Coupon{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, couponDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Coupon object: %s", err)
	}

	AddCouponHook(boil.BeforeInsertHook, couponBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	couponBeforeInsertHooks = []CouponHook{}

	AddCouponHook(boil.AfterInsertHook, couponAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	couponAfterInsertHooks = []CouponHook{}

	AddCouponHook(boil.AfterSelectHook, couponAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	couponAfterSelectHooks = []CouponHook{}

	AddCouponHook(boil.BeforeUpdateHook, couponBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	couponBeforeUpdateHooks = []CouponHook{}

	AddCouponHook(boil.AfterUpdateHook, couponAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	couponAfterUpdateHooks = []CouponHook{}

	AddCouponHook(boil.BeforeDeleteHook, couponBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	couponBeforeDeleteHooks = []CouponHook{}

	AddCouponHook(boil.AfterDeleteHook, couponAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	couponAfterDeleteHooks = []CouponHook{}

	AddCouponHook(boil.BeforeUpsertHook, couponBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	couponBeforeUpsertHooks = []CouponHook{}

	AddCouponHook(boil.AfterUpsertHook, couponAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	couponAfterUpsertHooks = []CouponHook{}
}

func testCouponsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCouponsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(couponColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCouponOneToOneCouponNoticeUsingCouponNotice(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign CouponNotice
	var local Coupon

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, couponNoticeDBTypes, true, couponNoticeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponNotice struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreign.CouponID = local.ID
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.CouponNotice().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.CouponID != foreign.CouponID {
		t.Errorf("want: %v, got %v", foreign.CouponID, check.CouponID)
	}

	ranAfterSelectHook := false
	AddCouponNoticeHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *CouponNotice) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := CouponSlice{&local}
	if err = local.L.LoadCouponNotice(ctx, tx, false, (*[]*Coupon)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.CouponNotice == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.CouponNotice = nil
	if err = local.L.LoadCouponNotice(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.CouponNotice == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testCouponOneToOneCouponStoreUsingCouponStore(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign CouponStore
	var local Coupon

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, couponStoreDBTypes, true, couponStoreColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CouponStore struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreign.CouponID = local.ID
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.CouponStore().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.CouponID != foreign.CouponID {
		t.Errorf("want: %v, got %v", foreign.CouponID, check.CouponID)
	}

	ranAfterSelectHook := false
	AddCouponStoreHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *CouponStore) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := CouponSlice{&local}
	if err = local.L.LoadCouponStore(ctx, tx, false, (*[]*Coupon)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.CouponStore == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.CouponStore = nil
	if err = local.L.LoadCouponStore(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.CouponStore == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testCouponOneToOneSetOpCouponNoticeUsingCouponNotice(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Coupon
	var b, c CouponNotice

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, couponDBTypes, false, strmangle.SetComplement(couponPrimaryKeyColumns, couponColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, couponNoticeDBTypes, false, strmangle.SetComplement(couponNoticePrimaryKeyColumns, couponNoticeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, couponNoticeDBTypes, false, strmangle.SetComplement(couponNoticePrimaryKeyColumns, couponNoticeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CouponNotice{&b, &c} {
		err = a.SetCouponNotice(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.CouponNotice != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Coupon != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.ID != x.CouponID {
			t.Error("foreign key was wrong value", a.ID)
		}

		if exists, err := CouponNoticeExists(ctx, tx, x.CouponID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'x' to exist")
		}

		if a.ID != x.CouponID {
			t.Error("foreign key was wrong value", a.ID, x.CouponID)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testCouponOneToOneSetOpCouponStoreUsingCouponStore(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Coupon
	var b, c CouponStore

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, couponDBTypes, false, strmangle.SetComplement(couponPrimaryKeyColumns, couponColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, couponStoreDBTypes, false, strmangle.SetComplement(couponStorePrimaryKeyColumns, couponStoreColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, couponStoreDBTypes, false, strmangle.SetComplement(couponStorePrimaryKeyColumns, couponStoreColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*CouponStore{&b, &c} {
		err = a.SetCouponStore(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.CouponStore != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Coupon != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.ID != x.CouponID {
			t.Error("foreign key was wrong value", a.ID)
		}

		if exists, err := CouponStoreExists(ctx, tx, x.CouponID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'x' to exist")
		}

		if a.ID != x.CouponID {
			t.Error("foreign key was wrong value", a.ID, x.CouponID)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testCouponToManyCouponAttachedUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Coupon
	var b, c CouponAttachedUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, couponAttachedUserDBTypes, false, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, couponAttachedUserDBTypes, false, couponAttachedUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.CouponID = a.ID
	c.CouponID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.CouponAttachedUsers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.CouponID == b.CouponID {
			bFound = true
		}
		if v.CouponID == c.CouponID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CouponSlice{&a}
	if err = a.L.LoadCouponAttachedUsers(ctx, tx, false, (*[]*Coupon)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CouponAttachedUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CouponAttachedUsers = nil
	if err = a.L.LoadCouponAttachedUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CouponAttachedUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCouponToManyAddOpCouponAttachedUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Coupon
	var b, c, d, e CouponAttachedUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, couponDBTypes, false, strmangle.SetComplement(couponPrimaryKeyColumns, couponColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*CouponAttachedUser{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, couponAttachedUserDBTypes, false, strmangle.SetComplement(couponAttachedUserPrimaryKeyColumns, couponAttachedUserColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*CouponAttachedUser{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCouponAttachedUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CouponID {
			t.Error("foreign key was wrong value", a.ID, first.CouponID)
		}
		if a.ID != second.CouponID {
			t.Error("foreign key was wrong value", a.ID, second.CouponID)
		}

		if first.R.Coupon != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Coupon != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CouponAttachedUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CouponAttachedUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CouponAttachedUsers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCouponsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
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

func testCouponsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CouponSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCouponsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Coupons().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	couponDBTypes = map[string]string{`ID`: `uuid`, `Name`: `character varying`, `CouponType`: `integer`, `DiscountAmount`: `integer`, `ExpireAt`: `timestamp with time zone`, `IsCombinationable`: `boolean`, `CouponStatus`: `integer`, `CreateAt`: `timestamp with time zone`, `UpdateAt`: `timestamp with time zone`}
	_             = bytes.MinRead
)

func testCouponsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(couponPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(couponAllColumns) == len(couponPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, couponDBTypes, true, couponPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCouponsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(couponAllColumns) == len(couponPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Coupon{}
	if err = randomize.Struct(seed, o, couponDBTypes, true, couponColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, couponDBTypes, true, couponPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(couponAllColumns, couponPrimaryKeyColumns) {
		fields = couponAllColumns
	} else {
		fields = strmangle.SetComplement(
			couponAllColumns,
			couponPrimaryKeyColumns,
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

	slice := CouponSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCouponsUpsert(t *testing.T) {
	t.Parallel()

	if len(couponAllColumns) == len(couponPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Coupon{}
	if err = randomize.Struct(seed, &o, couponDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Coupon: %s", err)
	}

	count, err := Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, couponDBTypes, false, couponPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Coupon struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Coupon: %s", err)
	}

	count, err = Coupons().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
