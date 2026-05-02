package utils

import (
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

// GetEnv retrieves the value of the environment variable named by the key. If the variable is not present, it returns the default value provided.
func GetEnv[T constraints.Ordered](key string, defaultValue T) T {
	value, exists := os.LookupEnv(key)
	switch any(defaultValue).(type) {
	case int:
		if !exists {
			return defaultValue
		}
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return T(intValue)
	case string:
		if !exists {
			return defaultValue
		}
		return any(value).(T)
	default:
		return defaultValue
	}
}
