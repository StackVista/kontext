package konfig

import (
	"os"
	"path"
)

func HomeKubeConfig() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".kube/config"), nil
}
