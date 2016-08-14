package converters

import "fmt"

type Redshift struct{}

func (r *Redshift) ColumnToString(col interface{}) (string, error) {
	// In all my testing, it seems like the redshift db adapter always returns []byte
	switch col.(type) {
	case []byte:
		return string(col.([]byte)), nil
	default:
		// Need to handle anything that ends up here
		return "", fmt.Errorf("Unknown column type %v", col)
	}
}
