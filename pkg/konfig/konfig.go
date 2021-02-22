package konfig

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const ConfigFileName = ".konfig"

type ErrNoKubeFound struct {
	FromDir string
}

func (e ErrNoKubeFound) Error() string {
	return fmt.Sprintf("No '%s' found in '%s' or parent directories", ConfigFileName, e.FromDir)
}

func FindKontextConfig() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return PathToKontextConfig(cwd)
}

func PathToKontextConfig(p string) (string, error) {
	absPath, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}

	for absPath != "/" && absPath != "." {
		kubePath := path.Join(absPath, ConfigFileName)
		ok, err := fileExists(kubePath)
		if err != nil {
			return "", err
		}

		if ok {
			return kubePath, nil
		}

		absPath = filepath.Dir(absPath)
	}

	return "", ErrNoKubeFound{FromDir: p}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return !info.IsDir(), nil
}
