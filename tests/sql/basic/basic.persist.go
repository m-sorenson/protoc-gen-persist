// This file is generated by protoc-gen-persist
// Source File: tests/sql/basic/basic.proto
// DO NOT EDIT !
package basic

import (
	sql "database/sql"
	driver "database/sql/driver"
	fmt "fmt"
	io "io"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	persist "github.com/tcncloud/protoc-gen-persist/persist"
	test "github.com/tcncloud/protoc-gen-persist/tests/test"
	context "golang.org/x/net/context"
	codes "google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

func NopPersistTx(r persist.Runnable) (persist.PersistTx, error) {
	return &ignoreTx{r}, nil
}

type ignoreTx struct {
	r persist.Runnable
}

func (this *ignoreTx) Commit() error   { return nil }
func (this *ignoreTx) Rollback() error { return nil }
func (this *ignoreTx) QueryContext(ctx context.Context, x string, ys ...interface{}) (*sql.Rows, error) {
	return this.r.QueryContext(ctx, x, ys...)
}
func (this *ignoreTx) ExecContext(ctx context.Context, x string, ys ...interface{}) (sql.Result, error) {
	return this.r.ExecContext(ctx, x, ys...)
}

type Runnable interface {
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

func DefaultClientStreamingPersistTx(ctx context.Context, db *sql.DB) (persist.PersistTx, error) {
	return db.BeginTx(ctx, nil)
}
func DefaultServerStreamingPersistTx(ctx context.Context, db *sql.DB) (persist.PersistTx, error) {
	return NopPersistTx(db)
}
func DefaultBidiStreamingPersistTx(ctx context.Context, db *sql.DB) (persist.PersistTx, error) {
	return NopPersistTx(db)
}
func DefaultUnaryPersistTx(ctx context.Context, db *sql.DB) (persist.PersistTx, error) {
	return NopPersistTx(db)
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

// Queries_Amazing holds all the queries found the proto service option as methods
type Queries_Amazing struct {
	opts Opts_Amazing
}

// QueriesAmazing returns all the known 'SQL' queires for the 'Amazing' service.
// If no opts are provided default implementations are used.
func QueriesAmazing(opts ...Opts_Amazing) *Queries_Amazing {
	var myOpts Opts_Amazing
	if len(opts) > 0 {
		myOpts = opts[0]
	} else {
		myOpts = OptsAmazing(&DefaultHooks_Amazing{}, &DefaultTypeMappings_Amazing{})
	}
	return &Queries_Amazing{
		opts: myOpts,
	}
}

// SelectById returns a struct that will perform the 'select_by_id' query.
// When Execute is called, it will use the following fields:
// [id start_time]
func (this *Queries_Amazing) SelectById(ctx context.Context, db persist.Runnable) *Query_Amazing_SelectById {
	return &Query_Amazing_SelectById{
		opts: this.opts,
		ctx:  ctx,
		db:   db,
	}
}

// Query_Amazing_SelectById (future doc string needed)
type Query_Amazing_SelectById struct {
	opts Opts_Amazing
	db   persist.Runnable
	ctx  context.Context
}

func (this *Query_Amazing_SelectById) QueryInType_PartialTable()  {}
func (this *Query_Amazing_SelectById) QueryOutType_ExampleTable() {}

// Executes the query 'select_by_id' with parameters retrieved from x.
// Fields used: [id start_time]
func (this *Query_Amazing_SelectById) Execute(x In_Amazing_SelectById) *Iter_Amazing_SelectById {
	var setupErr error
	params := []interface{}{
		func() (out interface{}) {
			out = x.GetId()
			return
		}(),
		func() (out interface{}) {
			mapper := this.opts.MAPPINGS.TimestampTimestamp()
			out = mapper.ToSql(x.GetStartTime())
			return
		}(),
	}
	result := &Iter_Amazing_SelectById{
		tm:  this.opts.MAPPINGS,
		ctx: this.ctx,
	}
	if setupErr != nil {
		result.err = setupErr
		return result
	}
	result.rows, result.err = this.db.QueryContext(this.ctx, "SELECT * from example_table Where id=$1 AND start_time>$2", params...)
	return result
}

// SelectByName returns a struct that will perform the 'select_by_name' query.
// When Execute is called, it will use the following fields:
// [name]
func (this *Queries_Amazing) SelectByName(ctx context.Context, db persist.Runnable) *Query_Amazing_SelectByName {
	return &Query_Amazing_SelectByName{
		opts: this.opts,
		ctx:  ctx,
		db:   db,
	}
}

// Query_Amazing_SelectByName (future doc string needed)
type Query_Amazing_SelectByName struct {
	opts Opts_Amazing
	db   persist.Runnable
	ctx  context.Context
}

func (this *Query_Amazing_SelectByName) QueryInType_Name()          {}
func (this *Query_Amazing_SelectByName) QueryOutType_ExampleTable() {}

// Executes the query 'select_by_name' with parameters retrieved from x.
// Fields used: [name]
func (this *Query_Amazing_SelectByName) Execute(x In_Amazing_SelectByName) *Iter_Amazing_SelectByName {
	var setupErr error
	params := []interface{}{
		func() (out interface{}) {
			out = x.GetName()
			return
		}(),
	}
	result := &Iter_Amazing_SelectByName{
		tm:  this.opts.MAPPINGS,
		ctx: this.ctx,
	}
	if setupErr != nil {
		result.err = setupErr
		return result
	}
	result.rows, result.err = this.db.QueryContext(this.ctx, "SELECT * FROM example_table WHERE name=$1", params...)
	return result
}

// Insert returns a struct that will perform the 'insert' query.
// When Execute is called, it will use the following fields:
// [id start_time name]
func (this *Queries_Amazing) Insert(ctx context.Context, db persist.Runnable) *Query_Amazing_Insert {
	return &Query_Amazing_Insert{
		opts: this.opts,
		ctx:  ctx,
		db:   db,
	}
}

// Query_Amazing_Insert (future doc string needed)
type Query_Amazing_Insert struct {
	opts Opts_Amazing
	db   persist.Runnable
	ctx  context.Context
}

func (this *Query_Amazing_Insert) QueryInType_ExampleTable() {}
func (this *Query_Amazing_Insert) QueryOutType_Empty()       {}

// Executes the query 'insert' with parameters retrieved from x.
// Fields used: [id start_time name]
func (this *Query_Amazing_Insert) Execute(x In_Amazing_Insert) *Iter_Amazing_Insert {
	var setupErr error
	params := []interface{}{
		func() (out interface{}) {
			out = x.GetId()
			return
		}(),
		func() (out interface{}) {
			mapper := this.opts.MAPPINGS.TimestampTimestamp()
			out = mapper.ToSql(x.GetStartTime())
			return
		}(),
		func() (out interface{}) {
			out = x.GetName()
			return
		}(),
	}
	result := &Iter_Amazing_Insert{
		tm:  this.opts.MAPPINGS,
		ctx: this.ctx,
	}
	if setupErr != nil {
		result.err = setupErr
		return result
	}
	result.result, result.err = this.db.ExecContext(this.ctx, "INSERT INTO example_table (id, start_time, name) VALUES ($1, $2, $3)", params...)
	return result
}

type Iter_Amazing_SelectById struct {
	result sql.Result
	rows   *sql.Rows
	err    error
	tm     TypeMappings_Amazing
	ctx    context.Context
}

func (this *Iter_Amazing_SelectById) IterOutTypeTestExampleTable() {}
func (this *Iter_Amazing_SelectById) IterInTypeTestPartialTable()  {}

// Each performs 'fun' on each row in the result set.
// Each respects the context passed to it.
// It will stop iteration, and returns this.ctx.Err() if encountered.
func (this *Iter_Amazing_SelectById) Each(fun func(*Row_Amazing_SelectById) error) error {
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
}

// One returns the sole row, or ensures an error if there was not one result when this row is converted
func (this *Iter_Amazing_SelectById) One() *Row_Amazing_SelectById {
	first, hasFirst := this.Next()
	if first != nil && first.err != nil && first.err != io.EOF {
		return &Row_Amazing_SelectById{err: first.err}
	}
	_, hasSecond := this.Next()
	if !hasFirst || hasSecond {
		amount := "none"
		if hasSecond {
			amount = "multiple"
		}
		return &Row_Amazing_SelectById{err: fmt.Errorf("expected exactly 1 result from query 'SelectById' found %s", amount)}
	}
	return first
}

// Zero returns an error if there were any rows in the result
func (this *Iter_Amazing_SelectById) Zero() error {
	row, ok := this.Next()
	if row != nil && row.err != nil && row.err != io.EOF {
		return row.err
	}
	if ok {
		return fmt.Errorf("expected exactly 0 results from query 'SelectById'")
	}
	return nil
}

// Next returns the next scanned row out of the database, or (nil, false) if there are no more rows
func (this *Iter_Amazing_SelectById) Next() (*Row_Amazing_SelectById, bool) {
	if this.err != io.EOF && this.err != nil {
		err := this.err
		this.err = io.EOF
		return &Row_Amazing_SelectById{err: err}, true
	}
	if this.rows == nil {
		this.err = io.EOF
		return nil, false
	}
	cols, err := this.rows.Columns()
	if err != nil {
		return &Row_Amazing_SelectById{err: err}, true
	}
	if !this.rows.Next() {
		if this.err = this.rows.Err(); this.err == nil {
			this.err = io.EOF
			return nil, false
		}
	}
	toScan := make([]interface{}, len(cols))
	scanned := make([]alwaysScanner, len(cols))
	for i := range scanned {
		toScan[i] = &scanned[i]
	}
	if this.err = this.rows.Scan(toScan...); this.err != nil {
		return &Row_Amazing_SelectById{err: this.err}, true
	}
	res := &test.ExampleTable{}
	for i, col := range cols {
		_ = i
		switch col {
		case "id":
			r, ok := (*scanned[i].i).(int64)
			if !ok {
				return &Row_Amazing_SelectById{err: fmt.Errorf("cant convert db column id to protobuf go type ")}, true
			}
			res.Id = r
		case "start_time":
			var converted = this.tm.TimestampTimestamp()
			if err := converted.Scan(*scanned[i].i); err != nil {
				return &Row_Amazing_SelectById{err: fmt.Errorf("could not convert mapped db column start_time to type on test.ExampleTable.StartTime: %v", err)}, true
			}
			if err := converted.ToProto(&res.StartTime); err != nil {
				return &Row_Amazing_SelectById{err: fmt.Errorf("could not convert mapped db column start_timeto type on test.ExampleTable.StartTime: %v", err)}, true
			}
		case "name":
			r, ok := (*scanned[i].i).(string)
			if !ok {
				return &Row_Amazing_SelectById{err: fmt.Errorf("cant convert db column name to protobuf go type ")}, true
			}
			res.Name = r

		default:
			return &Row_Amazing_SelectById{err: fmt.Errorf("unsupported column in output: %s", col)}, true
		}
	}
	return &Row_Amazing_SelectById{item: res}, true
}

// Slice returns all rows found in the iterator as a Slice.
func (this *Iter_Amazing_SelectById) Slice() []*Row_Amazing_SelectById {
	var results []*Row_Amazing_SelectById
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
func (r *Iter_Amazing_SelectById) Columns() ([]string, error) {
	if r.err != nil {
		return nil, r.err
	}
	if r.rows != nil {
		return r.rows.Columns()
	}
	return nil, nil
}

type Iter_Amazing_SelectByName struct {
	result sql.Result
	rows   *sql.Rows
	err    error
	tm     TypeMappings_Amazing
	ctx    context.Context
}

func (this *Iter_Amazing_SelectByName) IterOutTypeTestExampleTable() {}
func (this *Iter_Amazing_SelectByName) IterInTypeTestName()          {}

// Each performs 'fun' on each row in the result set.
// Each respects the context passed to it.
// It will stop iteration, and returns this.ctx.Err() if encountered.
func (this *Iter_Amazing_SelectByName) Each(fun func(*Row_Amazing_SelectByName) error) error {
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
}

// One returns the sole row, or ensures an error if there was not one result when this row is converted
func (this *Iter_Amazing_SelectByName) One() *Row_Amazing_SelectByName {
	first, hasFirst := this.Next()
	if first != nil && first.err != nil && first.err != io.EOF {
		return &Row_Amazing_SelectByName{err: first.err}
	}
	_, hasSecond := this.Next()
	if !hasFirst || hasSecond {
		amount := "none"
		if hasSecond {
			amount = "multiple"
		}
		return &Row_Amazing_SelectByName{err: fmt.Errorf("expected exactly 1 result from query 'SelectByName' found %s", amount)}
	}
	return first
}

// Zero returns an error if there were any rows in the result
func (this *Iter_Amazing_SelectByName) Zero() error {
	row, ok := this.Next()
	if row != nil && row.err != nil && row.err != io.EOF {
		return row.err
	}
	if ok {
		return fmt.Errorf("expected exactly 0 results from query 'SelectByName'")
	}
	return nil
}

// Next returns the next scanned row out of the database, or (nil, false) if there are no more rows
func (this *Iter_Amazing_SelectByName) Next() (*Row_Amazing_SelectByName, bool) {
	if this.err != io.EOF && this.err != nil {
		err := this.err
		this.err = io.EOF
		return &Row_Amazing_SelectByName{err: err}, true
	}
	if this.rows == nil {
		this.err = io.EOF
		return nil, false
	}
	cols, err := this.rows.Columns()
	if err != nil {
		return &Row_Amazing_SelectByName{err: err}, true
	}
	if !this.rows.Next() {
		if this.err = this.rows.Err(); this.err == nil {
			this.err = io.EOF
			return nil, false
		}
	}
	toScan := make([]interface{}, len(cols))
	scanned := make([]alwaysScanner, len(cols))
	for i := range scanned {
		toScan[i] = &scanned[i]
	}
	if this.err = this.rows.Scan(toScan...); this.err != nil {
		return &Row_Amazing_SelectByName{err: this.err}, true
	}
	res := &test.ExampleTable{}
	for i, col := range cols {
		_ = i
		switch col {
		case "id":
			r, ok := (*scanned[i].i).(int64)
			if !ok {
				return &Row_Amazing_SelectByName{err: fmt.Errorf("cant convert db column id to protobuf go type ")}, true
			}
			res.Id = r
		case "start_time":
			var converted = this.tm.TimestampTimestamp()
			if err := converted.Scan(*scanned[i].i); err != nil {
				return &Row_Amazing_SelectByName{err: fmt.Errorf("could not convert mapped db column start_time to type on test.ExampleTable.StartTime: %v", err)}, true
			}
			if err := converted.ToProto(&res.StartTime); err != nil {
				return &Row_Amazing_SelectByName{err: fmt.Errorf("could not convert mapped db column start_timeto type on test.ExampleTable.StartTime: %v", err)}, true
			}
		case "name":
			r, ok := (*scanned[i].i).(string)
			if !ok {
				return &Row_Amazing_SelectByName{err: fmt.Errorf("cant convert db column name to protobuf go type ")}, true
			}
			res.Name = r

		default:
			return &Row_Amazing_SelectByName{err: fmt.Errorf("unsupported column in output: %s", col)}, true
		}
	}
	return &Row_Amazing_SelectByName{item: res}, true
}

// Slice returns all rows found in the iterator as a Slice.
func (this *Iter_Amazing_SelectByName) Slice() []*Row_Amazing_SelectByName {
	var results []*Row_Amazing_SelectByName
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
func (r *Iter_Amazing_SelectByName) Columns() ([]string, error) {
	if r.err != nil {
		return nil, r.err
	}
	if r.rows != nil {
		return r.rows.Columns()
	}
	return nil, nil
}

type Iter_Amazing_Insert struct {
	result sql.Result
	rows   *sql.Rows
	err    error
	tm     TypeMappings_Amazing
	ctx    context.Context
}

func (this *Iter_Amazing_Insert) IterOutTypeEmpty()           {}
func (this *Iter_Amazing_Insert) IterInTypeTestExampleTable() {}

// Each performs 'fun' on each row in the result set.
// Each respects the context passed to it.
// It will stop iteration, and returns this.ctx.Err() if encountered.
func (this *Iter_Amazing_Insert) Each(fun func(*Row_Amazing_Insert) error) error {
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
}

// One returns the sole row, or ensures an error if there was not one result when this row is converted
func (this *Iter_Amazing_Insert) One() *Row_Amazing_Insert {
	first, hasFirst := this.Next()
	if first != nil && first.err != nil && first.err != io.EOF {
		return &Row_Amazing_Insert{err: first.err}
	}
	_, hasSecond := this.Next()
	if !hasFirst || hasSecond {
		amount := "none"
		if hasSecond {
			amount = "multiple"
		}
		return &Row_Amazing_Insert{err: fmt.Errorf("expected exactly 1 result from query 'Insert' found %s", amount)}
	}
	return first
}

// Zero returns an error if there were any rows in the result
func (this *Iter_Amazing_Insert) Zero() error {
	row, ok := this.Next()
	if row != nil && row.err != nil && row.err != io.EOF {
		return row.err
	}
	if ok {
		return fmt.Errorf("expected exactly 0 results from query 'Insert'")
	}
	return nil
}

// Next returns the next scanned row out of the database, or (nil, false) if there are no more rows
func (this *Iter_Amazing_Insert) Next() (*Row_Amazing_Insert, bool) {
	if this.err != io.EOF && this.err != nil {
		err := this.err
		this.err = io.EOF
		return &Row_Amazing_Insert{err: err}, true
	}
	if this.rows == nil {
		this.err = io.EOF
		return nil, false
	}
	cols, err := this.rows.Columns()
	if err != nil {
		return &Row_Amazing_Insert{err: err}, true
	}
	if !this.rows.Next() {
		if this.err = this.rows.Err(); this.err == nil {
			this.err = io.EOF
			return nil, false
		}
	}
	toScan := make([]interface{}, len(cols))
	scanned := make([]alwaysScanner, len(cols))
	for i := range scanned {
		toScan[i] = &scanned[i]
	}
	if this.err = this.rows.Scan(toScan...); this.err != nil {
		return &Row_Amazing_Insert{err: this.err}, true
	}
	res := &Empty{}
	for i, col := range cols {
		_ = i
		switch col {

		default:
			return &Row_Amazing_Insert{err: fmt.Errorf("unsupported column in output: %s", col)}, true
		}
	}
	return &Row_Amazing_Insert{item: res}, true
}

// Slice returns all rows found in the iterator as a Slice.
func (this *Iter_Amazing_Insert) Slice() []*Row_Amazing_Insert {
	var results []*Row_Amazing_Insert
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
func (r *Iter_Amazing_Insert) Columns() ([]string, error) {
	if r.err != nil {
		return nil, r.err
	}
	if r.rows != nil {
		return r.rows.Columns()
	}
	return nil, nil
}

type In_Amazing_SelectById interface {
	GetId() int64
	GetStartTime() *timestamp.Timestamp
}
type Out_Amazing_SelectById interface {
	GetId() int64
	GetStartTime() *timestamp.Timestamp
	GetName() string
}
type Row_Amazing_SelectById struct {
	item Out_Amazing_SelectById
	err  error
}

func newRowAmazingSelectById(item Out_Amazing_SelectById, err error) *Row_Amazing_SelectById {
	return &Row_Amazing_SelectById{item, err}
}

// Unwrap takes an address to a proto.Message as its only parameter
// Unwrap can only set into output protos of that match method return types + the out option on the query itself
func (this *Row_Amazing_SelectById) Unwrap(pointerToMsg proto.Message) error {
	if this.err != nil {
		return this.err
	}
	if o, ok := (pointerToMsg).(*test.ExampleTable); ok {
		if o == nil {
			return fmt.Errorf("must initialize *test.ExampleTable before giving to Unwrap()")
		}
		res, _ := this.TestExampleTable()
		_ = res
		o.Id = res.Id
		o.StartTime = res.StartTime
		o.Name = res.Name
		return nil
	}

	if o, ok := (pointerToMsg).(*test.ExampleTable); ok {
		if o == nil {
			return fmt.Errorf("must initialize *test.ExampleTable before giving to Unwrap()")
		}
		res, _ := this.TestExampleTable()
		_ = res
		o.Id = res.Id
		o.StartTime = res.StartTime
		o.Name = res.Name
		return nil
	}
	if o, ok := (pointerToMsg).(*test.ExampleTable); ok {
		if o == nil {
			return fmt.Errorf("must initialize *test.ExampleTable before giving to Unwrap()")
		}
		res, _ := this.TestExampleTable()
		_ = res
		o.Id = res.Id
		o.StartTime = res.StartTime
		o.Name = res.Name
		return nil
	}

	return nil
}
func (this *Row_Amazing_SelectById) TestExampleTable() (*test.ExampleTable, error) {
	if this.err != nil {
		return nil, this.err
	}
	return &test.ExampleTable{
		Id:        this.item.GetId(),
		StartTime: this.item.GetStartTime(),
		Name:      this.item.GetName(),
	}, nil
}

func (this *Row_Amazing_SelectById) Proto() (*test.ExampleTable, error) {
	if this.err != nil {
		return nil, this.err
	}
	return &test.ExampleTable{
		Id:        this.item.GetId(),
		StartTime: this.item.GetStartTime(),
		Name:      this.item.GetName(),
	}, nil
}

type In_Amazing_SelectByName interface {
	GetName() string
}
type Out_Amazing_SelectByName interface {
	GetId() int64
	GetStartTime() *timestamp.Timestamp
	GetName() string
}
type Row_Amazing_SelectByName struct {
	item Out_Amazing_SelectByName
	err  error
}

func newRowAmazingSelectByName(item Out_Amazing_SelectByName, err error) *Row_Amazing_SelectByName {
	return &Row_Amazing_SelectByName{item, err}
}

// Unwrap takes an address to a proto.Message as its only parameter
// Unwrap can only set into output protos of that match method return types + the out option on the query itself
func (this *Row_Amazing_SelectByName) Unwrap(pointerToMsg proto.Message) error {
	if this.err != nil {
		return this.err
	}
	if o, ok := (pointerToMsg).(*test.ExampleTable); ok {
		if o == nil {
			return fmt.Errorf("must initialize *test.ExampleTable before giving to Unwrap()")
		}
		res, _ := this.TestExampleTable()
		_ = res
		o.Id = res.Id
		o.StartTime = res.StartTime
		o.Name = res.Name
		return nil
	}

	if o, ok := (pointerToMsg).(*test.ExampleTable); ok {
		if o == nil {
			return fmt.Errorf("must initialize *test.ExampleTable before giving to Unwrap()")
		}
		res, _ := this.TestExampleTable()
		_ = res
		o.Id = res.Id
		o.StartTime = res.StartTime
		o.Name = res.Name
		return nil
	}
	if o, ok := (pointerToMsg).(*test.ExampleTable); ok {
		if o == nil {
			return fmt.Errorf("must initialize *test.ExampleTable before giving to Unwrap()")
		}
		res, _ := this.TestExampleTable()
		_ = res
		o.Id = res.Id
		o.StartTime = res.StartTime
		o.Name = res.Name
		return nil
	}

	return nil
}
func (this *Row_Amazing_SelectByName) TestExampleTable() (*test.ExampleTable, error) {
	if this.err != nil {
		return nil, this.err
	}
	return &test.ExampleTable{
		Id:        this.item.GetId(),
		StartTime: this.item.GetStartTime(),
		Name:      this.item.GetName(),
	}, nil
}

func (this *Row_Amazing_SelectByName) Proto() (*test.ExampleTable, error) {
	if this.err != nil {
		return nil, this.err
	}
	return &test.ExampleTable{
		Id:        this.item.GetId(),
		StartTime: this.item.GetStartTime(),
		Name:      this.item.GetName(),
	}, nil
}

type In_Amazing_Insert interface {
	GetId() int64
	GetStartTime() *timestamp.Timestamp
	GetName() string
}
type Out_Amazing_Insert interface {
}
type Row_Amazing_Insert struct {
	item Out_Amazing_Insert
	err  error
}

func newRowAmazingInsert(item Out_Amazing_Insert, err error) *Row_Amazing_Insert {
	return &Row_Amazing_Insert{item, err}
}

// Unwrap takes an address to a proto.Message as its only parameter
// Unwrap can only set into output protos of that match method return types + the out option on the query itself
func (this *Row_Amazing_Insert) Unwrap(pointerToMsg proto.Message) error {
	if this.err != nil {
		return this.err
	}
	if o, ok := (pointerToMsg).(*Empty); ok {
		if o == nil {
			return fmt.Errorf("must initialize *Empty before giving to Unwrap()")
		}
		res, _ := this.Empty()
		_ = res

		return nil
	}

	if o, ok := (pointerToMsg).(*test.NumRows); ok {
		if o == nil {
			return fmt.Errorf("must initialize *test.NumRows before giving to Unwrap()")
		}
		res, _ := this.TestNumRows()
		_ = res

		return nil
	}
	if o, ok := (pointerToMsg).(*test.Ids); ok {
		if o == nil {
			return fmt.Errorf("must initialize *test.Ids before giving to Unwrap()")
		}
		res, _ := this.TestIds()
		_ = res

		return nil
	}

	return nil
}
func (this *Row_Amazing_Insert) Empty() (*Empty, error) {
	if this.err != nil {
		return nil, this.err
	}
	return &Empty{}, nil
}
func (this *Row_Amazing_Insert) TestNumRows() (*test.NumRows, error) {
	if this.err != nil {
		return nil, this.err
	}
	return &test.NumRows{}, nil
}
func (this *Row_Amazing_Insert) TestIds() (*test.Ids, error) {
	if this.err != nil {
		return nil, this.err
	}
	return &test.Ids{}, nil
}

func (this *Row_Amazing_Insert) Proto() (*Empty, error) {
	if this.err != nil {
		return nil, this.err
	}
	return &Empty{}, nil
}

type Hooks_Amazing interface {
	UniarySelectWithHooksBeforeHook(context.Context, *test.PartialTable) (*test.ExampleTable, error)
	ServerStreamWithHooksBeforeHook(context.Context, *test.Name) (*test.ExampleTable, error)
	ClientStreamWithHookBeforeHook(context.Context, *test.ExampleTable) (*test.Ids, error)
	UniarySelectWithHooksAfterHook(context.Context, *test.PartialTable, *test.ExampleTable) error
	ServerStreamWithHooksAfterHook(context.Context, *test.Name, *test.ExampleTable) error
	ClientStreamWithHookAfterHook(context.Context, *test.ExampleTable, *test.Ids) error
}
type DefaultHooks_Amazing struct{}

func (*DefaultHooks_Amazing) UniarySelectWithHooksBeforeHook(context.Context, *test.PartialTable) (*test.ExampleTable, error) {
	return nil, nil
}
func (*DefaultHooks_Amazing) ServerStreamWithHooksBeforeHook(context.Context, *test.Name) (*test.ExampleTable, error) {
	return nil, nil
}
func (*DefaultHooks_Amazing) ClientStreamWithHookBeforeHook(context.Context, *test.ExampleTable) (*test.Ids, error) {
	return nil, nil
}
func (*DefaultHooks_Amazing) UniarySelectWithHooksAfterHook(context.Context, *test.PartialTable, *test.ExampleTable) error {
	return nil
}
func (*DefaultHooks_Amazing) ServerStreamWithHooksAfterHook(context.Context, *test.Name, *test.ExampleTable) error {
	return nil
}
func (*DefaultHooks_Amazing) ClientStreamWithHookAfterHook(context.Context, *test.ExampleTable, *test.Ids) error {
	return nil
}

type TypeMappings_Amazing interface {
	TimestampTimestamp() MappingImpl_Amazing_TimestampTimestamp
}
type DefaultTypeMappings_Amazing struct{}

func (this *DefaultTypeMappings_Amazing) TimestampTimestamp() MappingImpl_Amazing_TimestampTimestamp {
	return &DefaultMappingImpl_Amazing_TimestampTimestamp{}
}

type DefaultMappingImpl_Amazing_TimestampTimestamp struct{}

func (this *DefaultMappingImpl_Amazing_TimestampTimestamp) ToProto(**timestamp.Timestamp) error {
	return nil
}
func (this *DefaultMappingImpl_Amazing_TimestampTimestamp) ToSql(*timestamp.Timestamp) sql.Scanner {
	return this
}
func (this *DefaultMappingImpl_Amazing_TimestampTimestamp) Scan(interface{}) error {
	return nil
}
func (this *DefaultMappingImpl_Amazing_TimestampTimestamp) Value() (driver.Value, error) {
	return "DEFAULT_TYPE_MAPPING_VALUE", nil
}

type MappingImpl_Amazing_TimestampTimestamp interface {
	ToProto(**timestamp.Timestamp) error
	ToSql(*timestamp.Timestamp) sql.Scanner
	sql.Scanner
	driver.Valuer
}

type Opts_Amazing struct {
	MAPPINGS TypeMappings_Amazing
	HOOKS    Hooks_Amazing
}

func OptsAmazing(hooks Hooks_Amazing, mappings TypeMappings_Amazing) Opts_Amazing {
	opts := Opts_Amazing{
		HOOKS:    &DefaultHooks_Amazing{},
		MAPPINGS: &DefaultTypeMappings_Amazing{},
	}
	if hooks != nil {
		opts.HOOKS = hooks
	}
	if mappings != nil {
		opts.MAPPINGS = mappings
	}
	return opts
}

type Impl_Amazing struct {
	opts     *Opts_Amazing
	QUERIES  *Queries_Amazing
	HANDLERS RestOfHandlers_Amazing
	DB       *sql.DB
}

func ImplAmazing(db *sql.DB, handlers RestOfHandlers_Amazing, opts ...Opts_Amazing) *Impl_Amazing {
	var myOpts Opts_Amazing
	if len(opts) > 0 {
		myOpts = opts[0]
	} else {
		myOpts = OptsAmazing(&DefaultHooks_Amazing{}, &DefaultTypeMappings_Amazing{})
	}
	return &Impl_Amazing{
		opts:     &myOpts,
		QUERIES:  QueriesAmazing(myOpts),
		DB:       db,
		HANDLERS: handlers,
	}
}

type RestOfHandlers_Amazing interface {
	UnImplementedPersistMethod(context.Context, *test.ExampleTable) (*test.ExampleTable, error)
	NoGenerationForBadReturnTypes(context.Context, *test.ExampleTable) (*BadReturn, error)
}

func (this *Impl_Amazing) UnImplementedPersistMethod(ctx context.Context, req *test.ExampleTable) (*test.ExampleTable, error) {
	return this.HANDLERS.UnImplementedPersistMethod(ctx, req)
}

func (this *Impl_Amazing) NoGenerationForBadReturnTypes(ctx context.Context, req *test.ExampleTable) (*BadReturn, error) {
	return this.HANDLERS.NoGenerationForBadReturnTypes(ctx, req)
}

func (this *Impl_Amazing) UniarySelect(ctx context.Context, req *test.PartialTable) (*test.ExampleTable, error) {
	query := this.QUERIES.SelectById(ctx, this.DB)

	result := query.Execute(req)
	res, err := result.One().TestExampleTable()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (this *Impl_Amazing) UniarySelectWithHooks(ctx context.Context, req *test.PartialTable) (*test.ExampleTable, error) {
	query := this.QUERIES.SelectById(ctx, this.DB)

	{
		beforeRes, err := this.opts.HOOKS.UniarySelectWithHooksBeforeHook(ctx, req)
		if err != nil {
			return nil, gstatus.Errorf(codes.Unknown, "error in before hook: %v", err)
		} else if beforeRes != nil {
			return beforeRes, nil
		}
	}

	result := query.Execute(req)
	res, err := result.One().TestExampleTable()
	if err != nil {
		return nil, err
	}

	{
		if err := this.opts.HOOKS.UniarySelectWithHooksAfterHook(ctx, req, res); err != nil {
			return nil, gstatus.Errorf(codes.Unknown, "error in after hook: %v", err)
		}
	}

	return res, nil
}

func (this *Impl_Amazing) ServerStream(req *test.Name, stream Amazing_ServerStreamServer) error {
	tx, err := DefaultServerStreamingPersistTx(stream.Context(), this.DB)
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error creating persist tx: %v", err)
	}
	if err := this.ServerStreamTx(req, stream, tx); err != nil {
		return gstatus.Errorf(codes.Unknown, "error executing 'select_by_name' query: %v", err)
	}
	return nil
}
func (this *Impl_Amazing) ServerStreamTx(req *test.Name, stream Amazing_ServerStreamServer, tx persist.PersistTx) error {
	ctx := stream.Context()
	query := this.QUERIES.SelectByName(ctx, tx)
	iter := query.Execute(req)
	return iter.Each(func(row *Row_Amazing_SelectByName) error {
		res, err := row.TestExampleTable()
		if err != nil {
			return err
		}
		return stream.Send(res)
	})
}

func (this *Impl_Amazing) ServerStreamWithHooks(req *test.Name, stream Amazing_ServerStreamWithHooksServer) error {
	tx, err := DefaultServerStreamingPersistTx(stream.Context(), this.DB)
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error creating persist tx: %v", err)
	}
	if err := this.ServerStreamWithHooksTx(req, stream, tx); err != nil {
		return gstatus.Errorf(codes.Unknown, "error executing 'select_by_name' query: %v", err)
	}
	return nil
}
func (this *Impl_Amazing) ServerStreamWithHooksTx(req *test.Name, stream Amazing_ServerStreamWithHooksServer, tx persist.PersistTx) error {
	ctx := stream.Context()
	query := this.QUERIES.SelectByName(ctx, tx)
	iter := query.Execute(req)
	return iter.Each(func(row *Row_Amazing_SelectByName) error {
		res, err := row.TestExampleTable()
		if err != nil {
			return err
		}
		return stream.Send(res)
	})
}

func (this *Impl_Amazing) ClientStream(stream Amazing_ClientStreamServer) error {
	tx, err := DefaultClientStreamingPersistTx(stream.Context(), this.DB)
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error creating persist tx: %v", err)
	}
	if err := this.ClientStreamTx(stream, tx); err != nil {
		return gstatus.Errorf(codes.Unknown, "error executing 'insert' query: %v", err)
	}
	return nil
}
func (this *Impl_Amazing) ClientStreamTx(stream Amazing_ClientStreamServer, tx persist.PersistTx) error {
	query := this.QUERIES.Insert(stream.Context(), tx)
	var first *test.ExampleTable
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

		result := query.Execute(req)
		if err := result.Zero(); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("error executing 'insert' query :::AND COULD NOT ROLLBACK::: rollback err: %v, query err: %v", rollbackErr, err)
		}
	}
	res := &test.NumRows{}

	if err := stream.SendAndClose(res); err != nil {
		return gstatus.Errorf(codes.Unknown, "error sending back response: %v", err)
	}
	return nil
}

func (this *Impl_Amazing) ClientStreamWithHook(stream Amazing_ClientStreamWithHookServer) error {
	tx, err := DefaultClientStreamingPersistTx(stream.Context(), this.DB)
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error creating persist tx: %v", err)
	}
	if err := this.ClientStreamWithHookTx(stream, tx); err != nil {
		return gstatus.Errorf(codes.Unknown, "error executing 'insert' query: %v", err)
	}
	return nil
}
func (this *Impl_Amazing) ClientStreamWithHookTx(stream Amazing_ClientStreamWithHookServer, tx persist.PersistTx) error {
	query := this.QUERIES.Insert(stream.Context(), tx)
	var first *test.ExampleTable
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

		{
			beforeRes, err := this.opts.HOOKS.ClientStreamWithHookBeforeHook(stream.Context(), req)
			if err != nil {
				return gstatus.Errorf(codes.Unknown, "error in before hook: %v", err)
			} else if beforeRes != nil {
				continue
			}
		}

		result := query.Execute(req)
		if err := result.Zero(); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("error executing 'insert' query :::AND COULD NOT ROLLBACK::: rollback err: %v, query err: %v", rollbackErr, err)
		}
	}
	res := &test.Ids{}

	{
		if err := this.opts.HOOKS.ClientStreamWithHookAfterHook(stream.Context(), first, res); err != nil {
			return gstatus.Errorf(codes.Unknown, "error in after hook: %v", err)
		}
	}

	if err := stream.SendAndClose(res); err != nil {
		return gstatus.Errorf(codes.Unknown, "error sending back response: %v", err)
	}
	return nil
}
