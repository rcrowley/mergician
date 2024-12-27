package html

import (
	"testing"

	"golang.org/x/net/html/atom"
)

func TestAllIsAtomHasAttr(t *testing.T) {
	testFindAll(t, All(IsAtom(atom.Article), HasAttr("class", "body")), 1)
	testFindAll(t, All(IsAtom(atom.Article), HasAttr("id", "id")), 0)
}

func TestAnyIsAtomHasAttr(t *testing.T) {
	testFindAll(t, Any(IsAtom(atom.Article), HasAttr("class", "body")), 1)
	testFindAll(t, Any(IsAtom(atom.Article), HasAttr("id", "id")), 1)
	testFindAll(t, Any(IsAtom(atom.Div), HasAttr("id", "id")), 0)
}

func TestHasAttr(t *testing.T) {
	testFindAll(t, HasAttr("class", "body"), 1)
	testFindAll(t, HasAttr("id", "id"), 0)
}

func TestIsAtom(t *testing.T) {
	testFindAll(t, IsAtom(atom.P), 2)
}

func TestIsWhitespace(t *testing.T) {
	testFindAll(t, IsWhitespace, 18)
}

func TestMatch(t *testing.T) {
	testFindAll(t, Match(NewNode(atom.Article, "class", "body")), 1)
	testFindAll(t, Match(NewNode(atom.Div, "id", "id")), 0)
}

func TestNot(t *testing.T) {
	testFindAll(t, Not(IsAtom(atom.P)), 39)
}

func testFindAll(t *testing.T, f func(*Node) bool, i int) {
	t.Helper()
	n, err := ParseFile("template.html")
	if err != nil {
		t.Fatal(err)
	}
	if found := FindAll(n, f); len(found) != i {
		t.Fatal(found)
	}
}
