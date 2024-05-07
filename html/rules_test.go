package html

import (
	"testing"

	"golang.org/x/net/html/atom"
)

func TestDefaultRules(t *testing.T) {
	rules := DefaultRules()
	for _, rule := range rules {
		_ = rule // t.Log(rule)
	}
}

func TestParseRule(t *testing.T) {
	for _, s := range []string{
		`<body> = <body>`,
		`<div class="body"> = <body>`,
		`<div class="body"> = <body class="body">`,
		`<div class="body"> = <div class="body">`,
	} {
		if rule, err := ParseRule(s); err != nil || rule.String() != s {
			t.Fatal(rule, s, err)
		}
	}
}

func TestRuleString(t *testing.T) {
	rule := Rule{NewNode(atom.Div, "class", "body"), "=", NewNode(atom.Body)}
	if rule.String() != `<div class="body"> = <body>` {
		t.Fatal(rule)
	}
}
