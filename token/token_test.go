package token

import "testing"

func TestLookup(t *testing.T) {
	test := LookupIdent("==")
	t.Log(test)
	if test == "" {
		t.Fail()
	}
}
