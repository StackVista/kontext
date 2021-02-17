package hook

import (
	"bytes"
	"html/template"
)

type Context struct {
	// SelfPath is the path to the executable
	SelfPath string
	// Name is the name of the executable progam
	Name string
}

type Hook struct {
	Contents string
}

func (h *Hook) Render(hookCtx Context) (string, error) {
	hookTemplate, err := template.New("hook").Parse(h.Contents)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer([]byte{})

	err = hookTemplate.Execute(buf, hookCtx)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
