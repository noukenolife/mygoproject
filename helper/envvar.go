package helper

import "os"

func GetenvWithDefaultValue(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}

func Getenv(key string) string {
	return GetenvWithDefaultValue(key, "")
}
