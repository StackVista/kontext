package kube

import (
	"fmt"
	"os"
	"path"
)

type ErrNoKubeFound struct {
	FromDir string
}

func (e ErrNoKubeFound) Error() string {
	return fmt.Sprintf("No .kube found in %s or parent directories", e.FromDir)
}

const DotKube = ".kube"

func FindDotKube() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return PathToDotKube(cwd)
}

func PathToDotKube(p string) (string, error) {
	kubePath := path.Join(p, DotKube)
	ok, err := fileExists(kubePath)
	if err != nil {
		return "", err
	}

	if ok {
		return kubePath, nil
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
