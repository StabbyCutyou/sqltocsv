package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	// blank to add the mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Config is
type Config struct {
	dbAdapter  string
	connString string
	sqlQuery   string
	outputFile string
}

func main() {
	cfg := getConfig()
	db, err := sqlx.Open(cfg.dbAdapter, cfg.connString)
	if err != nil {
		log.Fatal(err)
	}

	results, err := db.Queryx(cfg.sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create(cfg.outputFile)
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(output)
	csvWriter.Comma = 0x0009
	firstLine := true
	for results.Next() {
		row, err := results.SliceScan()
		if err != nil {
			log.Fatal(err)
		}
		if firstLine {
			cols, err := results.Columns()
			if err != nil {
				log.Fatal(err)
			}
			csvWriter.Write(cols)
		}

		rowStrings := make([]string, len(row))
		// It seems for mysql, the case is always []byte of a string?
		for i, col := range row {
			//log.Print(reflect.TypeOf(col))
			switch col.(type) {
			case float64:
				rowStrings[i] = strconv.FormatFloat(col.(float64), 'f', 6, 64)
			case int64:
				rowStrings[i] = strconv.FormatInt(col.(int64), 10)
			case bool:
				rowStrings[i] = strconv.FormatBool(col.(bool))
			case []byte:
				rowStrings[i] = string(col.([]byte))
			case string:
				rowStrings[i] = col.(string)
			case time.Time:
				rowStrings[i] = col.(time.Time).String()
			case nil:
				rowStrings[i] = "NULL"
			default:
				log.Print(col)
			}
		}
		csvWriter.Write(rowStrings)
	}

	csvWriter.Flush()
	output.Close()
}

func getConfig() *Config {
	cfg := &Config{
		dbAdapter:  os.Getenv("STC_DBADAPTER"),
		connString: os.Getenv("STC_CONNSTRING"),
		sqlQuery:   os.Getenv("STC_QUERY"),
		outputFile: os.Getenv("STC_OUTPUTFILE"),
	}
	if cfg.dbAdapter == "" {
		log.Fatal("You must provide a connection string via STC_DBADAPTER")
	}
	if cfg.connString == "" {
		log.Fatal("You must provide a connection string via STC_CONNSTRING")
	}
	if cfg.sqlQuery == "" {
		log.Fatal("You must provide a query to run via STC_QUERY")
	}
	if cfg.outputFile == "" {
		log.Fatal("You must provide an output file via STC_OUTPUTFILE")
	}

	return cfg
}
