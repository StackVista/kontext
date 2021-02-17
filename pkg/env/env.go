package env

import (
	"os"
	"strings"
)

// Environment is a map of environment variables.
type Environment map[string]string

// GetEnvironment converts the operating system environment variables to a map
func GetEnvironment() Environment {
	m := make(Environment)
	for _, kv := range os.Environ() {
		k, v := splitEnvKeyValue(kv)
		m[k] = v
	}

	return m
}

func splitEnvKeyValue(kv string) (string, string) {
	switch {
	case kv == "":
		return "", ""
	case strings.HasPrefix(kv, "="):
		k, v := splitEnvKeyValue(kv[1:])
		return "=" + k, v
	case strings.Contains(kv, "="):
		parts := strings.SplitN(kv, "=", 2)
		return parts[0], parts[1]
	default:
		return kv, ""
	}
}
