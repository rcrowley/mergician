package http

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rcrowley/mergician/html"
)

type Handler struct {
	Logger *log.Logger
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	mr, err := req.MultipartReader()
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err)
		return
	}
	var in []*html.Node
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			h.writeError(w, http.StatusBadRequest, err)
			return
		}
		n, err := html.Parse(p)
		if err != nil {
			h.writeError(w, http.StatusInternalServerError, err)
			return
		}
		in = append(in, n)
	}

	rules := html.DefaultRules() // TODO pull rules out of headers

	out, err := html.Merge(in, rules)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := html.Render(w, out); err != nil && h.Logger != nil {
		h.Logger.Print(err)
	}
}

func (h *Handler) writeError(w http.ResponseWriter, statusCode int, err error) {
	if h.Logger != nil {
		h.Logger.Output(2, fmt.Sprintf("%s\n", err))
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusBadRequest)
	if _, err = fmt.Fprintln(w, err); err != nil && h.Logger != nil {
		h.Logger.Output(2, fmt.Sprintf("%s\n", err))
	}
}
