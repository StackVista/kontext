package shell

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellDetection(t *testing.T) {
	shells := []string{"-bash", "-/bin/bash", "-/usr/local/bin/bash", "-zsh", "-/bin/zsh", "-/usr/local/bin/zsh", "bash", "zsh"}
	for _, shell := range shells {
		sh := shell
		t.Run(fmt.Sprintf("Detect %s", sh), func(t *testing.T) {
			s, err := DetectShell(sh)
			assert.NoError(t, err)
			assert.NotNil(t, s)
		})
	}
}
