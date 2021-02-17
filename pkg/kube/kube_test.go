package kube

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/tj/assert"
)

func TestShouldFindDotKubeInCurrentPath(t *testing.T) {
	assert.NoError(t, ioutil.WriteFile(".kube", []byte("foo"), 0600))
	defer os.Remove(".kube")

	p, err := PathToDotKube(".")
	assert.NoError(t, err)
	assert.Equal(t, ".kube", p)
}
