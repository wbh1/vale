package code

import (
	"regexp"

	"github.com/smacker/go-tree-sitter/typescript/tsx"
	"github.com/wbh1/vale/v3/internal/core"
)

func Tsx() *Language {
	return &Language{
		Delims:  regexp.MustCompile(`//|/\*|\*/`),
		Parser:  tsx.GetLanguage(),
		Queries: []core.Scope{{Name: "", Expr: "(comment) @comment", Type: ""}},
		Padding: cStyle,
	}
}
