package util

import (
	"fmt"
	"strconv"
)

func Round(f float64, i int) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.FormatInt(int64(i), 10)+"f", f), 64)
	return value
}
