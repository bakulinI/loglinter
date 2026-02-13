package analyzer

import (
	"go/ast"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

// Analyzer экспортируется для singlechecker / golangci-lint
var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages",
	Run:  run,
}

// Список запрещённых ключевых слов
var sensitiveWords = []string{
	"password",
	"token",
	"apikey",
	"secret",
}

func run(pass *analysis.Pass) (interface{}, error) {
	var checkSensitive func(expr ast.Expr) bool
	checkSensitive = func(expr ast.Expr) bool {
		switch e := expr.(type) {
		case *ast.BasicLit:
			m, err := strconv.Unquote(e.Value)
			if err != nil {
				return false
			}
			for _, word := range sensitiveWords {
				if strings.Contains(strings.ToLower(m), word) {
					return true
				}
			}
		case *ast.BinaryExpr:

			return checkSensitive(e.X) || checkSensitive(e.Y)
		case *ast.Ident:

			for _, word := range sensitiveWords {
				if strings.Contains(strings.ToLower(e.Name), word) {
					return true
				}
			}
		}
		return false
	}

	for _, file := range pass.Files {

		ast.Inspect(file, func(n ast.Node) bool {

			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			if sel.Sel.Name != "Info" && sel.Sel.Name != "Debug" && sel.Sel.Name != "Error" && sel.Sel.Name != "Warn" {
				return true
			}

			ident, ok := sel.X.(*ast.Ident)
			if !ok || ident.Name != "slog" {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			arg := call.Args[0]

			if lit, ok := arg.(*ast.BasicLit); ok {
				msg, err := strconv.Unquote(lit.Value)
				if err == nil && msg != "" {

					// 1. lowercase
					r := []rune(msg)[0]
					if unicode.IsUpper(r) {
						pass.Reportf(lit.Pos(), "log message must start with lowercase letter")
					}

					// 2. english only
					for _, r := range msg {
						if unicode.In(r, unicode.Cyrillic) {
							pass.Reportf(lit.Pos(), "log message must be in English")
							break
						}
					}

					// 3. no emoji
					for _, r := range msg {
						if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
							continue
						}
						pass.Reportf(lit.Pos(), "log message must not contain special characters or emoji")
						break
					}
				}
			}

			// 4. sensitive data
			if checkSensitive(arg) {
				pass.Reportf(arg.Pos(), "log message contains sensitive data")
			}

			return true
		})
	}

	return nil, nil
}
