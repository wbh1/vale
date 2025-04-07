package code

import (
	"regexp"

	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/wbh1/vale/v3/internal/core"
)

func JavaScript() *Language {
	return &Language{
		Delims: regexp.MustCompile(`//|/\*\*?|\*/`),
		Parser: javascript.GetLanguage(),
		//Cutset:  " *",
		Queries: []core.Scope{{Name: "", Expr: "(comment) @comment", Type: ""}},
		Padding: cStyle,
	}
}
