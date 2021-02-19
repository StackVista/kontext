package shell

import (
	"fmt"
	"testing"

	testenv "github.com/hierynomus/go-testenv"
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

func TestShellDetectionThroughEnv(t *testing.T) {
	shells := []string{"/bin/bash", "/usr/local/bin/zsh"}
	for _, shell := range shells {
		sh := shell
		t.Run(fmt.Sprintf("Detect %s through env", sh), func(t *testing.T) {
			defer testenv.PatchEnv(t, map[string]string{
				"SHELL": sh,
			})()

			s, err := DetectShell("-")
			assert.NoError(t, err)
			assert.NotNil(t, s)
		})
	}
}
