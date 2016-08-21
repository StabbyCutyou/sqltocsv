package converters

import (
	"fmt"
	"strconv"
	"time"
)

type pg struct{}

func (pg) ColumnToString(col interface{}) (string, error) {
	switch col.(type) {
	case float64:
		return strconv.FormatFloat(col.(float64), 'f', 6, 64), nil
	case int64:
		return strconv.FormatInt(col.(int64), 10), nil
	case bool:
		return strconv.FormatBool(col.(bool)), nil
	case []byte:
		return string(col.([]byte)), nil
	case string:
		return col.(string), nil
	case time.Time:
		return col.(time.Time).String(), nil
	case nil:
		return "NULL", nil
	default:
		// Need to handle anything that ends up here
		return "", fmt.Errorf("Unknown column type %v", col)
	}
}
