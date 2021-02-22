package konfig

import (
	"github.com/stackvista/kontext/pkg/env"
)

const (
	PreviousKubeConfigVar = "KONTEXT_KUBECONFIG_OLD"
	KontextLoadedVar      = "KONTEXT_CONFIG"
	KubeConfigVar         = "KUBECONFIG"
)

type CurrentKontext struct {
	CurrentKubeConfig  *string
	LoadedKontext      *string
	PreviousKubeConfig *string
}

func LoadCurrentKontext() CurrentKontext {
	e := env.GetEnvironment()

	return CurrentKontext{
		CurrentKubeConfig:  fromEnv(e, KubeConfigVar),
		LoadedKontext:      fromEnv(e, KontextLoadedVar),
		PreviousKubeConfig: fromEnv(e, PreviousKubeConfigVar),
	}
}

func fromEnv(e env.Environment, k string) *string {
	if v, ok := e[k]; ok {
		return &v
	}

	return nil
}

func BuildKontextExport(pathToKonfig string) (env.Export, error) {
	currentKontext := LoadCurrentKontext()

	if currentKontext.LoadedKontext == nil {
		return buildWithoutPreviousKontext(pathToKonfig, currentKontext)
	}

	if currentKontext.PreviousKubeConfig != nil {
		// Have a currKubeConfig and an oldKubeConfig
		// Keep the oldKubeConfig as is.
		exp := env.Export{}
		exp.Add(KubeConfigVar, env.Join(pathToKonfig, *currentKontext.PreviousKubeConfig))
		exp.Add(KontextLoadedVar, pathToKonfig)
		exp.Add(PreviousKubeConfigVar, *currentKontext.PreviousKubeConfig)

		return exp, nil
	}

	// Kontext loaded, but no previous kubeconfig found, this means that previously KUBECONFIG was unset
	kubeHomeConfig, err := HomeKubeConfig()
	if err != nil {
		return nil, err
	}

	exp := env.Export{}
	exp.Add(KubeConfigVar, env.Join(pathToKonfig, kubeHomeConfig))
	exp.Add(KontextLoadedVar, pathToKonfig)
	exp.Remove(PreviousKubeConfigVar)

	return exp, nil
}

func buildWithoutPreviousKontext(pathToKonfig string, currentKontext CurrentKontext) (env.Export, error) {
	exp := env.Export{}

	if currentKontext.CurrentKubeConfig == nil {
		// no KUBECONFIG env var set, default to `~/.kube/config`
		kubeConfig, err := HomeKubeConfig()
		if err != nil {
			return nil, err
		}

		exp.Add(KubeConfigVar, env.Join(pathToKonfig, kubeConfig))
		exp.Add(KontextLoadedVar, pathToKonfig)
		exp.Remove(PreviousKubeConfigVar)

		return exp, nil
	}

	exp.Add(PreviousKubeConfigVar, *currentKontext.CurrentKubeConfig)
	exp.Add(KontextLoadedVar, pathToKonfig)
	exp.Add(KubeConfigVar, env.Join(pathToKonfig, *currentKontext.CurrentKubeConfig))

	return exp, nil
}

func UnsetKontextExport() env.Export {
	exp := env.Export{}

	oldKubeConfig, err := env.FindEnvironment(PreviousKubeConfigVar)
	if err != nil {
		// Just remove the exported KUBECONFIG as originally it was not set
		exp.Remove(KubeConfigVar)
	} else {
		exp.Remove(PreviousKubeConfigVar)
		exp.Add(KubeConfigVar, oldKubeConfig)
	}

	return exp
}
