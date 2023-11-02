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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// BookGuestDatum is an object representing the database table.
type BookGuestDatum struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	FirstName     string      `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	LastName      string      `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	FirstNameKana string      `boil:"first_name_kana" json:"first_name_kana" toml:"first_name_kana" yaml:"first_name_kana"`
	LastNameKana  string      `boil:"last_name_kana" json:"last_name_kana" toml:"last_name_kana" yaml:"last_name_kana"`
	CompanyName   null.String `boil:"company_name" json:"company_name,omitempty" toml:"company_name" yaml:"company_name,omitempty"`
	BirthDate     time.Time   `boil:"birth_date" json:"birth_date" toml:"birth_date" yaml:"birth_date"`
	ZipCode       null.String `boil:"zip_code" json:"zip_code,omitempty" toml:"zip_code" yaml:"zip_code,omitempty"`
	Prefecture    int         `boil:"prefecture" json:"prefecture" toml:"prefecture" yaml:"prefecture"`
	City          null.String `boil:"city" json:"city,omitempty" toml:"city" yaml:"city,omitempty"`
	Address       null.String `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Tel           null.String `boil:"tel" json:"tel,omitempty" toml:"tel" yaml:"tel,omitempty"`
	Mail          string      `boil:"mail" json:"mail" toml:"mail" yaml:"mail"`
	CreateAt      time.Time   `boil:"create_at" json:"create_at" toml:"create_at" yaml:"create_at"`
	UpdateAt      time.Time   `boil:"update_at" json:"update_at" toml:"update_at" yaml:"update_at"`

	R *bookGuestDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookGuestDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookGuestDatumColumns = struct {
	ID            string
	FirstName     string
	LastName      string
	FirstNameKana string
	LastNameKana  string
	CompanyName   string
	BirthDate     string
	ZipCode       string
	Prefecture    string
	City          string
	Address       string
	Tel           string
	Mail          string
	CreateAt      string
	UpdateAt      string
}{
	ID:            "id",
	FirstName:     "first_name",
	LastName:      "last_name",
	FirstNameKana: "first_name_kana",
	LastNameKana:  "last_name_kana",
	CompanyName:   "company_name",
	BirthDate:     "birth_date",
	ZipCode:       "zip_code",
	Prefecture:    "prefecture",
	City:          "city",
	Address:       "address",
	Tel:           "tel",
	Mail:          "mail",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
}

var BookGuestDatumTableColumns = struct {
	ID            string
	FirstName     string
	LastName      string
	FirstNameKana string
	LastNameKana  string
	CompanyName   string
	BirthDate     string
	ZipCode       string
	Prefecture    string
	City          string
	Address       string
	Tel           string
	Mail          string
	CreateAt      string
	UpdateAt      string
}{
	ID:            "book_guest_data.id",
	FirstName:     "book_guest_data.first_name",
	LastName:      "book_guest_data.last_name",
	FirstNameKana: "book_guest_data.first_name_kana",
	LastNameKana:  "book_guest_data.last_name_kana",
	CompanyName:   "book_guest_data.company_name",
	BirthDate:     "book_guest_data.birth_date",
	ZipCode:       "book_guest_data.zip_code",
	Prefecture:    "book_guest_data.prefecture",
	City:          "book_guest_data.city",
	Address:       "book_guest_data.address",
	Tel:           "book_guest_data.tel",
	Mail:          "book_guest_data.mail",
	CreateAt:      "book_guest_data.create_at",
	UpdateAt:      "book_guest_data.update_at",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

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

var BookGuestDatumWhere = struct {
	ID            whereHelperstring
	FirstName     whereHelperstring
	LastName      whereHelperstring
	FirstNameKana whereHelperstring
	LastNameKana  whereHelperstring
	CompanyName   whereHelpernull_String
	BirthDate     whereHelpertime_Time
	ZipCode       whereHelpernull_String
	Prefecture    whereHelperint
	City          whereHelpernull_String
	Address       whereHelpernull_String
	Tel           whereHelpernull_String
	Mail          whereHelperstring
	CreateAt      whereHelpertime_Time
	UpdateAt      whereHelpertime_Time
}{
	ID:            whereHelperstring{field: "\"book_guest_data\".\"id\""},
	FirstName:     whereHelperstring{field: "\"book_guest_data\".\"first_name\""},
	LastName:      whereHelperstring{field: "\"book_guest_data\".\"last_name\""},
	FirstNameKana: whereHelperstring{field: "\"book_guest_data\".\"first_name_kana\""},
	LastNameKana:  whereHelperstring{field: "\"book_guest_data\".\"last_name_kana\""},
	CompanyName:   whereHelpernull_String{field: "\"book_guest_data\".\"company_name\""},
	BirthDate:     whereHelpertime_Time{field: "\"book_guest_data\".\"birth_date\""},
	ZipCode:       whereHelpernull_String{field: "\"book_guest_data\".\"zip_code\""},
	Prefecture:    whereHelperint{field: "\"book_guest_data\".\"prefecture\""},
	City:          whereHelpernull_String{field: "\"book_guest_data\".\"city\""},
	Address:       whereHelpernull_String{field: "\"book_guest_data\".\"address\""},
	Tel:           whereHelpernull_String{field: "\"book_guest_data\".\"tel\""},
	Mail:          whereHelperstring{field: "\"book_guest_data\".\"mail\""},
	CreateAt:      whereHelpertime_Time{field: "\"book_guest_data\".\"create_at\""},
	UpdateAt:      whereHelpertime_Time{field: "\"book_guest_data\".\"update_at\""},
}

// BookGuestDatumRels is where relationship names are stored.
var BookGuestDatumRels = struct {
	GuestDatumUserBooks string
}{
	GuestDatumUserBooks: "GuestDatumUserBooks",
}

// bookGuestDatumR is where relationships are stored.
type bookGuestDatumR struct {
	GuestDatumUserBooks UserBookSlice `boil:"GuestDatumUserBooks" json:"GuestDatumUserBooks" toml:"GuestDatumUserBooks" yaml:"GuestDatumUserBooks"`
}

// NewStruct creates a new relationship struct
func (*bookGuestDatumR) NewStruct() *bookGuestDatumR {
	return &bookGuestDatumR{}
}

func (r *bookGuestDatumR) GetGuestDatumUserBooks() UserBookSlice {
	if r == nil {
		return nil
	}
	return r.GuestDatumUserBooks
}

// bookGuestDatumL is where Load methods for each relationship are stored.
type bookGuestDatumL struct{}

var (
	bookGuestDatumAllColumns            = []string{"id", "first_name", "last_name", "first_name_kana", "last_name_kana", "company_name", "birth_date", "zip_code", "prefecture", "city", "address", "tel", "mail", "create_at", "update_at"}
	bookGuestDatumColumnsWithoutDefault = []string{"id", "first_name", "last_name", "first_name_kana", "last_name_kana", "birth_date", "prefecture", "mail"}
	bookGuestDatumColumnsWithDefault    = []string{"company_name", "zip_code", "city", "address", "tel", "create_at", "update_at"}
	bookGuestDatumPrimaryKeyColumns     = []string{"id"}
	bookGuestDatumGeneratedColumns      = []string{}
)

type (
	// BookGuestDatumSlice is an alias for a slice of pointers to BookGuestDatum.
	// This should almost always be used instead of []BookGuestDatum.
	BookGuestDatumSlice []*BookGuestDatum
	// BookGuestDatumHook is the signature for custom BookGuestDatum hook methods
	BookGuestDatumHook func(context.Context, boil.ContextExecutor, *BookGuestDatum) error

	bookGuestDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookGuestDatumType                 = reflect.TypeOf(&BookGuestDatum{})
	bookGuestDatumMapping              = queries.MakeStructMapping(bookGuestDatumType)
	bookGuestDatumPrimaryKeyMapping, _ = queries.BindMapping(bookGuestDatumType, bookGuestDatumMapping, bookGuestDatumPrimaryKeyColumns)
	bookGuestDatumInsertCacheMut       sync.RWMutex
	bookGuestDatumInsertCache          = make(map[string]insertCache)
	bookGuestDatumUpdateCacheMut       sync.RWMutex
	bookGuestDatumUpdateCache          = make(map[string]updateCache)
	bookGuestDatumUpsertCacheMut       sync.RWMutex
	bookGuestDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookGuestDatumAfterSelectHooks []BookGuestDatumHook

var bookGuestDatumBeforeInsertHooks []BookGuestDatumHook
var bookGuestDatumAfterInsertHooks []BookGuestDatumHook

var bookGuestDatumBeforeUpdateHooks []BookGuestDatumHook
var bookGuestDatumAfterUpdateHooks []BookGuestDatumHook

var bookGuestDatumBeforeDeleteHooks []BookGuestDatumHook
var bookGuestDatumAfterDeleteHooks []BookGuestDatumHook

var bookGuestDatumBeforeUpsertHooks []BookGuestDatumHook
var bookGuestDatumAfterUpsertHooks []BookGuestDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BookGuestDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookGuestDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BookGuestDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookGuestDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BookGuestDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookGuestDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BookGuestDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookGuestDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BookGuestDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookGuestDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BookGuestDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookGuestDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BookGuestDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookGuestDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BookGuestDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookGuestDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BookGuestDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookGuestDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookGuestDatumHook registers your hook function for all future operations.
func AddBookGuestDatumHook(hookPoint boil.HookPoint, bookGuestDatumHook BookGuestDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bookGuestDatumAfterSelectHooks = append(bookGuestDatumAfterSelectHooks, bookGuestDatumHook)
	case boil.BeforeInsertHook:
		bookGuestDatumBeforeInsertHooks = append(bookGuestDatumBeforeInsertHooks, bookGuestDatumHook)
	case boil.AfterInsertHook:
		bookGuestDatumAfterInsertHooks = append(bookGuestDatumAfterInsertHooks, bookGuestDatumHook)
	case boil.BeforeUpdateHook:
		bookGuestDatumBeforeUpdateHooks = append(bookGuestDatumBeforeUpdateHooks, bookGuestDatumHook)
	case boil.AfterUpdateHook:
		bookGuestDatumAfterUpdateHooks = append(bookGuestDatumAfterUpdateHooks, bookGuestDatumHook)
	case boil.BeforeDeleteHook:
		bookGuestDatumBeforeDeleteHooks = append(bookGuestDatumBeforeDeleteHooks, bookGuestDatumHook)
	case boil.AfterDeleteHook:
		bookGuestDatumAfterDeleteHooks = append(bookGuestDatumAfterDeleteHooks, bookGuestDatumHook)
	case boil.BeforeUpsertHook:
		bookGuestDatumBeforeUpsertHooks = append(bookGuestDatumBeforeUpsertHooks, bookGuestDatumHook)
	case boil.AfterUpsertHook:
		bookGuestDatumAfterUpsertHooks = append(bookGuestDatumAfterUpsertHooks, bookGuestDatumHook)
	}
}

// One returns a single bookGuestDatum record from the query.
func (q bookGuestDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BookGuestDatum, error) {
	o := &BookGuestDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for book_guest_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BookGuestDatum records from the query.
func (q bookGuestDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookGuestDatumSlice, error) {
	var o []*BookGuestDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BookGuestDatum slice")
	}

	if len(bookGuestDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BookGuestDatum records in the query.
func (q bookGuestDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count book_guest_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bookGuestDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if book_guest_data exists")
	}

	return count > 0, nil
}

// GuestDatumUserBooks retrieves all the user_book's UserBooks with an executor via guest_data_id column.
func (o *BookGuestDatum) GuestDatumUserBooks(mods ...qm.QueryMod) userBookQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_book\".\"guest_data_id\"=?", o.ID),
	)

	return UserBooks(queryMods...)
}

// LoadGuestDatumUserBooks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bookGuestDatumL) LoadGuestDatumUserBooks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBookGuestDatum interface{}, mods queries.Applicator) error {
	var slice []*BookGuestDatum
	var object *BookGuestDatum

	if singular {
		var ok bool
		object, ok = maybeBookGuestDatum.(*BookGuestDatum)
		if !ok {
			object = new(BookGuestDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBookGuestDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBookGuestDatum))
			}
		}
	} else {
		s, ok := maybeBookGuestDatum.(*[]*BookGuestDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBookGuestDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBookGuestDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookGuestDatumR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookGuestDatumR{}
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
		qm.WhereIn(`user_book.guest_data_id in ?`, args...),
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
		object.R.GuestDatumUserBooks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userBookR{}
			}
			foreign.R.GuestDatum = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GuestDataID {
				local.R.GuestDatumUserBooks = append(local.R.GuestDatumUserBooks, foreign)
				if foreign.R == nil {
					foreign.R = &userBookR{}
				}
				foreign.R.GuestDatum = local
				break
			}
		}
	}

	return nil
}

// AddGuestDatumUserBooks adds the given related objects to the existing relationships
// of the book_guest_datum, optionally inserting them as new records.
// Appends related to o.R.GuestDatumUserBooks.
// Sets related.R.GuestDatum appropriately.
func (o *BookGuestDatum) AddGuestDatumUserBooks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserBook) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GuestDataID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_book\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"guest_data_id"}),
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

			rel.GuestDataID = o.ID
		}
	}

	if o.R == nil {
		o.R = &bookGuestDatumR{
			GuestDatumUserBooks: related,
		}
	} else {
		o.R.GuestDatumUserBooks = append(o.R.GuestDatumUserBooks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userBookR{
				GuestDatum: o,
			}
		} else {
			rel.R.GuestDatum = o
		}
	}
	return nil
}

// BookGuestData retrieves all the records using an executor.
func BookGuestData(mods ...qm.QueryMod) bookGuestDatumQuery {
	mods = append(mods, qm.From("\"book_guest_data\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"book_guest_data\".*"})
	}

	return bookGuestDatumQuery{q}
}

// FindBookGuestDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookGuestDatum(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*BookGuestDatum, error) {
	bookGuestDatumObj := &BookGuestDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"book_guest_data\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bookGuestDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from book_guest_data")
	}

	if err = bookGuestDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bookGuestDatumObj, err
	}

	return bookGuestDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BookGuestDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book_guest_data provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(bookGuestDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookGuestDatumInsertCacheMut.RLock()
	cache, cached := bookGuestDatumInsertCache[key]
	bookGuestDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookGuestDatumAllColumns,
			bookGuestDatumColumnsWithDefault,
			bookGuestDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookGuestDatumType, bookGuestDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookGuestDatumType, bookGuestDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"book_guest_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"book_guest_data\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into book_guest_data")
	}

	if !cached {
		bookGuestDatumInsertCacheMut.Lock()
		bookGuestDatumInsertCache[key] = cache
		bookGuestDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BookGuestDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BookGuestDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdateAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookGuestDatumUpdateCacheMut.RLock()
	cache, cached := bookGuestDatumUpdateCache[key]
	bookGuestDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookGuestDatumAllColumns,
			bookGuestDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update book_guest_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"book_guest_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bookGuestDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookGuestDatumType, bookGuestDatumMapping, append(wl, bookGuestDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update book_guest_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for book_guest_data")
	}

	if !cached {
		bookGuestDatumUpdateCacheMut.Lock()
		bookGuestDatumUpdateCache[key] = cache
		bookGuestDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bookGuestDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for book_guest_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for book_guest_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookGuestDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookGuestDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"book_guest_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bookGuestDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bookGuestDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bookGuestDatum")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BookGuestDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book_guest_data provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(bookGuestDatumColumnsWithDefault, o)

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

	bookGuestDatumUpsertCacheMut.RLock()
	cache, cached := bookGuestDatumUpsertCache[key]
	bookGuestDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookGuestDatumAllColumns,
			bookGuestDatumColumnsWithDefault,
			bookGuestDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bookGuestDatumAllColumns,
			bookGuestDatumPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert book_guest_data, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bookGuestDatumPrimaryKeyColumns))
			copy(conflict, bookGuestDatumPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"book_guest_data\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bookGuestDatumType, bookGuestDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookGuestDatumType, bookGuestDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert book_guest_data")
	}

	if !cached {
		bookGuestDatumUpsertCacheMut.Lock()
		bookGuestDatumUpsertCache[key] = cache
		bookGuestDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BookGuestDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BookGuestDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BookGuestDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookGuestDatumPrimaryKeyMapping)
	sql := "DELETE FROM \"book_guest_data\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from book_guest_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for book_guest_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookGuestDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bookGuestDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from book_guest_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book_guest_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookGuestDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookGuestDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookGuestDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"book_guest_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookGuestDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bookGuestDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book_guest_data")
	}

	if len(bookGuestDatumAfterDeleteHooks) != 0 {
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
func (o *BookGuestDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookGuestDatum(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookGuestDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookGuestDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookGuestDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"book_guest_data\".* FROM \"book_guest_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookGuestDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BookGuestDatumSlice")
	}

	*o = slice

	return nil
}

// BookGuestDatumExists checks if the BookGuestDatum row exists.
func BookGuestDatumExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"book_guest_data\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if book_guest_data exists")
	}

	return exists, nil
}

// Exists checks if the BookGuestDatum row exists.
func (o *BookGuestDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BookGuestDatumExists(ctx, exec, o.ID)
}