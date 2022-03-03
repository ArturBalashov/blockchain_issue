package helpers

import (
	"os"
	"strconv"
)

// EnvToInt convert env variable to int
func EnvToInt(value *int, key string, defaultValue int) {
	*value = defaultValue

	envValue, exists := os.LookupEnv(key)
	if !exists {
		return
	}

	if res, err := strconv.Atoi(envValue); err == nil {
		*value = res
	}
}

// EnvToBool convert env variable to bool
func EnvToBool(value *bool, key string, defaultValue bool) {
	*value = defaultValue

	envValue, exists := os.LookupEnv(key)
	if !exists {
		return
	}

	if res, err := strconv.ParseBool(envValue); err == nil {
		*value = res
	}
}

// EnvToString convert env variable to string
func EnvToString(value *string, key string, defaultValue string) {
	*value = getEnv(key, defaultValue)
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
