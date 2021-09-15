// Package clientts generates client-side TypeScript library
package clientts

import (
	"fmt"
	"path/filepath"

	"github.com/go-generalize/api_gen/pkg/common"
	"github.com/go-generalize/api_gen/pkg/parser"
	go2tsgenerator "github.com/go-generalize/go2ts/pkg/generator"
	go2tsparser "github.com/go-generalize/go2ts/pkg/parser"
	"golang.org/x/xerrors"
)

// GenerateTypes generates request/response types in TypeScript
func (g *generator) GenerateTypes(fn func(relPath, code string) error) error {
	err := g.generateTypes(g.root, fn)

	if err != nil {
		return xerrors.Errorf("failed to generate request/response types: %w", err)
	}

	return nil
}

func (g *generator) generateTypes(gr *parser.Group, fn func(relPath, code string) error) error {
	for _, child := range gr.Children {
		if err := g.generateTypes(child, fn); err != nil {
			return xerrors.Errorf("an error occurred in %s: %w", gr.Dir, err)
		}
	}

	if len(gr.Endpoints) == 0 {
		return nil
	}

	psr, err := go2tsparser.NewParser(gr.Dir, common.ParserFilter)

	if err != nil {
		return xerrors.Errorf("failed to parse %s: %w", gr.Dir, err)
	}

	parsed, err := psr.Parse()

	if err != nil {
		return xerrors.Errorf("failed to parse %s: %w", gr.Dir, err)
	}

	code := go2tsgenerator.NewGenerator(parsed).Generate()

	relative := gr.GetFullPath(string(filepath.Separator), func(rawPath, path, placeholder string) string {
		return rawPath
	})

	code = fmt.Sprintf(headerComment, g.AppVersion) + code

	p := filepath.Join("classes", relative, "types.ts")
	if err := fn(p, code); err != nil {
		return xerrors.Errorf("failed to save %s: %w", p, err)
	}

	return nil
}
