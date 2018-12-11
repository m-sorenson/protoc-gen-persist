// This file is generated by protoc-gen-persist
// Source File: pb/user.proto
// DO NOT EDIT !

// TXs, Queries, Hooks, TypeMappings, Handlers, Rows, Iters
package pb

import (
	sql "database/sql"
	"database/sql/driver"
	"fmt"
	io "io"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
	codes "google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

// WriteHandlers
// RestOf<S>Handlers
type RestOfUServHandlers interface {
	UpdateAllNames(req *Empty, stream UServ_UpdateAllNamesServer) error
}

// WriteTypeMappigns
type UServTypeMappings interface {
	TimestampTimestamp() TimestampTimestampMappingImpl
	SliceStringParam() SliceStringParamMappingImpl
}

type TimestampTimestampMappingImpl interface {
	ToProto(**timestamp.Timestamp) error
	Empty() TimestampTimestampMappingImpl
	ToSql(*timestamp.Timestamp) sql.Scanner
	sql.Scanner
	driver.Valuer
}
type SliceStringParamMappingImpl interface {
	ToProto(**SliceStringParam) error
	Empty() SliceStringParamMappingImpl
	ToSql(*SliceStringParam) sql.Scanner
	sql.Scanner
	driver.Valuer
}
type UServHooks interface {
	InsertUsersBeforeHook(*User) (*Empty, error)
	InsertUsersAfterHook(*User, *Empty) error
	GetAllUsersBeforeHook(*Empty) ([]*User, error)
	GetAllUsersAfterHook(*Empty, *User) error
}
type alwaysScanner struct {
	i *interface{}
}

func (s *alwaysScanner) Scan(src interface{}) error {
	s.i = &src
	return nil
}

type scanable interface {
	Scan(...interface{}) error
	Columns() ([]string, error)
}
type Runable interface {
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

func DefaultClientStreamingPersistTx(ctx context.Context, db *sql.DB) (PersistTx, error) {
	return db.BeginTx(ctx, nil)
}
func DefaultServerStreamingPersistTx(ctx context.Context, db *sql.DB) (PersistTx, error) {
	return NopPersistTx(db)
}
func DefaultBidiStreamingPersistTx(ctx context.Context, db *sql.DB) (PersistTx, error) {
	return NopPersistTx(db)
}
func DefaultUnaryPersistTx(ctx context.Context, db *sql.DB) (PersistTx, error) {
	return NopPersistTx(db)
}

type ignoreTx struct {
	r Runable
}

func (this *ignoreTx) Commit() error   { return nil }
func (this *ignoreTx) Rollback() error { return nil }
func (this *ignoreTx) QueryContext(ctx context.Context, x string, ys ...interface{}) (*sql.Rows, error) {
	return this.r.QueryContext(ctx, x, ys...)
}
func (this *ignoreTx) ExecContext(ctx context.Context, x string, ys ...interface{}) (sql.Result, error) {
	return this.r.ExecContext(ctx, x, ys...)
}

type UServ_QueryOpts struct {
	MAPPINGS UServTypeMappings
	db       Runable
	ctx      context.Context
}

func DefaultUServQueryOpts(db Runable) UServ_QueryOpts {
	return UServ_QueryOpts{
		db: db,
	}
}

type UServ_Queries struct {
	opts UServ_QueryOpts
}
type PersistTx interface {
	Commit() error
	Rollback() error
	Runable
}

func NopPersistTx(r Runable) (PersistTx, error) {
	return &ignoreTx{r}, nil
}

type UServ_InsertUsersOut interface {
	GetId() int64
	GetName() string
	GetFriends() *Friends
	GetCreatedOn() *timestamp.Timestamp
}

type UServ_InsertUsersRow struct {
	item UServ_InsertUsersOut
	err  error
}

func newUServ_InsertUsersRow(item UServ_InsertUsersOut, err error) *UServ_InsertUsersRow {
	return &UServ_InsertUsersRow{item, err}
}

// Unwrap takes an address to a proto.Message as its only parameter
// Unwrap can only set into output protos of that match method return types + the out option on the query itself
func (this *UServ_InsertUsersRow) Unwrap(pointerToMsg proto.Message) error {
	if this.err != nil {
		return this.err
	}
	// for each known method result
	if o, ok := (pointerToMsg).(*User); ok {
		if o == nil {
			return fmt.Errorf("must initialize *User before giving to Unwrap()")
		}
		res, _ := this.User()
		// set shared fields
		o.Id = res.Id
		o.Name = res.Name
		o.Friends = res.Friends
		o.CreatedOn = res.CreatedOn
		return nil
	}
	if o, ok := (pointerToMsg).(*Friends); ok {
		if o == nil {
			return fmt.Errorf("must initialize *Friends before giving to Unwrap()")
		}
	}

	return nil
}

// one for each Output type of the methods that use this query + the output proto itself

func (this *UServ_InsertUsersRow) User() (*User, error) {
	return &User{
		Id:        this.item.GetId(),
		Name:      this.item.GetName(),
		Friends:   this.item.GetFriends(),
		CreatedOn: this.item.GetCreatedOn(),
	}, nil
}

// just for example
func (this *UServ_InsertUsersRow) Friends() (*Friends, error) {
	return nil, nil
}

// UServPersistQueries returns all the known 'SQL' queires for the 'UServ' service.
func UServPersistQueries(db Runable, opts ...UServ_QueryOpts) *UServ_Queries {
	var myOpts UServ_QueryOpts
	if len(opts) > 0 {
		myOpts = opts[0]
	} else {
		myOpts = DefaultUServQueryOpts(db)
	}
	return &UServ_Queries{
		opts: myOpts,
	}
}

// camel case the services query name
// method for every query

// InsertUsersQuery returns a new struct wrapping the current UServ_QueryOpts
// that will perform 'UServ' services 'insert_users_query' on the database
// when executed
func (this *UServ_Queries) InsertUsersQuery(ctx context.Context) *UServ_InsertUsersQuery {
	return &UServ_InsertUsersQuery{
		opts: UServ_QueryOpts{
			MAPPINGS: this.opts.MAPPINGS,
			db:       this.opts.db,
			ctx:      ctx,
		},
	}
}

// I dont know this is a insert query, I only know this is a query
type UServ_InsertUsersQuery struct {
	opts UServ_QueryOpts
	ctx  context.Context
}

func (this *UServ_InsertUsersQuery) QueryInTypeUser()  {}
func (this *UServ_InsertUsersQuery) QueryOutTypeUser() {}

// the main execute function
func (this *UServ_InsertUsersQuery) Execute(x UServ_InsertUsersOut) *UServ_InsertUsersIter {
	var setupErr error
	params := []interface{}{
		func() (out interface{}) {
			out = x.GetId()
			return
		}(),
		func() (out interface{}) {
			out = x.GetName()
			return
		}(),
		func() (out interface{}) {
			raw, err := proto.Marshal(x.GetFriends())
			if err != nil {
				setupErr = err
			}
			out = raw
			return
		}(),
		func() (out interface{}) {
			mapper := this.opts.MAPPINGS.TimestampTimestamp()
			out = mapper.ToSql(x.GetCreatedOn())
			return
		}(),
	}
	result := &UServ_InsertUsersIter{
		tm:  this.opts.MAPPINGS,
		ctx: this.ctx,
	}
	if setupErr != nil {
		result.err = setupErr
		return result
	}
	result.result, result.err = this.opts.db.ExecContext(this.ctx, "INSERT INTO users (id, name, friends, created_on) VALUES ($1, $2, $3, $4)", params...)

	return result
}

//<SERVICE><QUERY (camel)><MESSAGE TYPE>Iter
type UServ_InsertUsersIter struct {
	result sql.Result
	rows   *sql.Rows
	err    error
	tm     UServTypeMappings
	ctx    context.Context
}

func (this *UServ_InsertUsersIter) IterOutTypeUser() {}
func (this *UServ_InsertUsersIter) IterInTypeUser()  {}

// Each performs 'fun' on each row in the result set.
// Each respects the context passed to it.
// It will stop iteration, and returns ctx.Err() if encountered.
func (this *UServ_InsertUsersIter) Each(fun func(*UServ_InsertUsersRow) error) error {
	for {
		select {
		case <-this.ctx.Done():
			return this.ctx.Err()
		default:
			if row, ok := this.Next(); !ok {
				return nil
			} else if err := fun(row); err != nil {
				return err
			}
		}
	}
	return nil
}

// One returns the sole row, or ensures an error if there was not one result when this row is converted
func (this *UServ_InsertUsersIter) One() *UServ_InsertUsersRow {
	first, hasFirst := this.Next()
	_, hasSecond := this.Next()
	if !hasFirst || hasSecond {
		return newUServ_InsertUsersRow(first.item, fmt.Errorf("expected exactly 1 result from query 'InsertUsers'"))
	}
	return first
}

// Zero returns an error if there were any rows in the result
func (this *UServ_InsertUsersIter) Zero() error {
	if _, ok := this.Next(); ok {
		return fmt.Errorf("expected exactly 0 results from query 'InsertUsers'")
	}
	return nil
}

// Next returns the next scanned row out of the database, or (nil, false) if there are no more rows
func (this *UServ_InsertUsersIter) Next() (*UServ_InsertUsersRow, bool) {
	if this.rows == nil || this.err == io.EOF {
		return nil, false
	} else if this.err != nil {
		err := this.err
		this.err = io.EOF
		return &UServ_InsertUsersRow{err: err}, true
	}
	cols, err := this.rows.Columns()
	if err != nil {
		return &UServ_InsertUsersRow{err: err}, true
	}
	toScan := make([]interface{}, len(cols))
	scanned := make([]alwaysScanner, len(cols))
	for i := range scanned {
		toScan[i] = &scanned[i]
	}
	if this.err = this.rows.Scan(toScan...); this.err != nil {
		return &UServ_InsertUsersRow{err: err}, true
	}
	// this sets the error for next time, after this point,
	// we do not return this.err
	if !this.rows.Next() {
		if this.err = this.rows.Err(); this.err == nil {
			this.err = io.EOF
			return nil, false
		}
	}
	res := &User{}
	for i, col := range cols {
		_ = i
		switch col {
		case "id":
			r, ok := (*scanned[i].i).(int64)
			if !ok {
				return &UServ_InsertUsersRow{err: fmt.Errorf("cant convert db column id to protobuf go type int64")}, true
			}
			res.Id = r
		case "name":
			r, ok := (*scanned[i].i).(string)
			if !ok {
				return &UServ_InsertUsersRow{err: fmt.Errorf("cant convert db column name to protobuf go type string")}, true
			}
			res.Name = r
		case "friends":
			r, ok := (*scanned[i].i).([]byte)
			if !ok {
				return &UServ_InsertUsersRow{err: fmt.Errorf("cant convert db column friends to protobuf go type *Friends")}, true
			}
			var converted = new(Friends)
			if err := proto.Unmarshal(r, converted); err != nil {
				return &UServ_InsertUsersRow{err: err}, true
			}
			res.Friends = converted
		case "created_on":
			var converted = this.tm.TimestampTimestamp().Empty()
			if err := converted.Scan(*scanned[i].i); err != nil {
				return &UServ_InsertUsersRow{err: fmt.Errorf("could not convert mapped db column created_on to type on User.CreatedOn: %v", err)}, true
			}
			if err := converted.ToProto(&res.CreatedOn); err != nil {
				return &UServ_InsertUsersRow{err: fmt.Errorf("could not convert mapped db column created_on to type on User.CreatedOn: %v", err)}, true
			}
		default:
			return &UServ_InsertUsersRow{err: fmt.Errorf("unsupported column in output: %s", col)}, true
		}
	}
	return &UServ_InsertUsersRow{item: res}, true
}

// Slice returns all rows found in the iterator as a Slice.
func (this *UServ_InsertUsersIter) Slice() []*UServ_InsertUsersRow {
	var results []*UServ_InsertUsersRow
	for {
		if i, ok := this.Next(); ok {
			results = append(results, i)
		} else {
			break
		}
	}
	return results
}

// returns the known columns for this result
func (r *UServ_InsertUsersIter) Columns() ([]string, error) {
	if r.err != nil {
		return nil, r.err
	}
	if r.rows != nil {
		return r.rows.Columns()
	}
	return nil, nil
}

type UServ_ImplOpts struct {
	MAPPINGS UServTypeMappings
	HOOKS    UServHooks
}

func DefaultUServImplOpts() UServ_ImplOpts {
	return UServ_ImplOpts{}
}

type UServ_Impl struct {
	opts    *UServ_ImplOpts
	QUERIES *UServ_Queries
	DB      *sql.DB
}

func UServPersistImpl(db *sql.DB, opts ...UServ_ImplOpts) *UServ_Impl {
	var myOpts UServ_ImplOpts
	if len(opts) > 0 {
		myOpts = opts[0]
	} else {
		myOpts = DefaultUServImplOpts()
	}
	return &UServ_Impl{
		opts:    &myOpts,
		QUERIES: UServPersistQueries(db, UServ_QueryOpts{MAPPINGS: myOpts.MAPPINGS}),
		DB:      db,
	}
}

// THIS is the grpc handler
func (this *UServ_Impl) InsertUsers(stream UServ_InsertUsersServer) error {
	tx, err := DefaultClientStreamingPersistTx(stream.Context(), this.DB)
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error creating persist tx: %v", err)
	}
	if err := this.InsertUsersTx(stream, tx); err != nil {
		return gstatus.Errorf(codes.Unknown, "error executing 'insert_users' query: %v", err)
	}
	return nil
}

func (this *UServ_Impl) InsertUsersTx(stream UServ_InsertUsersServer, tx PersistTx) error {
	query := this.QUERIES.InsertUsersQuery(stream.Context())
	var first *User
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return gstatus.Errorf(codes.Unknown, "error receiving request: %v", err)
		}
		if first == nil {
			first = req
		}
		// TODO UPDATE HOOKS FOR CTX
		beforeRes, err := this.opts.HOOKS.InsertUsersBeforeHook( /*stream.Context(),*/ req)
		if err != nil {
			return gstatus.Errorf(codes.Unknown, "error in before hook: %v", err)
		} else if beforeRes != nil {
			continue
		}
		result := query.Execute(req)
		/*for {
			res := new(User)
			if row, ok := result.Next(); !ok {
				break
			} else if err := row.Unwrap(res); err != nil {
				return err
			} else {
				stream.Send(res)
			}
		}
		err := result.Each(stream.Context(), func(row *UServ_InsertUsersRow) error {
			res, err := row.User()
			if err != nil {
				return err
			}
			return stream.Send(res)
		})
		users := result.Slice()
		for _, row := range users {
			user, err := row.User()
			if err != nil {
				return err
			}
			stream.Send(user)
		}
		res, err := result.One().Friends()
		err := result.Zero()*/
		// TODO allow results to be returned here?
		if err := result.Zero(); err != nil {
			return gstatus.Errorf(codes.InvalidArgument, "client streaming queries must return zero results")
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("executed 'insert_users' query without error, but received error on commit: %v", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("error executing 'insert_users' query :::AND COULD NOT ROLLBACK::: rollback err: %v, query err: %v", rollbackErr, err)
		}
	}
	res := &Empty{}
	if err := this.opts.HOOKS.InsertUsersAfterHook( /*stream.Context(),*/ first, res); err != nil {
		return gstatus.Errorf(codes.Unknown, "error in after hook: %v", err)
	}
	if err := stream.SendAndClose(res); err != nil {
		return gstatus.Errorf(codes.Unknown, "error sending back response: %v", err)
	}

	return nil
}
