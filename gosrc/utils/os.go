package utils

import (
	"os"
	"strconv"
)

func GetEnv(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func GetEnvInt(key string, defaultValue int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	val_int, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return val_int
}
