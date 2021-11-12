package codemod

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"

	"github.com/cacastelmeli/eir/context"
)

const (
	// Skip object resolution for fast path
	optimizedParserMode parser.Mode = parser.ParseComments | parser.SkipObjectResolution
)

type TransformFunc func(fileSet *token.FileSet, fileNode *ast.File, ctx *context.Context) (ast.Node, error)

// Transformer utility for code modification
type Transformer struct {
	ctx     *context.Context
	fileSet *token.FileSet
}

func NewTransformer(ctx *context.Context) *Transformer {
	return &Transformer{
		ctx:     ctx,
		fileSet: token.NewFileSet(),
	}
}

func (transformer *Transformer) TransformFile(filename string, transformFunc TransformFunc) error {
	// Parse given `filename` into an AST
	fileNode, err := parser.ParseFile(transformer.fileSet, filename, nil, optimizedParserMode)

	if err != nil {
		return err
	}

	resultFileNode, err := transformFunc(transformer.fileSet, fileNode, transformer.ctx)

	if err != nil {
		return err
	}

	if file, err := os.Create(filename); err != nil {
		return err
	} else {
		defer file.Close()

		// Transform file
		return format.Node(file, transformer.fileSet, resultFileNode)
	}
}
