package migrator

import (
	"context"
	"database/sql"
	"time"
)

// InvokerFunc is the function type of the actual invoker. It should be called in an interceptor.
type InvokerFunc = func(ctx context.Context, sql string) error

// InterceptorFunc is the function type of an interceptor. An interceptor should implement this function to fulfill it's purpose.
type InterceptorFunc = func(ctx context.Context, sql string, invoker InvokerFunc) error

type database struct {
	db               *sql.DB
	tx               *sql.Tx
	logger           func(sql string, durationNano int64)
	retryPolicy      func(error) bool
	enableCallerInfo bool
	interceptor      InterceptorFunc
}

type cursor struct {
	rows *sql.Rows
}

type Cursor interface {
	Next() bool
	Scan(dest ...interface{}) error
	GetMap() (map[string]interface{}, error)
	Close() error
}

// Database is the interface of a database with underlying sql.DB object.
type Database interface {
	// GetDB Get the underlying sql.DB object of the database
	GetDB() *sql.DB

	// Query Executes a query and return the cursor
	Query(sql string) (Cursor, error)
	// QueryContext Executes a query with context and return the cursor
	QueryContext(ctx context.Context, sqlString string) (Cursor, error)
	// Execute Executes a statement
	Execute(sql string) (sql.Result, error)
	// ExecuteContext Executes a statement with context
	ExecuteContext(ctx context.Context, sql string) (sql.Result, error)
	// SetLogger Set the logger function
	SetLogger(logger func(sql string, durationNano int64))
	// SetRetryPolicy Set the retry policy function.
	// The retry policy function returns true if needs retry.
	SetRetryPolicy(retryPolicy func(err error) bool)
	// EnableCallerInfo enable or disable caller info
	EnableCallerInfo(enableCallerInfo bool)
	// SetInterceptor Set a interceptor function
	SetInterceptor(interceptor InterceptorFunc)

	// Initiate a SELECT statement
	//Select(fields ...interface{}) selectWithFields
	// Initiate a SELECT DISTINCT statement
	//SelectDistinct(fields ...interface{}) selectWithFields
	// Initiate a SELECT * FROM statement
	//SelectFrom(tables ...Table) selectWithTables
	// Initiate a INSERT INTO statement
	//InsertInto(table Table) insertWithTable
	// Initiate a REPLACE INTO statement
	//ReplaceInto(table Table) insertWithTable
	// Initiate a UPDATE statement
	//Update(table Table) updateWithSet
	// Initiate a DELETE FROM statement
	//DeleteFrom(table Table) deleteWithTable
}

type txOrDB interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func (d *database) SetLogger(logger func(sql string, durationNano int64)) {
	d.logger = logger
}

func (d *database) SetRetryPolicy(retryPolicy func(err error) bool) {
	d.retryPolicy = retryPolicy
}

func (d *database) EnableCallerInfo(enableCallerInfo bool) {
	d.enableCallerInfo = enableCallerInfo
}

func (d *database) SetInterceptor(interceptor InterceptorFunc) {
	d.interceptor = interceptor
}

// Open a database, similar to sql.Open
func Open(driver, databaseName string) (db *database, err error) {
	var sqlDB *sql.DB
	if databaseName != "" {
		sqlDB, err = sql.Open(driver, databaseName)
		if err != nil {
			return
		}
	}
	db = &database{
		db: sqlDB,
	}
	return
}

func (d database) GetDB() *sql.DB {
	return d.db
}

func (d database) getTxOrDB() txOrDB {
	if d.tx != nil {
		return d.tx
	}
	return d.db
}

func (d database) Query(sqlString string) (*cursor, error) {
	return d.QueryContext(context.Background(), sqlString)
}

func (d database) QueryContext(ctx context.Context, sqlString string) (*cursor, error) {
	isRetry := false
	for {
		sqlStringWithCallerInfo := getCallerInfo(d, isRetry) + sqlString

		rows, err := d.queryContextOnce(ctx, sqlStringWithCallerInfo)
		if err != nil {
			isRetry = d.tx == nil && d.retryPolicy != nil && d.retryPolicy(err)
			if isRetry {
				continue
			}
			return nil, err
		}
		return &cursor{rows: rows}, nil
	}
}

func (d database) queryContextOnce(ctx context.Context, sqlStringWithCallerInfo string) (*sql.Rows, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	startTime := time.Now().UnixNano()
	defer func() {
		endTime := time.Now().UnixNano()
		if d.logger != nil {
			d.logger(sqlStringWithCallerInfo, endTime-startTime)
		}
	}()

	interceptor := d.interceptor
	var rows *sql.Rows
	invoker := func(ctx context.Context, sql string) (err error) {
		rows, err = d.getTxOrDB().QueryContext(ctx, sql)
		return
	}

	var err error
	if interceptor == nil {
		err = invoker(ctx, sqlStringWithCallerInfo)
	} else {
		err = interceptor(ctx, sqlStringWithCallerInfo, invoker)
	}
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (d database) Execute(sql string) (sql.Result, error) {
	return d.ExecuteContext(context.Background(), sql)
}

func (d database) ExecuteContext(ctx context.Context, sql string) (sql.Result, error) {
	sql = getCallerInfo(d, false) + sql
	startTime := time.Now().UnixNano()
	result, err := d.getTxOrDB().ExecContext(ctx, sql)
	endTime := time.Now().UnixNano()
	if d.logger != nil {
		d.logger(sql, endTime-startTime)
	}
	return result, err
}

func getCallerInfo(db database, f bool) string {
	return ""

}
