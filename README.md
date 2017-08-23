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

## TODO

### More supports

- postgres

- mssql

- sqlite

### Execute duration

To print how long take quering.