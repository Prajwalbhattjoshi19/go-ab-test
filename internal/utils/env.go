package utils

import (
	"os"
	"strconv"
)

// GetEnv retrieves the value of the environment variable named by the key. If the variable is not present, it returns the default value provided.
func GetEnv(key string, defaultValue any) any {
	value, exists := os.LookupEnv(key)
	switch defaultValue.(type) {
	case int:
		if !exists {
			return defaultValue
		}
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return intValue
	case string:
		if !exists {
			return defaultValue
		}
		return value
	default:
		return value
	}
}
