package utils

import "os"

// Get env variable. Accepts second argument as default value.
func GetEnv(env string, defaultVal ...string) string {
	val, ok := os.LookupEnv(env)

	if !ok {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		} else {
			return ""
		}
	}
	return val
}
