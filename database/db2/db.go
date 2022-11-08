package db2

import (
	"context"
	"errors"
	"log"
	"sync"

	_ "github.com/ibmdb/go_ibm_db"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var once sync.Once

func Init(dsn string) error {
	var initError error

	once.Do(func() {
		db, err := sqlx.Connect("go_ibm_db", dsn)
		if err != nil {
			initError = err
			return
		}

		DB = db
	})

	return initError
}

var Health = func() error {
	if err := DB.Ping(); err != nil {
		log.Println(err)
		return errors.New("DB2 is not available")
	}

	return nil
}

var Exec = func(ctx context.Context, query string, args ...interface{}) (affected int64, err error) {
	rs, err := DB.ExecContext(ctx, query, args)
	if err != nil {
		return
	}
	return rs.RowsAffected()
}

var NamedExec = func(ctx context.Context, query string, args interface{}) (affected int64, err error) {
	rs, err := DB.NamedExecContext(ctx, query, args)
	if err != nil {
		return
	}
	return rs.RowsAffected()
}

var Get = func(ctx context.Context, dest interface{}, query string, args interface{}) error {
	return DB.GetContext(ctx, dest, query, args)
}

var Select = func(ctx context.Context, dest interface{}, query string, args interface{}) error {
	return DB.SelectContext(ctx, dest, query, args)
}
