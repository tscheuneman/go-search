package utils

import (
	"os"
)

func GetEnvValue(envVar string, defaultValue string) string {
	envVal := os.Getenv(envVar)
	if envVal == "" {
		envVal = defaultValue
	}

	return envVal
}

func GetEnvValueOrDie(envVar string) string {
	envVal := os.Getenv(envVar)
	if envVal == "" {
		panic("Error: " + envVar + " does not exist")
	}

	return envVal
}
