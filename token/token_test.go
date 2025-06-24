package token

import "testing"

func TestLookup(t *testing.T) {
	test := LookupIdent("let")
	if test == "" {
		t.Fail()
	}
}
