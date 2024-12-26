package files

/*
func TestAll(t *testing.T) {
	l, err := All([]string{"."}, []string{}, []string{".html", ".md", ".zip"}) // skipping ".htm"
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.md", "test/html.html", "test2/html.html", "zip.zip"}) {
		t.Fatal(pathnames)
	}
}

func TestAllHTML(t *testing.T) {
	l, err := AllHTML([]string{"."}, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.html", "test/html.html", "test/test/htm.htm", "test2/html.html"}) {
		t.Fatal(pathnames)
	}
}

func TestAllHTMLExclude(t *testing.T) {
	l, err := AllHTML([]string{"."}, []string{"test"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.html", "test2/html.html"}) {
		t.Fatal(pathnames)
	}
}

func TestAllHTMLExclude2(t *testing.T) {
	l, err := AllHTML([]string{"."}, []string{"test", "test2"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.html"}) {
		t.Fatal(pathnames)
	}
}

func TestAllInputs(t *testing.T) {
	l, err := AllInputs([]string{"."}, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.md", "test/html.html", "test/test/htm.htm", "test2/html.html", "zip.zip"}) {
		t.Fatal(pathnames)
	}
}

func TestAllMarkdown(t *testing.T) {
	l, err := All([]string{"."}, []string{}, []string{".md"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"md.md"}) {
		t.Fatal(pathnames)
	}
}

func TestAllSub(t *testing.T) {
	l, err := All([]string{"test"}, []string{}, []string{".htm", ".html"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"test/html.html", "test/test/htm.htm"}) {
		t.Fatal(pathnames)
	}
}

func TestAllSub2(t *testing.T) {
	l, err := All([]string{"test", "test2"}, []string{}, []string{".htm", ".html"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"test/html.html", "test/test/htm.htm", "test2/html.html"}) {
		t.Fatal(pathnames)
	}
}
*/
