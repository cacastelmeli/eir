package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/cacastelmeli/eir/context"
)

type contentTemplate struct {
	*template.Template
	ctx *context.Context
}

func newContentTemplate(ctx *context.Context) *contentTemplate {
	tpl := template.New("content")

	// Register custom pipes
	tpl.Funcs(template.FuncMap{
		// Func to prepend module path
		"mod_path": func(path string) string {
			return fmt.Sprintf("%s/%s", ctx.ModulePath(), path)
		},
	})

	return &contentTemplate{
		Template: tpl,
		ctx:      ctx,
	}
}

func (template *contentTemplate) Compile(content string) (string, error) {
	tpl := template.New("")
	_, err := tpl.Parse(content)

	if err != nil {
		return "", err
	}

	resultContent := bytes.NewBufferString("")
	err = tpl.Execute(resultContent, content)

	return resultContent.String(), err
}
