package html

import (
	"os"
	"testing"

	"golang.org/x/net/html/atom"
)

func TestMergeCustom(t *testing.T) {
	n0, err := ParseFile("testdata/template.html")
	if err != nil {
		t.Fatal(err)
	}
	n1, err := ParseFile("testdata/article.html")
	if err != nil {
		t.Fatal(err)
	}
	n, err := Merge([]*Node{n0, n1}, []Rule{
		{NewNode(atom.Article, "class", "body"), "=", NewNode(atom.H1)},
	})
	if err != nil {
		t.Fatal(err)
	}

	actual := String(n)
	expected := testTemplatePlusArticleCustomHTML(t)
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(String(n))
}

func TestMergeNoRules(t *testing.T) {
	n0, err := ParseFile("testdata/template.html")
	if err != nil {
		t.Fatal(err)
	}
	n1, err := ParseFile("testdata/article.html")
	if err != nil {
		t.Fatal(err)
	}
	n, err := Merge([]*Node{n0, n1}, []Rule{})
	if err != nil {
		t.Fatal(err)
	}

	actual := String(n)
	expected := testTemplatePlusArticleNoRulesHTML(t)
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(String(n))
}

func TestMergeDefault(t *testing.T) {
	n0, err := ParseFile("testdata/template.html")
	if err != nil {
		t.Fatal(err)
	}
	n1, err := ParseFile("testdata/article.html")
	if err != nil {
		t.Fatal(err)
	}
	n, err := Merge([]*Node{n0, n1}, DefaultRules())
	if err != nil {
		t.Fatal(err)
	}

	actual := String(n)
	expected := testTemplatePlusArticleHTML(t)
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(String(n))
}

func TestMergePlusEquals(t *testing.T) {
	n0, err := ParseFile("testdata/template.html")
	if err != nil {
		t.Fatal(err)
	}
	n1, err := ParseFile("testdata/article.html")
	if err != nil {
		t.Fatal(err)
	}
	n, err := Merge([]*Node{n0, n1}, []Rule{
		{NewNode(atom.Article, "class", "body"), "+=", NewNode(atom.Body)},
	})
	if err != nil {
		t.Fatal(err)
	}

	actual := String(n)
	expected := testTemplatePlusEqualsArticleHTML(t)
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(String(n))
}

func testTemplatePlusArticleCustomHTML(t *testing.T) string {
	b, err := os.ReadFile("testdata/template+article.custom.html")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func testTemplatePlusArticleHTML(t *testing.T) string {
	b, err := os.ReadFile("testdata/template+article.html")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func testTemplatePlusArticleNoRulesHTML(t *testing.T) string {
	b, err := os.ReadFile("testdata/template+article.no-rules.html")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func testTemplatePlusEqualsArticleHTML(t *testing.T) string {
	b, err := os.ReadFile("testdata/template+=article.html")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
