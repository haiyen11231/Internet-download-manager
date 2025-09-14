package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/haiyen11231/Internet-download-manager/internal/configs"
)

type Database interface {
	Delete(table interface{}) *goqu.DeleteDataset
	Dialect() string
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	From(from ...interface{}) *goqu.SelectDataset
	Insert(table interface{}) *goqu.InsertDataset
	Logger(logger goqu.Logger)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ScanStruct(i interface{}, query string, args ...interface{}) (bool, error)
	ScanStructContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
	ScanStructs(i interface{}, query string, args ...interface{}) error
	ScanStructsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
	ScanVal(i interface{}, query string, args ...interface{}) (bool, error)
	ScanValContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
	ScanVals(i interface{}, query string, args ...interface{}) error
	ScanValsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
	Select(cols ...interface{}) *goqu.SelectDataset
	Trace(op string, sqlString string, args ...interface{})
	Truncate(table ...interface{}) *goqu.TruncateDataset
	Update(table interface{}) *goqu.UpdateDataset
}

// return sql.db de cam tu ben ngoai va tuong tac db
func InitializeDB(DBConfig configs.DBConfig) (db *sql.DB, cleanup func(), err error) {
	// Construct connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		DBConfig.Username, 
		DBConfig.Password, 
		DBConfig.Host, 
		DBConfig.Port, 
		DBConfig.DBName,
	)
	// Open database connection
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Printf("Error connecting to the database: ", err)
		return nil, nil, err
	}

	// Close the database connection when the function exits
	// defer db.Close()	
	cleanup = func() {
		db.Close()
	}
	log.Println("Successfully connected to the database")
	return db, cleanup, nil
}

func InitializeGoquDB(db *sql.DB) *goqu.Database{
	return goqu.New("mysql", db)
}