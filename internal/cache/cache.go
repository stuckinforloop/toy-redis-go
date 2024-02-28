package cache

import (
	"time"
)

type Node struct {
	Value      string
	CreatedAt  time.Time
	Expiration time.Duration
}

// TODO: Implement LRU cache
var cache map[string]Node

func init() {
	cache = make(map[string]Node)
}

func Get(key string) (string, bool) {
	node, ok := cache[key]
	if !ok {
		return "", false
	}

	// if created_at not present, no need to check for expiration
	if node.CreatedAt.IsZero() {
		return node.Value, ok
	}

	timeElapsed := time.Now().Sub(node.CreatedAt)
	if timeElapsed > node.Expiration {
		// remove the key value from cache -- passive expiry
		delete(cache, key)

		return "", false
	}

	return node.Value, ok
}

func Set(key string, node Node) error {
	cache[key] = node

	return nil
}
