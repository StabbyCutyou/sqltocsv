package main

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
)

func getMySQLConfig() *config {
	return &config{
		dbAdapter:       "mysql",
		connString:      "root@/sqltocsvtest",
		sqlQuery:        "SELECT * FROM sample",
		obfuscateFields: "",
		delimiter:       "comma",
		quoteFields:     "",
		quoteType:       "",
	}
}

func getPostgreSQLConfig() *config {
	return &config{
		dbAdapter:       "postgres",
		connString:      "dbname=sqltocsvtest sslmode=disable",
		sqlQuery:        "SELECT * FROM sample",
		obfuscateFields: "",
		delimiter:       "comma",
		quoteFields:     "",
		quoteType:       "",
	}
}

func setupMySQL(db *sqlx.DB) {
	// Create the table
	_, err := db.Exec(`CREATE TABLE sample(
    VARCHAR50 varchar(50),
    CHAR2 char(2),
    NUM int,
    DECIMAL82 decimal(8,2),
    TEXT text,
    BOOL bit
  )`)
	// Insert the data
	if err != nil {
		log.Fatal(err)
	}
	// Insert the data
	_, err = db.Exec(`INSERT INTO sample(VARCHAR50,CHAR2,NUM,DECIMAL82,TEXT,BOOL) VALUES ('Words','ZZ',100,12.89,'HI, EVERYBODY!',1)`)

	if err != nil {
		log.Fatal(err)
	}
}

func tearDownMySQL(db *sqlx.DB) {
	// Drop the table
	db.Exec("DROP TABLE sample")
}

func setupPostgreSQL(db *sqlx.DB) {
	// Create the table
	_, err := db.Exec(`CREATE TABLE sample(
    VARCHAR50 varchar(50),
    CHAR2 char(2),
    NUM int,
    DECIMAL82 decimal(8,2),
    TEXT text,
    BOOL bit
  )`)

	if err != nil {
		log.Fatal(err)
	}
	// Insert the data
	_, err = db.Exec(`INSERT INTO sample(VARCHAR50,CHAR2,NUM,DECIMAL82,TEXT,BOOL) VALUES ('Words','ZZ',100,12.89,'HI, EVERYBODY!',B'1')`)

	if err != nil {
		log.Fatal(err)
	}
}

func tearDownPostgreSQL(db *sqlx.DB) {
	// Drop the table
	db.Exec("DROP TABLE sample")
}

func TestMySQL(t *testing.T) {
	cfg := getMySQLConfig()
	db, err := sqlx.Open(cfg.dbAdapter, cfg.connString)
	if err != nil {
		log.Fatal(err)
	}
	defer tearDownMySQL(db)
	setupMySQL(db)
	run(cfg)
}

func TestPostgreSQL(t *testing.T) {
	cfg := getPostgreSQLConfig()
	db, err := sqlx.Open(cfg.dbAdapter, cfg.connString)
	if err != nil {
		log.Fatal(err)
	}

	defer tearDownPostgreSQL(db)
	setupPostgreSQL(db)
	run(cfg)
}
