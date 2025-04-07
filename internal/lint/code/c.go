package code

import (
	"regexp"

	"github.com/smacker/go-tree-sitter/c"
	"github.com/wbh1/vale/v3/internal/core"
)

func C() *Language {
	return &Language{
		Delims: regexp.MustCompile(`//|/\*|\*/`),
		Parser: c.GetLanguage(),
		Queries: []core.Scope{
			{Name: "", Expr: "(comment) @comment", Type: ""},
		},
		Padding: cStyle,
	}
}
