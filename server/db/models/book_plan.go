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

// BookPlan is an object representing the database table.
type BookPlan struct {
	ID              string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	PlanID          string    `boil:"plan_id" json:"plan_id" toml:"plan_id" yaml:"plan_id"`
	Title           string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Price           int       `boil:"price" json:"price" toml:"price" yaml:"price"`
	ImageURL        string    `boil:"image_url" json:"image_url" toml:"image_url" yaml:"image_url"`
	RoomType        int       `boil:"room_type" json:"room_type" toml:"room_type" yaml:"room_type"`
	MealTypeMorning bool      `boil:"meal_type_morning" json:"meal_type_morning" toml:"meal_type_morning" yaml:"meal_type_morning"`
	MealTypeDinner  bool      `boil:"meal_type_dinner" json:"meal_type_dinner" toml:"meal_type_dinner" yaml:"meal_type_dinner"`
	SmokeType       int       `boil:"smoke_type" json:"smoke_type" toml:"smoke_type" yaml:"smoke_type"`
	Overview        string    `boil:"overview" json:"overview" toml:"overview" yaml:"overview"`
	StoreID         string    `boil:"store_id" json:"store_id" toml:"store_id" yaml:"store_id"`
	CreateAt        time.Time `boil:"create_at" json:"create_at" toml:"create_at" yaml:"create_at"`
	UpdateAt        time.Time `boil:"update_at" json:"update_at" toml:"update_at" yaml:"update_at"`

	R *bookPlanR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookPlanL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookPlanColumns = struct {
	ID              string
	PlanID          string
	Title           string
	Price           string
	ImageURL        string
	RoomType        string
	MealTypeMorning string
	MealTypeDinner  string
	SmokeType       string
	Overview        string
	StoreID         string
	CreateAt        string
	UpdateAt        string
}{
	ID:              "id",
	PlanID:          "plan_id",
	Title:           "title",
	Price:           "price",
	ImageURL:        "image_url",
	RoomType:        "room_type",
	MealTypeMorning: "meal_type_morning",
	MealTypeDinner:  "meal_type_dinner",
	SmokeType:       "smoke_type",
	Overview:        "overview",
	StoreID:         "store_id",
	CreateAt:        "create_at",
	UpdateAt:        "update_at",
}

var BookPlanTableColumns = struct {
	ID              string
	PlanID          string
	Title           string
	Price           string
	ImageURL        string
	RoomType        string
	MealTypeMorning string
	MealTypeDinner  string
	SmokeType       string
	Overview        string
	StoreID         string
	CreateAt        string
	UpdateAt        string
}{
	ID:              "book_plan.id",
	PlanID:          "book_plan.plan_id",
	Title:           "book_plan.title",
	Price:           "book_plan.price",
	ImageURL:        "book_plan.image_url",
	RoomType:        "book_plan.room_type",
	MealTypeMorning: "book_plan.meal_type_morning",
	MealTypeDinner:  "book_plan.meal_type_dinner",
	SmokeType:       "book_plan.smoke_type",
	Overview:        "book_plan.overview",
	StoreID:         "book_plan.store_id",
	CreateAt:        "book_plan.create_at",
	UpdateAt:        "book_plan.update_at",
}

// Generated where

var BookPlanWhere = struct {
	ID              whereHelperstring
	PlanID          whereHelperstring
	Title           whereHelperstring
	Price           whereHelperint
	ImageURL        whereHelperstring
	RoomType        whereHelperint
	MealTypeMorning whereHelperbool
	MealTypeDinner  whereHelperbool
	SmokeType       whereHelperint
	Overview        whereHelperstring
	StoreID         whereHelperstring
	CreateAt        whereHelpertime_Time
	UpdateAt        whereHelpertime_Time
}{
	ID:              whereHelperstring{field: "\"book_plan\".\"id\""},
	PlanID:          whereHelperstring{field: "\"book_plan\".\"plan_id\""},
	Title:           whereHelperstring{field: "\"book_plan\".\"title\""},
	Price:           whereHelperint{field: "\"book_plan\".\"price\""},
	ImageURL:        whereHelperstring{field: "\"book_plan\".\"image_url\""},
	RoomType:        whereHelperint{field: "\"book_plan\".\"room_type\""},
	MealTypeMorning: whereHelperbool{field: "\"book_plan\".\"meal_type_morning\""},
	MealTypeDinner:  whereHelperbool{field: "\"book_plan\".\"meal_type_dinner\""},
	SmokeType:       whereHelperint{field: "\"book_plan\".\"smoke_type\""},
	Overview:        whereHelperstring{field: "\"book_plan\".\"overview\""},
	StoreID:         whereHelperstring{field: "\"book_plan\".\"store_id\""},
	CreateAt:        whereHelpertime_Time{field: "\"book_plan\".\"create_at\""},
	UpdateAt:        whereHelpertime_Time{field: "\"book_plan\".\"update_at\""},
}

// BookPlanRels is where relationship names are stored.
var BookPlanRels = struct {
	Store     string
	UserBooks string
}{
	Store:     "Store",
	UserBooks: "UserBooks",
}

// bookPlanR is where relationships are stored.
type bookPlanR struct {
	Store     *Store        `boil:"Store" json:"Store" toml:"Store" yaml:"Store"`
	UserBooks UserBookSlice `boil:"UserBooks" json:"UserBooks" toml:"UserBooks" yaml:"UserBooks"`
}

// NewStruct creates a new relationship struct
func (*bookPlanR) NewStruct() *bookPlanR {
	return &bookPlanR{}
}

func (r *bookPlanR) GetStore() *Store {
	if r == nil {
		return nil
	}
	return r.Store
}

func (r *bookPlanR) GetUserBooks() UserBookSlice {
	if r == nil {
		return nil
	}
	return r.UserBooks
}

// bookPlanL is where Load methods for each relationship are stored.
type bookPlanL struct{}

var (
	bookPlanAllColumns            = []string{"id", "plan_id", "title", "price", "image_url", "room_type", "meal_type_morning", "meal_type_dinner", "smoke_type", "overview", "store_id", "create_at", "update_at"}
	bookPlanColumnsWithoutDefault = []string{"id", "plan_id", "title", "price", "image_url", "room_type", "meal_type_morning", "meal_type_dinner", "smoke_type", "overview", "store_id"}
	bookPlanColumnsWithDefault    = []string{"create_at", "update_at"}
	bookPlanPrimaryKeyColumns     = []string{"id"}
	bookPlanGeneratedColumns      = []string{}
)

type (
	// BookPlanSlice is an alias for a slice of pointers to BookPlan.
	// This should almost always be used instead of []BookPlan.
	BookPlanSlice []*BookPlan
	// BookPlanHook is the signature for custom BookPlan hook methods
	BookPlanHook func(context.Context, boil.ContextExecutor, *BookPlan) error

	bookPlanQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookPlanType                 = reflect.TypeOf(&BookPlan{})
	bookPlanMapping              = queries.MakeStructMapping(bookPlanType)
	bookPlanPrimaryKeyMapping, _ = queries.BindMapping(bookPlanType, bookPlanMapping, bookPlanPrimaryKeyColumns)
	bookPlanInsertCacheMut       sync.RWMutex
	bookPlanInsertCache          = make(map[string]insertCache)
	bookPlanUpdateCacheMut       sync.RWMutex
	bookPlanUpdateCache          = make(map[string]updateCache)
	bookPlanUpsertCacheMut       sync.RWMutex
	bookPlanUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookPlanAfterSelectHooks []BookPlanHook

var bookPlanBeforeInsertHooks []BookPlanHook
var bookPlanAfterInsertHooks []BookPlanHook

var bookPlanBeforeUpdateHooks []BookPlanHook
var bookPlanAfterUpdateHooks []BookPlanHook

var bookPlanBeforeDeleteHooks []BookPlanHook
var bookPlanAfterDeleteHooks []BookPlanHook

var bookPlanBeforeUpsertHooks []BookPlanHook
var bookPlanAfterUpsertHooks []BookPlanHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BookPlan) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookPlanAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BookPlan) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookPlanBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BookPlan) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookPlanAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BookPlan) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookPlanBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BookPlan) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookPlanAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BookPlan) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookPlanBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BookPlan) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookPlanAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BookPlan) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookPlanBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BookPlan) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookPlanAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookPlanHook registers your hook function for all future operations.
func AddBookPlanHook(hookPoint boil.HookPoint, bookPlanHook BookPlanHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bookPlanAfterSelectHooks = append(bookPlanAfterSelectHooks, bookPlanHook)
	case boil.BeforeInsertHook:
		bookPlanBeforeInsertHooks = append(bookPlanBeforeInsertHooks, bookPlanHook)
	case boil.AfterInsertHook:
		bookPlanAfterInsertHooks = append(bookPlanAfterInsertHooks, bookPlanHook)
	case boil.BeforeUpdateHook:
		bookPlanBeforeUpdateHooks = append(bookPlanBeforeUpdateHooks, bookPlanHook)
	case boil.AfterUpdateHook:
		bookPlanAfterUpdateHooks = append(bookPlanAfterUpdateHooks, bookPlanHook)
	case boil.BeforeDeleteHook:
		bookPlanBeforeDeleteHooks = append(bookPlanBeforeDeleteHooks, bookPlanHook)
	case boil.AfterDeleteHook:
		bookPlanAfterDeleteHooks = append(bookPlanAfterDeleteHooks, bookPlanHook)
	case boil.BeforeUpsertHook:
		bookPlanBeforeUpsertHooks = append(bookPlanBeforeUpsertHooks, bookPlanHook)
	case boil.AfterUpsertHook:
		bookPlanAfterUpsertHooks = append(bookPlanAfterUpsertHooks, bookPlanHook)
	}
}

// One returns a single bookPlan record from the query.
func (q bookPlanQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BookPlan, error) {
	o := &BookPlan{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for book_plan")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BookPlan records from the query.
func (q bookPlanQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookPlanSlice, error) {
	var o []*BookPlan

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BookPlan slice")
	}

	if len(bookPlanAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BookPlan records in the query.
func (q bookPlanQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count book_plan rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bookPlanQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if book_plan exists")
	}

	return count > 0, nil
}

// Store pointed to by the foreign key.
func (o *BookPlan) Store(mods ...qm.QueryMod) storeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.StoreID),
	}

	queryMods = append(queryMods, mods...)

	return Stores(queryMods...)
}

// UserBooks retrieves all the user_book's UserBooks with an executor.
func (o *BookPlan) UserBooks(mods ...qm.QueryMod) userBookQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_book\".\"book_plan_id\"=?", o.ID),
	)

	return UserBooks(queryMods...)
}

// LoadStore allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (bookPlanL) LoadStore(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBookPlan interface{}, mods queries.Applicator) error {
	var slice []*BookPlan
	var object *BookPlan

	if singular {
		var ok bool
		object, ok = maybeBookPlan.(*BookPlan)
		if !ok {
			object = new(BookPlan)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBookPlan)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBookPlan))
			}
		}
	} else {
		s, ok := maybeBookPlan.(*[]*BookPlan)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBookPlan)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBookPlan))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookPlanR{}
		}
		args = append(args, object.StoreID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookPlanR{}
			}

			for _, a := range args {
				if a == obj.StoreID {
					continue Outer
				}
			}

			args = append(args, obj.StoreID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`store`),
		qm.WhereIn(`store.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Store")
	}

	var resultSlice []*Store
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Store")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for store")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for store")
	}

	if len(storeAfterSelectHooks) != 0 {
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
		object.R.Store = foreign
		if foreign.R == nil {
			foreign.R = &storeR{}
		}
		foreign.R.BookPlans = append(foreign.R.BookPlans, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StoreID == foreign.ID {
				local.R.Store = foreign
				if foreign.R == nil {
					foreign.R = &storeR{}
				}
				foreign.R.BookPlans = append(foreign.R.BookPlans, local)
				break
			}
		}
	}

	return nil
}

// LoadUserBooks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bookPlanL) LoadUserBooks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBookPlan interface{}, mods queries.Applicator) error {
	var slice []*BookPlan
	var object *BookPlan

	if singular {
		var ok bool
		object, ok = maybeBookPlan.(*BookPlan)
		if !ok {
			object = new(BookPlan)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBookPlan)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBookPlan))
			}
		}
	} else {
		s, ok := maybeBookPlan.(*[]*BookPlan)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBookPlan)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBookPlan))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookPlanR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookPlanR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_book`),
		qm.WhereIn(`user_book.book_plan_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_book")
	}

	var resultSlice []*UserBook
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_book")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_book")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_book")
	}

	if len(userBookAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserBooks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userBookR{}
			}
			foreign.R.BookPlan = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.BookPlanID {
				local.R.UserBooks = append(local.R.UserBooks, foreign)
				if foreign.R == nil {
					foreign.R = &userBookR{}
				}
				foreign.R.BookPlan = local
				break
			}
		}
	}

	return nil
}

// SetStore of the bookPlan to the related item.
// Sets o.R.Store to related.
// Adds o to related.R.BookPlans.
func (o *BookPlan) SetStore(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Store) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"book_plan\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"store_id"}),
		strmangle.WhereClause("\"", "\"", 2, bookPlanPrimaryKeyColumns),
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

	o.StoreID = related.ID
	if o.R == nil {
		o.R = &bookPlanR{
			Store: related,
		}
	} else {
		o.R.Store = related
	}

	if related.R == nil {
		related.R = &storeR{
			BookPlans: BookPlanSlice{o},
		}
	} else {
		related.R.BookPlans = append(related.R.BookPlans, o)
	}

	return nil
}

// AddUserBooks adds the given related objects to the existing relationships
// of the book_plan, optionally inserting them as new records.
// Appends related to o.R.UserBooks.
// Sets related.R.BookPlan appropriately.
func (o *BookPlan) AddUserBooks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserBook) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.BookPlanID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_book\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"book_plan_id"}),
				strmangle.WhereClause("\"", "\"", 2, userBookPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.BookPlanID = o.ID
		}
	}

	if o.R == nil {
		o.R = &bookPlanR{
			UserBooks: related,
		}
	} else {
		o.R.UserBooks = append(o.R.UserBooks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userBookR{
				BookPlan: o,
			}
		} else {
			rel.R.BookPlan = o
		}
	}
	return nil
}

// BookPlans retrieves all the records using an executor.
func BookPlans(mods ...qm.QueryMod) bookPlanQuery {
	mods = append(mods, qm.From("\"book_plan\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"book_plan\".*"})
	}

	return bookPlanQuery{q}
}

// FindBookPlan retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookPlan(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*BookPlan, error) {
	bookPlanObj := &BookPlan{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"book_plan\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bookPlanObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from book_plan")
	}

	if err = bookPlanObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bookPlanObj, err
	}

	return bookPlanObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BookPlan) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book_plan provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreateAt.IsZero() {
			o.CreateAt = currTime
		}
		if o.UpdateAt.IsZero() {
			o.UpdateAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookPlanColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookPlanInsertCacheMut.RLock()
	cache, cached := bookPlanInsertCache[key]
	bookPlanInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookPlanAllColumns,
			bookPlanColumnsWithDefault,
			bookPlanColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookPlanType, bookPlanMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookPlanType, bookPlanMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"book_plan\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"book_plan\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into book_plan")
	}

	if !cached {
		bookPlanInsertCacheMut.Lock()
		bookPlanInsertCache[key] = cache
		bookPlanInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BookPlan.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BookPlan) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdateAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookPlanUpdateCacheMut.RLock()
	cache, cached := bookPlanUpdateCache[key]
	bookPlanUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookPlanAllColumns,
			bookPlanPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update book_plan, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"book_plan\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bookPlanPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookPlanType, bookPlanMapping, append(wl, bookPlanPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update book_plan row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for book_plan")
	}

	if !cached {
		bookPlanUpdateCacheMut.Lock()
		bookPlanUpdateCache[key] = cache
		bookPlanUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bookPlanQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for book_plan")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for book_plan")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookPlanSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookPlanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"book_plan\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bookPlanPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bookPlan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bookPlan")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BookPlan) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book_plan provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreateAt.IsZero() {
			o.CreateAt = currTime
		}
		o.UpdateAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookPlanColumnsWithDefault, o)

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

	bookPlanUpsertCacheMut.RLock()
	cache, cached := bookPlanUpsertCache[key]
	bookPlanUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookPlanAllColumns,
			bookPlanColumnsWithDefault,
			bookPlanColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bookPlanAllColumns,
			bookPlanPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert book_plan, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bookPlanPrimaryKeyColumns))
			copy(conflict, bookPlanPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"book_plan\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bookPlanType, bookPlanMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookPlanType, bookPlanMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert book_plan")
	}

	if !cached {
		bookPlanUpsertCacheMut.Lock()
		bookPlanUpsertCache[key] = cache
		bookPlanUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BookPlan record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BookPlan) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BookPlan provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookPlanPrimaryKeyMapping)
	sql := "DELETE FROM \"book_plan\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from book_plan")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for book_plan")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookPlanQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bookPlanQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from book_plan")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book_plan")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookPlanSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookPlanBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookPlanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"book_plan\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookPlanPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookPlan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book_plan")
	}

	if len(bookPlanAfterDeleteHooks) != 0 {
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
func (o *BookPlan) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookPlan(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookPlanSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookPlanSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookPlanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"book_plan\".* FROM \"book_plan\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookPlanPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BookPlanSlice")
	}

	*o = slice

	return nil
}

// BookPlanExists checks if the BookPlan row exists.
func BookPlanExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"book_plan\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if book_plan exists")
	}

	return exists, nil
}

// Exists checks if the BookPlan row exists.
func (o *BookPlan) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BookPlanExists(ctx, exec, o.ID)
}
