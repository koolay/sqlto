package db

import (
	"errors"
	"strings"

	"github.com/koolay/sqlto/config"
)

// ExecSQL execute sql
func ExecSQL(driver string, ctx *config.Context) (items []map[string]interface{}, err error) {
	if strings.TrimSpace(ctx.SQL) == "" {
		return items, errors.New("sql is empty")
	}
	if driver == "mysql" {
		return mysqlQuery(ctx)
	}
	return items, errors.New("not supported driver")
}
