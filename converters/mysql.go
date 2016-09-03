package converters

import (
	"fmt"
	"strconv"
)

type mySQL struct{}

func (mySQL) ColumnToString(col interface{}) (string, error) {
	// In all my testing, it seems like the mysql db adapter always returns []byte
	switch col.(type) {
	case []byte:
		byts := col.([]byte) // MySQL driver does not make it easy to deal with bit fields
		if len(byts) == 1 && (byts[0] == 0 || byts[0] == 1) {
			return strconv.Itoa(int(byts[0])), nil
		}
		return string(col.([]byte)), nil
	default:
		// Need to handle anything that ends up here
		return "", fmt.Errorf("Unknown column type %v", col)
	}
}
