package env

import (
	"fmt"
	"os"
	"strings"
)

// Environment is a map of environment variables.
type Environment map[string]string

type ErrNoSuchEnvVar struct {
	Var string
}

func (e ErrNoSuchEnvVar) Error() string {
	return fmt.Sprintf("No environment variable '%s' defined", e.Var)
}

// FindEnvironment finds the value for the given key in the environment variables
func FindEnvironment(key string) (string, error) {
	for _, kv := range os.Environ() {
		k, v := splitEnvKeyValue(kv)
		if k == key {
			return v, nil
		}
	}

	return "", ErrNoSuchEnvVar{Var: key}
}

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

func Join(values ...string) string {
	return strings.Join(values, ":")
}
