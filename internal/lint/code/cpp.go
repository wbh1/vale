package code

import (
	"regexp"

	"github.com/smacker/go-tree-sitter/cpp"
	"github.com/wbh1/vale/v3/internal/core"
)

func Cpp() *Language {
	return &Language{
		Delims:  regexp.MustCompile(`//|/\*!?|\*/`),
		Parser:  cpp.GetLanguage(),
		Queries: []core.Scope{{Name: "", Expr: "(comment) @comment", Type: ""}},
		Padding: cStyle,
	}
}
