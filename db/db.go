package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koolay/sqlto/config"
	_ "github.com/lib/pq"
)

// IDB db accessor
type IDB interface {
	Query(ctx *config.Context) (items []map[string]interface{}, err error)
}

// DB inheirt IDB
type DB struct {
	driver string
}

// NewDB instance a db
func NewDB(driver string) IDB {
	return &DB{driver: driver}
}

// Query execute sql
func (p *DB) Query(ctx *config.Context) (items []map[string]interface{}, err error) {
	fmt.Printf("connecting %s ... \n", p.driver)
	sess, err := sql.Open(p.driver, ctx.Source)
	if err != nil {
		return items, err
	}
	defer sess.Close()
	rows, err := sess.Query(ctx.SQL)
	if err != nil {
		return items, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return items, err
	}
	cols := make([]interface{}, len(columns))
	colsPtrs := make([]interface{}, len(columns))
	for i := 0; i < len(columns); i++ {
		colsPtrs[i] = &cols[i]
	}
	fmt.Println("executing quering ...")
	for rows.Next() {
		err = rows.Scan(colsPtrs...)
		if err != nil {
			log.Fatal(err)
		}
		var rowMap = make(map[string]interface{})
		for i, colName := range columns {
			val := cols[i]
			valBytes, ok := val.([]byte)
			if ok {
				rowMap[colName] = string(valBytes)
			} else {
				rowMap[colName] = ""
			}
		}
		items = append(items, rowMap)
	}
	fmt.Println("executed quering")
	return
}
