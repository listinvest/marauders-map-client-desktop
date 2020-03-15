package tools

import (
	"marauders-map-client-desktop/tools/string_tools"
	"testing"
)

func TestCleanWhisteSpaces(t *testing.T) {
	if "foo bar" != string_tools.CleanWhiteSpaces("   foo    bar    ") {
		t.Fail()
	}
}
