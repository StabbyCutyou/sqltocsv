package converters

import "fmt"

var converters = make(map[string]Converter)

func init() {
	register("postgres", &Pg{})
	register("mysql", &MySQL{})
	register("redshift", &Redshift{})
}

// Converter is the adapter for handling datatypes from different databases
type Converter interface {
	ColumnToString(interface{}) (string, error)
}

func register(name string, c Converter) {
	converters[name] = c
}

func GetConverter(name string) Converter {
	// lol go
	fmt.Println(name)
	c := converters[name]
	return c
}
