package environment

import (
	"os"
)

func GetEnvironmentValue(envVariable string) (value string) {
	value, exists := os.LookupEnv(envVariable)

	if !exists {
		return ""
	}

	return value
}
