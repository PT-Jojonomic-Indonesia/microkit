package db2

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reflect"
	"strings"
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
	rs, err := DB.ExecContext(ctx, query, args...)
	if err != nil {
		return
	}
	return rs.RowsAffected()
}

var NamedExecWithTx = func(ctx context.Context, tx *sqlx.Tx, query string, args interface{}) (affected int64, err error) {
	rs, err := tx.NamedExecContext(ctx, query, args)
	if err != nil {
		tx.Rollback()
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

var Get = func(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return DB.GetContext(ctx, dest, query, args...)
}

var Select = func(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return DB.SelectContext(ctx, dest, query, args...)
}

var GetList = func(ctx context.Context, query string, conditions []WhereStatementEntry, paginate *Paginate, dest interface{}, selectColumn ...string) (err error) {
	if reflect.TypeOf(dest).Kind() != reflect.Pointer {
		return errors.New("destination is not pointer")
	}

	whereQuery, args := BuildWhereCondition(conditions...)
	if whereQuery != "" {
		query = query + whereQuery
	}

	if paginate != nil {
		query += "OFFSET ? LIMIT ?"
		args = append(args, paginate.Offset, paginate.Limit)
	}

	err = Select(ctx, dest, query, args...)
	return
}

var GetDetail = func(ctx context.Context, query string, conditions []WhereStatementEntry, dest interface{}, selectColumn ...string) (err error) {
	if reflect.TypeOf(dest).Kind() != reflect.Pointer {
		return errors.New("destination is not pointer")
	}

	whereQuery, args := BuildWhereCondition(conditions...)
	if whereQuery != "" {
		query = query + whereQuery
	}

	err = Get(ctx, dest, query, args...)
	return
}

var GetSelectedColumnt = func(selectedColumn []string) string {
	if len(selectedColumn) == 0 {
		return "*"
	}

	return strings.Join(selectedColumn, ",")
}

var Update = func(ctx context.Context, query string, args interface{}) error {
	affected, err := NamedExec(ctx, query, args)
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return err
}

var Delete = func(ctx context.Context, query string, conditions []WhereStatementEntry) (err error) {
	whereQuery, args := BuildWhereCondition(conditions...)
	if whereQuery != "" {
		query = query + whereQuery
	}

	affected, err := Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return
}
