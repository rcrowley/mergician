package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHandlerOK(t *testing.T) {
	h := &Handler{Logger: log.Default()}
	req := testRequest(t)
	rw := httptest.NewRecorder()
	h.ServeHTTP(rw, req)
	body := rw.Body.String()
	if rw.Code != http.StatusOK || body != testTemplatePlusArticleHTML(t) {
		t.Fatal(rw.Code, rw.HeaderMap, body)
	}
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func testRequest(t *testing.T) *http.Request {
	b := &bytes.Buffer{}
	mw := multipart.NewWriter(b)
	var (
		w   io.Writer
		err error
	)
	for i, filename := range []string{"template.html", "article.html"} {
		if w, err = mw.CreateFormFile(fmt.Sprintf("in[%d]", i), filename); err != nil {
			t.Fatal(err)
		}
		f, err := os.Open(filepath.Join("../html/testdata", filename))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		if _, err := io.Copy(w, f); err != nil {
			t.Fatal(err)
		}
	}
	mw.Close()

	req := httptest.NewRequest("POST", "/", b)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	return req
}

func testTemplatePlusArticleHTML(t *testing.T) string {
	b, err := os.ReadFile("../html/testdata/template+article.html")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
