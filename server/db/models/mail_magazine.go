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

// MailMagazine is an object representing the database table.
type MailMagazine struct {
	ID                 string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title              string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	Content            string      `boil:"content" json:"content" toml:"content" yaml:"content"`
	Author             null.String `boil:"author" json:"author,omitempty" toml:"author" yaml:"author,omitempty"`
	SentAt             null.Time   `boil:"sent_at" json:"sent_at,omitempty" toml:"sent_at" yaml:"sent_at,omitempty"`
	SentCount          null.Int    `boil:"sent_count" json:"sent_count,omitempty" toml:"sent_count" yaml:"sent_count,omitempty"`
	MailMagazineStatus int         `boil:"mail_magazine_status" json:"mail_magazine_status" toml:"mail_magazine_status" yaml:"mail_magazine_status"`
	CreateAt           time.Time   `boil:"create_at" json:"create_at" toml:"create_at" yaml:"create_at"`
	UpdateAt           time.Time   `boil:"update_at" json:"update_at" toml:"update_at" yaml:"update_at"`

	R *mailMagazineR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mailMagazineL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MailMagazineColumns = struct {
	ID                 string
	Title              string
	Content            string
	Author             string
	SentAt             string
	SentCount          string
	MailMagazineStatus string
	CreateAt           string
	UpdateAt           string
}{
	ID:                 "id",
	Title:              "title",
	Content:            "content",
	Author:             "author",
	SentAt:             "sent_at",
	SentCount:          "sent_count",
	MailMagazineStatus: "mail_magazine_status",
	CreateAt:           "create_at",
	UpdateAt:           "update_at",
}

var MailMagazineTableColumns = struct {
	ID                 string
	Title              string
	Content            string
	Author             string
	SentAt             string
	SentCount          string
	MailMagazineStatus string
	CreateAt           string
	UpdateAt           string
}{
	ID:                 "mail_magazine.id",
	Title:              "mail_magazine.title",
	Content:            "mail_magazine.content",
	Author:             "mail_magazine.author",
	SentAt:             "mail_magazine.sent_at",
	SentCount:          "mail_magazine.sent_count",
	MailMagazineStatus: "mail_magazine.mail_magazine_status",
	CreateAt:           "mail_magazine.create_at",
	UpdateAt:           "mail_magazine.update_at",
}

// Generated where

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var MailMagazineWhere = struct {
	ID                 whereHelperstring
	Title              whereHelperstring
	Content            whereHelperstring
	Author             whereHelpernull_String
	SentAt             whereHelpernull_Time
	SentCount          whereHelpernull_Int
	MailMagazineStatus whereHelperint
	CreateAt           whereHelpertime_Time
	UpdateAt           whereHelpertime_Time
}{
	ID:                 whereHelperstring{field: "\"mail_magazine\".\"id\""},
	Title:              whereHelperstring{field: "\"mail_magazine\".\"title\""},
	Content:            whereHelperstring{field: "\"mail_magazine\".\"content\""},
	Author:             whereHelpernull_String{field: "\"mail_magazine\".\"author\""},
	SentAt:             whereHelpernull_Time{field: "\"mail_magazine\".\"sent_at\""},
	SentCount:          whereHelpernull_Int{field: "\"mail_magazine\".\"sent_count\""},
	MailMagazineStatus: whereHelperint{field: "\"mail_magazine\".\"mail_magazine_status\""},
	CreateAt:           whereHelpertime_Time{field: "\"mail_magazine\".\"create_at\""},
	UpdateAt:           whereHelpertime_Time{field: "\"mail_magazine\".\"update_at\""},
}

// MailMagazineRels is where relationship names are stored.
var MailMagazineRels = struct {
	AuthorAdmin string
}{
	AuthorAdmin: "AuthorAdmin",
}

// mailMagazineR is where relationships are stored.
type mailMagazineR struct {
	AuthorAdmin *Admin `boil:"AuthorAdmin" json:"AuthorAdmin" toml:"AuthorAdmin" yaml:"AuthorAdmin"`
}

// NewStruct creates a new relationship struct
func (*mailMagazineR) NewStruct() *mailMagazineR {
	return &mailMagazineR{}
}

func (r *mailMagazineR) GetAuthorAdmin() *Admin {
	if r == nil {
		return nil
	}
	return r.AuthorAdmin
}

// mailMagazineL is where Load methods for each relationship are stored.
type mailMagazineL struct{}

var (
	mailMagazineAllColumns            = []string{"id", "title", "content", "author", "sent_at", "sent_count", "mail_magazine_status", "create_at", "update_at"}
	mailMagazineColumnsWithoutDefault = []string{"id", "title", "content", "mail_magazine_status"}
	mailMagazineColumnsWithDefault    = []string{"author", "sent_at", "sent_count", "create_at", "update_at"}
	mailMagazinePrimaryKeyColumns     = []string{"id"}
	mailMagazineGeneratedColumns      = []string{}
)

type (
	// MailMagazineSlice is an alias for a slice of pointers to MailMagazine.
	// This should almost always be used instead of []MailMagazine.
	MailMagazineSlice []*MailMagazine
	// MailMagazineHook is the signature for custom MailMagazine hook methods
	MailMagazineHook func(context.Context, boil.ContextExecutor, *MailMagazine) error

	mailMagazineQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mailMagazineType                 = reflect.TypeOf(&MailMagazine{})
	mailMagazineMapping              = queries.MakeStructMapping(mailMagazineType)
	mailMagazinePrimaryKeyMapping, _ = queries.BindMapping(mailMagazineType, mailMagazineMapping, mailMagazinePrimaryKeyColumns)
	mailMagazineInsertCacheMut       sync.RWMutex
	mailMagazineInsertCache          = make(map[string]insertCache)
	mailMagazineUpdateCacheMut       sync.RWMutex
	mailMagazineUpdateCache          = make(map[string]updateCache)
	mailMagazineUpsertCacheMut       sync.RWMutex
	mailMagazineUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mailMagazineAfterSelectHooks []MailMagazineHook

var mailMagazineBeforeInsertHooks []MailMagazineHook
var mailMagazineAfterInsertHooks []MailMagazineHook

var mailMagazineBeforeUpdateHooks []MailMagazineHook
var mailMagazineAfterUpdateHooks []MailMagazineHook

var mailMagazineBeforeDeleteHooks []MailMagazineHook
var mailMagazineAfterDeleteHooks []MailMagazineHook

var mailMagazineBeforeUpsertHooks []MailMagazineHook
var mailMagazineAfterUpsertHooks []MailMagazineHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MailMagazine) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MailMagazine) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MailMagazine) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MailMagazine) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MailMagazine) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MailMagazine) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MailMagazine) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MailMagazine) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MailMagazine) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMailMagazineHook registers your hook function for all future operations.
func AddMailMagazineHook(hookPoint boil.HookPoint, mailMagazineHook MailMagazineHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		mailMagazineAfterSelectHooks = append(mailMagazineAfterSelectHooks, mailMagazineHook)
	case boil.BeforeInsertHook:
		mailMagazineBeforeInsertHooks = append(mailMagazineBeforeInsertHooks, mailMagazineHook)
	case boil.AfterInsertHook:
		mailMagazineAfterInsertHooks = append(mailMagazineAfterInsertHooks, mailMagazineHook)
	case boil.BeforeUpdateHook:
		mailMagazineBeforeUpdateHooks = append(mailMagazineBeforeUpdateHooks, mailMagazineHook)
	case boil.AfterUpdateHook:
		mailMagazineAfterUpdateHooks = append(mailMagazineAfterUpdateHooks, mailMagazineHook)
	case boil.BeforeDeleteHook:
		mailMagazineBeforeDeleteHooks = append(mailMagazineBeforeDeleteHooks, mailMagazineHook)
	case boil.AfterDeleteHook:
		mailMagazineAfterDeleteHooks = append(mailMagazineAfterDeleteHooks, mailMagazineHook)
	case boil.BeforeUpsertHook:
		mailMagazineBeforeUpsertHooks = append(mailMagazineBeforeUpsertHooks, mailMagazineHook)
	case boil.AfterUpsertHook:
		mailMagazineAfterUpsertHooks = append(mailMagazineAfterUpsertHooks, mailMagazineHook)
	}
}

// One returns a single mailMagazine record from the query.
func (q mailMagazineQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MailMagazine, error) {
	o := &MailMagazine{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mail_magazine")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MailMagazine records from the query.
func (q mailMagazineQuery) All(ctx context.Context, exec boil.ContextExecutor) (MailMagazineSlice, error) {
	var o []*MailMagazine

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MailMagazine slice")
	}

	if len(mailMagazineAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MailMagazine records in the query.
func (q mailMagazineQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mail_magazine rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mailMagazineQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mail_magazine exists")
	}

	return count > 0, nil
}

// AuthorAdmin pointed to by the foreign key.
func (o *MailMagazine) AuthorAdmin(mods ...qm.QueryMod) adminQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.Author),
	}

	queryMods = append(queryMods, mods...)

	return Admins(queryMods...)
}

// LoadAuthorAdmin allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (mailMagazineL) LoadAuthorAdmin(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMailMagazine interface{}, mods queries.Applicator) error {
	var slice []*MailMagazine
	var object *MailMagazine

	if singular {
		var ok bool
		object, ok = maybeMailMagazine.(*MailMagazine)
		if !ok {
			object = new(MailMagazine)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMailMagazine)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMailMagazine))
			}
		}
	} else {
		s, ok := maybeMailMagazine.(*[]*MailMagazine)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMailMagazine)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMailMagazine))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &mailMagazineR{}
		}
		if !queries.IsNil(object.Author) {
			args = append(args, object.Author)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &mailMagazineR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.Author) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.Author) {
				args = append(args, obj.Author)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`admin`),
		qm.WhereIn(`admin.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Admin")
	}

	var resultSlice []*Admin
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Admin")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for admin")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for admin")
	}

	if len(adminAfterSelectHooks) != 0 {
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
		object.R.AuthorAdmin = foreign
		if foreign.R == nil {
			foreign.R = &adminR{}
		}
		foreign.R.AuthorMailMagazines = append(foreign.R.AuthorMailMagazines, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.Author, foreign.ID) {
				local.R.AuthorAdmin = foreign
				if foreign.R == nil {
					foreign.R = &adminR{}
				}
				foreign.R.AuthorMailMagazines = append(foreign.R.AuthorMailMagazines, local)
				break
			}
		}
	}

	return nil
}

// SetAuthorAdmin of the mailMagazine to the related item.
// Sets o.R.AuthorAdmin to related.
// Adds o to related.R.AuthorMailMagazines.
func (o *MailMagazine) SetAuthorAdmin(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Admin) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"mail_magazine\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"author"}),
		strmangle.WhereClause("\"", "\"", 2, mailMagazinePrimaryKeyColumns),
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

	queries.Assign(&o.Author, related.ID)
	if o.R == nil {
		o.R = &mailMagazineR{
			AuthorAdmin: related,
		}
	} else {
		o.R.AuthorAdmin = related
	}

	if related.R == nil {
		related.R = &adminR{
			AuthorMailMagazines: MailMagazineSlice{o},
		}
	} else {
		related.R.AuthorMailMagazines = append(related.R.AuthorMailMagazines, o)
	}

	return nil
}

// RemoveAuthorAdmin relationship.
// Sets o.R.AuthorAdmin to nil.
// Removes o from all passed in related items' relationships struct.
func (o *MailMagazine) RemoveAuthorAdmin(ctx context.Context, exec boil.ContextExecutor, related *Admin) error {
	var err error

	queries.SetScanner(&o.Author, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("author")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.AuthorAdmin = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.AuthorMailMagazines {
		if queries.Equal(o.Author, ri.Author) {
			continue
		}

		ln := len(related.R.AuthorMailMagazines)
		if ln > 1 && i < ln-1 {
			related.R.AuthorMailMagazines[i] = related.R.AuthorMailMagazines[ln-1]
		}
		related.R.AuthorMailMagazines = related.R.AuthorMailMagazines[:ln-1]
		break
	}
	return nil
}

// MailMagazines retrieves all the records using an executor.
func MailMagazines(mods ...qm.QueryMod) mailMagazineQuery {
	mods = append(mods, qm.From("\"mail_magazine\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"mail_magazine\".*"})
	}

	return mailMagazineQuery{q}
}

// FindMailMagazine retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMailMagazine(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*MailMagazine, error) {
	mailMagazineObj := &MailMagazine{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mail_magazine\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mailMagazineObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mail_magazine")
	}

	if err = mailMagazineObj.doAfterSelectHooks(ctx, exec); err != nil {
		return mailMagazineObj, err
	}

	return mailMagazineObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MailMagazine) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mail_magazine provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(mailMagazineColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mailMagazineInsertCacheMut.RLock()
	cache, cached := mailMagazineInsertCache[key]
	mailMagazineInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mailMagazineAllColumns,
			mailMagazineColumnsWithDefault,
			mailMagazineColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mailMagazineType, mailMagazineMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mailMagazineType, mailMagazineMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mail_magazine\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mail_magazine\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into mail_magazine")
	}

	if !cached {
		mailMagazineInsertCacheMut.Lock()
		mailMagazineInsertCache[key] = cache
		mailMagazineInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MailMagazine.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MailMagazine) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdateAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mailMagazineUpdateCacheMut.RLock()
	cache, cached := mailMagazineUpdateCache[key]
	mailMagazineUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mailMagazineAllColumns,
			mailMagazinePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mail_magazine, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mail_magazine\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mailMagazinePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mailMagazineType, mailMagazineMapping, append(wl, mailMagazinePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update mail_magazine row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mail_magazine")
	}

	if !cached {
		mailMagazineUpdateCacheMut.Lock()
		mailMagazineUpdateCache[key] = cache
		mailMagazineUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mailMagazineQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mail_magazine")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mail_magazine")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MailMagazineSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mailMagazinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mail_magazine\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mailMagazinePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mailMagazine slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mailMagazine")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MailMagazine) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mail_magazine provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(mailMagazineColumnsWithDefault, o)

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

	mailMagazineUpsertCacheMut.RLock()
	cache, cached := mailMagazineUpsertCache[key]
	mailMagazineUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mailMagazineAllColumns,
			mailMagazineColumnsWithDefault,
			mailMagazineColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			mailMagazineAllColumns,
			mailMagazinePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert mail_magazine, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(mailMagazinePrimaryKeyColumns))
			copy(conflict, mailMagazinePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mail_magazine\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(mailMagazineType, mailMagazineMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mailMagazineType, mailMagazineMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert mail_magazine")
	}

	if !cached {
		mailMagazineUpsertCacheMut.Lock()
		mailMagazineUpsertCache[key] = cache
		mailMagazineUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MailMagazine record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MailMagazine) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MailMagazine provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mailMagazinePrimaryKeyMapping)
	sql := "DELETE FROM \"mail_magazine\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mail_magazine")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mail_magazine")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mailMagazineQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mailMagazineQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mail_magazine")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mail_magazine")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MailMagazineSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(mailMagazineBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mailMagazinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mail_magazine\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mailMagazinePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mailMagazine slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mail_magazine")
	}

	if len(mailMagazineAfterDeleteHooks) != 0 {
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
func (o *MailMagazine) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMailMagazine(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MailMagazineSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MailMagazineSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mailMagazinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mail_magazine\".* FROM \"mail_magazine\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mailMagazinePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MailMagazineSlice")
	}

	*o = slice

	return nil
}

// MailMagazineExists checks if the MailMagazine row exists.
func MailMagazineExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mail_magazine\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mail_magazine exists")
	}

	return exists, nil
}

// Exists checks if the MailMagazine row exists.
func (o *MailMagazine) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MailMagazineExists(ctx, exec, o.ID)
}
