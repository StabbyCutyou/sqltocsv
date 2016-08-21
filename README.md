SQL To CSV
============

This is a utility to make use of golangs sql package and it's streaming capabilities,
for exporting large result sets out of an RDBMS and into a text file.

Currently, it supports MySQL and Postgres drivers, which means any database that
supports those drivers should work. For example, RedShift works with the Postgres driver,
and MemSQL works with the MySQL driver.

If there are unsupported drivers, or issues with "compatible" drivers, please
file an issue, or better yet open a PR to address the problem.

Flags
=====

SqlToCSV supports a number of flags to direct how it works

* d : The (d)atabase adapter to use
* c : The (c)onnection string to use
* q : The (q)uery to use
* m : The deli(m)eter to use: 'comma' or 'tab'. Defaults to 'comma'
* o : The indices of the fields to (o)bfuscate. Ex: "1,3,4"
* w : The indices of the fields to (w)rap in quotes. Ex: "2,6,8"
* t : The (t)ype of quote to use with -w: 'single' or 'double'. Defaults to 'double'

Obfuscation
===========
One key feature of SqlToCSV is the ability to obfuscate data if needed. You might need
to anonymize certain fields when exporting, for the purpose of giving the data to a
third party. The `-o` flag will let you specify which fields get this treatment.

When obfuscating, SqlToCSV will track the original value, and replace it with an
auto-incrementing number (future processes for obfuscation to better match the original
field type are coming). Each occurance of the same underlying value will get the same
obfuscated id

Example
========
To run it, call from bash like so:

```bash
sqltocsv -d "mysql" -c "root@tcp(localhost:3306)/mydb" -q "SELECT * FROM table" > outfile.csv
```

TODO
=====
* Obfuscation should be more fleshed out - generate values that match the original type, option to not use the same value, etc.

LICENSE
=========
Apache v2 - See LICENSE
