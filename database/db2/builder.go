package db2

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	DBLikeTemplate string = "%%%v%%"

	DBLikeClause  string = "LIKE ?"
	DBEqualClause string = "= ?"
)

func CreateWhereStatements(entries ...WhereStatementEntry) []WhereStatementEntry {
	return entries
}

func NewWhereStatementEntry(column string, value interface{}, clause string, template *string) WhereStatementEntry {
	return WhereStatementEntry{
		Column:   column,
		Value:    value,
		Clause:   clause,
		Template: template,
	}
}

type WhereStatementEntry struct {
	Column   string
	Template *string
	Clause   string
	Value    interface{}
}

func (ws *WhereStatementEntry) GetClouse() string {
	return fmt.Sprintf("%s %s", ws.Column, ws.Clause)
}

func (ws *WhereStatementEntry) GetValue() interface{} {
	if ws.Template == nil {
		return ws.Value
	}

	return fmt.Sprintf(*ws.Template, reflect.Indirect(reflect.ValueOf(ws.Value)))
}

var BuildWhereCondition = func(params ...WhereStatementEntry) (whereQuery string, args []interface{}) {
	if len(params) == 0 {
		return
	}

	var whereClause []string
	for i := range params {
		if !reflect.ValueOf(params[i].Value).IsNil() {
			whereClause = append(whereClause, params[i].GetClouse())
			args = append(args, params[i].GetValue())
		}
	}

	if len(whereClause) == 0 {
		return
	}

	return fmt.Sprintf(" WHERE %s", strings.Join(whereClause, " AND ")), args
}
