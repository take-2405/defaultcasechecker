package defaultcasechecker_test

import (
	"testing"

	"github.com/take-2405/defaultcasechecker"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	// testdataディレクトリ内のソースコードを解析する
	analysistest.Run(t, testdata, defaultcasechecker.Analyzer, "a")
	analysistest.Run(t, testdata, defaultcasechecker.Analyzer, "b")
	analysistest.Run(t, testdata, defaultcasechecker.Analyzer, "c")
}
