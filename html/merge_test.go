package html

import (
	"testing"

	"golang.org/x/net/html/atom"
)

func TestMergeCustom(t *testing.T) {
	n0, err := ParseFile("template.html")
	if err != nil {
		t.Fatal(err)
	}
	n1, err := ParseFile("article.html")
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
	expected := `<!DOCTYPE html>
<html lang="en">
<head>
<link href="template.css" rel="stylesheet"/>
<meta charset="utf-8"/>
<meta content="width=device-width,initial-scale=1" name="viewport"/>
<title>My cool webpage — Website</title>
</head>
<body>
<header><h1>Website</h1></header>
<br/><!-- explicit self-closing -->
<article class="body">Things</article>
<br/><!-- implied self-closing -->
<footer><p>© 2023</p></footer>
</body>
</html>
`
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(String(n))
}

func TestMergeCustomEmpty(t *testing.T) {
	n0, err := ParseFile("template.html")
	if err != nil {
		t.Fatal(err)
	}
	n1, err := ParseFile("article.html")
	if err != nil {
		t.Fatal(err)
	}
	n, err := Merge([]*Node{n0, n1}, []Rule{})
	if err != nil {
		t.Fatal(err)
	}

	actual := String(n)
	expected := `<!DOCTYPE html>
<html lang="en">
<head>
<link href="template.css" rel="stylesheet"/>
<meta charset="utf-8"/>
<meta content="width=device-width,initial-scale=1" name="viewport"/>
<title>My cool webpage — Website</title>
</head>
<body>
<header><h1>Website</h1></header>
<br/><!-- explicit self-closing -->
<article class="body">
<p>Overwritten.</p>
</article>
<br/><!-- implied self-closing -->
<footer><p>© 2023</p></footer>
</body>
</html>
`
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(String(n))
}

func TestMergeDefault(t *testing.T) {
	n0, err := ParseFile("template.html")
	if err != nil {
		t.Fatal(err)
	}
	n1, err := ParseFile("article.html")
	if err != nil {
		t.Fatal(err)
	}
	n, err := Merge([]*Node{n0, n1}, DefaultRules())
	if err != nil {
		t.Fatal(err)
	}

	actual := String(n)
	expected := `<!DOCTYPE html>
<html lang="en">
<head>
<link href="template.css" rel="stylesheet"/>
<meta charset="utf-8"/>
<meta content="width=device-width,initial-scale=1" name="viewport"/>
<title>My cool webpage — Website</title>
</head>
<body>
<header><h1>Website</h1></header>
<br/><!-- explicit self-closing -->
<article class="body">
<h1>Things</h1>
<p>Stuff</p>
</article>
<br/><!-- implied self-closing -->
<footer><p>© 2023</p></footer>
</body>
</html>
` // TODO deal with the deranged whitespace the merge algorithm produces
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(String(n))
}
