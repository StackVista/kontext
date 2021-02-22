package konfig

import (
	"fmt"
	"testing"

	"github.com/hierynomus/go-testenv"
	"github.com/stackvista/kontext/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestBuildKontextExport(t *testing.T) {
	tests := map[string]struct {
		EnvMap        map[string]string
		KubeCtx       *string
		KontextOld    *string
		KontextLoaded *string
	}{
		"Fresh env": {
			EnvMap:        map[string]string{},
			KubeCtx:       test.StringP(fmt.Sprintf("/mydir/.kontext:%s/.kube/config", test.Home(t))),
			KontextOld:    nil,
			KontextLoaded: test.StringP("/mydir/.kontext"),
		},
		"Kontext loaded, current kubeconfig": {
			EnvMap:        map[string]string{KontextLoadedVar: "/dummy/.kontext", KubeConfigVar: fmt.Sprintf("/dummy/.kontext:%s/.kube/config", test.Home(t))},
			KubeCtx:       test.StringP(fmt.Sprintf("/mydir/.kontext:%s/.kube/config", test.Home(t))),
			KontextOld:    nil,
			KontextLoaded: test.StringP("/mydir/.kontext"),
		},
		"No kontext loaded, current kubeconfig": {
			EnvMap:        map[string]string{KubeConfigVar: "/a/random/.kube/config"},
			KubeCtx:       test.StringP("/mydir/.kontext:/a/random/.kube/config"),
			KontextOld:    test.StringP("/a/random/.kube/config"),
			KontextLoaded: test.StringP("/mydir/.kontext"),
		},
		"Kontext loaded, current and previous kubeconfig": {
			EnvMap:        map[string]string{KontextLoadedVar: "/dummy/.kontext", KubeConfigVar: "/dummy/.kontext/a/random/.kube/config", PreviousKubeConfigVar: "/a/random/.kube/config"},
			KubeCtx:       test.StringP("/mydir/.kontext:/a/random/.kube/config"),
			KontextOld:    test.StringP("/a/random/.kube/config"),
			KontextLoaded: test.StringP("/mydir/.kontext"),
		},
	}

	for n, s := range tests {
		e := s // Assign loop var
		t.Run(n, func(t *testing.T) {
			defer testenv.PatchEnv(t, e.EnvMap)()
			exp, err := BuildKontextExport("/mydir/.kontext")
			assert.NoError(t, err)
			assert.Equal(t, e.KubeCtx, exp[KubeConfigVar])
			assert.Equal(t, e.KontextOld, exp[PreviousKubeConfigVar])
			assert.Equal(t, e.KontextLoaded, exp[KontextLoadedVar])
		})
	}
}
