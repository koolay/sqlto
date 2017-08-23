sqlto
--------

Export data to xls, json, csv files from a SQL.

## How to run

- To xls

`sqlto mysql -s "root:password@tcp(localhost:3306)/dbname?charset=utf8mb4,utf8" -l "select * from user limit 3"  -p -o user.xls`

- To stdout

**from mysql**

`sqlto mysql -s "root:password@tcp(localhost:3306)/dbname?charset=utf8mb4,utf8" -l "select * from user limit 3"`

**from postgres**

`sqlto postgres -s "postgres://root:dev@localhost:5433/sample?sslmode=disable" -l "select * from user "  -p`

- To json file

`sqlto mysql -s "root:password@tcp(localhost:3306)/dbname?charset=utf8mb4,utf8" -l "select * from user limit 3"  -p -o user.json`

- More help

`sqlto -h`

```

Support mysql, pg, sqlite, mssql.

Usage:
  sqlto [flags]
  sqlto [command]

Available Commands:
  help        Help about any command
  mysql       Export records from mysql
  postgres    Export records from postgres
  version     Print version

Flags:
  -h, --help            help for sqlto
  -o, --output string   Output file eg. employees.xls
  -p, --pretty          Pretty json output. It is only affective for json.
  -l, --sql string      SQL. eg. select * from employee limit 20

Use "sqlto [command] --help" for more information about a command.

```

## TODO

### More supports

- postgres

- mssql

- sqlite

### Execute duration

To print how long take quering.