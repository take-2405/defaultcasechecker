package defaultcasechecker

import (
	"go/ast"
	"strings"

	"github.com/gostaticanalysis/comment"
	"github.com/gostaticanalysis/comment/passes/commentmap"

	"golang.org/x/tools/go/analysis"
)

var (
	checkTest     bool
	checkGenerate bool
)

const (
	CheckTest     = "check-test"
	CheckGenerate = "check-generate"
)

func init() {
	Analyzer.Flags.BoolVar(&checkTest, CheckTest, false, "check test file")
	Analyzer.Flags.BoolVar(&checkGenerate, CheckGenerate, false, "check generated file")
}

var Analyzer = &analysis.Analyzer{
	Name: "defaultcasechecker",
	Doc:  "checks for switch statements without default cases",
	Run:  run,
	Requires: []*analysis.Analyzer{
		commentmap.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	cmaps := pass.ResultOf[commentmap.Analyzer].(comment.Maps)
	for _, file := range pass.Files {
		if strings.HasSuffix(pass.Fset.File(file.Pos()).Name(), "test.go") && checkTest {
			continue
		}

		if isGeneratedFile(file) && checkGenerate {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			switchStmt, ok := n.(*ast.SwitchStmt)
			if !ok {
				return true
			}

			// check if there is a default case
			hasDefault := false
			for _, clause := range switchStmt.Body.List {
				caseClause, ok := clause.(*ast.CaseClause)
				if !ok {
					continue
				}
				if len(caseClause.List) == 0 {
					hasDefault = true
					break
				}
			}

			if !hasDefault {
				pos := pass.Fset.Position(switchStmt.Pos())
				if cmaps.IgnorePos(switchStmt.Pos(), "defaultcasechecker") {
					return true
				}
				pass.Reportf(switchStmt.Pos(), "%s default case not declared in switch statement", pos)
			}

			return true
		})
	}

	return nil, nil
}

func isGeneratedFile(file *ast.File) bool {
	for _, c := range file.Comments {
		if strings.Contains(c.Text(), "Code generated by") {
			return true
		}
	}
	return false
}
