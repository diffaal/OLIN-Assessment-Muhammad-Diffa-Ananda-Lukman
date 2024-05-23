package helper

import (
	"fmt"
)

func MapKeyMaker(triple []int) string {
	var key string
	for _, v := range triple {
		el := fmt.Sprintf("%d", v)
		key += el
	}
	return key
}
