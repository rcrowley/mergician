package html

import (
	"testing"

	"golang.org/x/net/html/atom"
)

func TestDefaultRules(t *testing.T) {
	rules := DefaultRules()
	for _, rule := range rules {
		t.Log(rule)
	}
}

func TestRuleString(t *testing.T) {
	rule := Rule{NewNode(atom.Div, "class", "body"), "=", NewNode(atom.Body)}
	if rule.String() != `<div class="body"> = <body>` {
		t.Fatal(rule)
	}
}
