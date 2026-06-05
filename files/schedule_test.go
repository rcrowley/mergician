package files

import (
	"testing"
	"time"

	"github.com/rcrowley/mergician/html"
)

func doc(t *testing.T, body string) *html.Node {
	t.Helper()
	n, err := html.ParseString("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n<title>x</title>\n</head>\n<body>\n" + body + "\n</body>\n</html>\n")
	if err != nil {
		t.Fatal(err)
	}
	return n
}

func TestPublishStatus(t *testing.T) {
	now := time.Date(2026, 6, 4, 0, 0, 0, 0, time.UTC)

	for _, tt := range []struct {
		name string
		body string
		want State
	}{
		{"no dates", `<article class="body"><h1>Hi</h1></article>`, Published},
		{"draft token", `<article class="body draft"><h1>Hi</h1></article>`, Draft},
		{"draft alone", `<article class="draft"><h1>Hi</h1></article>`, Draft},
		{"published past", `<article class="body"><time class="published" datetime="2026-01-01 00:00:00"></time></article>`, Published},
		{"published future", `<article class="body"><time class="published" datetime="2026-12-31 00:00:00"></time></article>`, Scheduled},
		{"feed fallback future", `<article class="body"><time class="feed" datetime="2026-12-31 00:00:00"></time></article>`, Scheduled},
		{"feed fallback past", `<article class="body"><time class="feed" datetime="2026-01-01 00:00:00"></time></article>`, Published},
		{"expired", `<article class="body"><time class="expires" datetime="2026-01-01 00:00:00"></time></article>`, Expired},
		{"not yet expired", `<article class="body"><time class="expires" datetime="2026-12-31 00:00:00"></time></article>`, Published},
		{"date-only future", `<article class="body"><time class="published" datetime="2026-07-01"></time></article>`, Scheduled},
		{"draft beats schedule", `<article class="body draft"><time class="published" datetime="2026-12-31 00:00:00"></time></article>`, Draft},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := PublishStatus(doc(t, tt.body), now).State; got != tt.want {
				t.Errorf("PublishStatus = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDate(t *testing.T) {
	if _, err := ParseDate("2026-06-04 10:00:00"); err != nil {
		t.Errorf("DateTime: %v", err)
	}
	if _, err := ParseDate("2026-06-04"); err != nil {
		t.Errorf("DateOnly: %v", err)
	}
	if _, err := ParseDate("nonsense"); err == nil {
		t.Error("expected error for nonsense")
	}
}
