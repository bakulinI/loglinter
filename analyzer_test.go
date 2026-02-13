package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bakulinI/loglinter/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	testdataPath := filepath.Join(dir, "testdata")

	analysistest.Run(t, testdataPath, analyzer.Analyzer, "testdata")
}
