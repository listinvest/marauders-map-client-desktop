package tools_tests

import (
	"marauders-map-client-desktop/tools"
	"testing"
)

func TestCleanWhisteSpaces(t *testing.T) {
	if "foo bar" != tools.CleanWhiteSpaces("   foo    bar    ") {
		t.Fail()
	}
}
