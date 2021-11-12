package template

import (
	"bytes"
	"regexp"
	"text/template"

	"github.com/cacastelmeli/eir/context"
)

var (
	templateExtRegexp = regexp.MustCompile(".tpl$")
)

type filenameTemplate struct {
	*template.Template
	ctx *context.Context
}

func newFilenameTemplate(ctx *context.Context) *filenameTemplate {
	tpl := template.New("filename")

	// Register custom delimiters
	tpl.Delims("[", "]")

	return &filenameTemplate{
		Template: tpl,
		ctx:      ctx,
	}
}

func (template *filenameTemplate) Compile(filename string) (string, error) {
	tpl := template.New("")
	_, err := tpl.Parse(filename)

	if err != nil {
		return "", err
	}

	filenameBuffer := bytes.NewBufferString("")
	err = tpl.Execute(filenameBuffer, template.ctx.FileNames())

	if err != nil {
		return "", err
	}

	resultFilename := filenameBuffer.String()

	return templateExtRegexp.ReplaceAllString(resultFilename, ""), nil
}
