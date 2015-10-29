SQL To CSV
============

This is a utility to make use of golangs sql package and it's streaming capabilities,
for exporting large result sets out of an RDBMS and into a text file.

Currently, it supports TSV and mysql. Additional support coming for setting the delimeter,
as well as additional databases.

To run it, call from bash like so:

```bash
STC_DBADAPTER="mysql" STC_CONNSTRING="root@tcp(localhost:3306)/mydb" STC_QUERY="SELECT * FROM table" STC_OUTPUTFILE="data.csv" go run sqltocsv.go
```

TODO
=====
Proper Readme
Postgres support
Variable delimiter

LICENSE
=========
Apache v2 - See LICENSE
