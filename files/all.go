package files

import (
	"io/fs"
	"os"
	"path/filepath"
)

func All(includes, excludes, extensions []string) (*List, error) {
	l := &List{}
	for _, include := range includes {
		if err := fs.WalkDir(
			os.DirFS(include),
			".",
			func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if !d.Type().IsRegular() {
					for _, e := range excludes {
						if e == path {
							return fs.SkipDir
						}
					}
					return nil
				}

				ext := filepath.Ext(path)
				for _, e := range extensions {
					if e == ext {
						l.Add(path)
					}
				}

				return nil
			},
		); err != nil {
			return nil, err
		}
	}
	return l, nil
}

func AllHTML(includes, excludes []string) (*List, error) {
	return All(includes, excludes, []string{".htm", ".html"})
}

func AllInputs(includes, excludes []string) (*List, error) {
	return All(includes, excludes, []string{".htm", ".html", ".md", ".zip"})
}
