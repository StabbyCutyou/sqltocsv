package converters

var converters = make(map[string]Converter)

func init() {
	register("postgres", pg{})
	register("mysql", mySQL{})
}

// Converter is the adapter for handling datatypes from different databases
type Converter interface {
	ColumnToString(interface{}) (string, error)
}

func register(name string, c Converter) {
	converters[name] = c
}

// GetConverter will return a given converter by name from the registry
func GetConverter(name string) Converter {
	return converters[name]
}
