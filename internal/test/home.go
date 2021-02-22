package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Home(t *testing.T) string {
	h, err := os.UserHomeDir()
	assert.NoError(t, err)
	return h
}
