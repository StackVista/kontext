package konfig

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFindDotKontextInCurrentPath(t *testing.T) {
	assert.NoError(t, ioutil.WriteFile(ConfigFileName, []byte("foo"), 0600))
	defer os.Remove(ConfigFileName)

	p, err := PathToKontextConfig(".")
	assert.NoError(t, err)
	assert.Equal(t, path.Join(workingDir(t), ConfigFileName), p)
}

func TestShouldFindDotKontextInParentPath(t *testing.T) {
	assert.NoError(t, ioutil.WriteFile(ConfigFileName, []byte("foo"), 0600))
	defer os.Remove(ConfigFileName)
	assert.NoError(t, os.Mkdir("tst", 0755))
	defer os.Remove("tst")

	p, err := PathToKontextConfig("tst")
	assert.NoError(t, err)
	assert.Equal(t, path.Join(workingDir(t), ConfigFileName), p)
}

func workingDir(t *testing.T) string {
	w, err := os.Getwd()
	assert.NoError(t, err)
	return w
}
