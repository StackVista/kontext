package kube

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFindDotKubeInCurrentPath(t *testing.T) {
	assert.NoError(t, ioutil.WriteFile(".kube", []byte("foo"), 0600))
	defer os.Remove(".kube")

	p, err := PathToDotKube(".")
	assert.NoError(t, err)
	assert.Equal(t, ".kube", p)
}
