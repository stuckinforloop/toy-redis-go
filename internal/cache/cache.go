package cache

import (
	"fmt"
)

// TODO: Implement LRU cache
var cache map[string]string

func init() {
	cache = make(map[string]string)
}

func Get(key string) (string, bool) {
	val, ok := cache[key]

	return val, ok
}

func Set(key string, value any) error {
	cache[key] = fmt.Sprint(value)

	return nil
}
