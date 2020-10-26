package http

import (
	"ExcelTools/tools/settings"
	"testing"
)

func TestRun(t *testing.T) {
	run := Run(&settings.HttpConfig{Port: 8090})
	if run != nil {
		t.Error(run)
	}
}
