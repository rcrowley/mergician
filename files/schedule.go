package files

import (
	"strings"
	"time"

	"github.com/rcrowley/mergician/html"
	"golang.org/x/net/html/atom"
)

// State describes a document's editorial state relative to a reference time.
type State int

const (
	Published State = iota // build it
	Draft                  // marked class="draft"; never build
	Scheduled              // its published <time> is in the future
	Expired                // past its expires <time>
)

func (s State) String() string {
	switch s {
	case Draft:
		return "draft"
	case Scheduled:
		return "scheduled"
	case Expired:
		return "expired"
	default:
		return "published"
	}
}

// Status is the result of evaluating a document's editorial state, including
// the relevant publish or expiry instant (the zero time when not applicable).
type Status struct {
	State State
	When  time.Time
}

// ParseDate parses a datetime attribute the way feed does: a full timestamp
// (time.DateTime) first, then a date alone (time.DateOnly).
func ParseDate(s string) (time.Time, error) {
	if t, err := time.Parse(time.DateTime, s); err == nil {
		return t, nil
	}
	return time.Parse(time.DateOnly, s)
}

// PublishStatus reports a parsed document's editorial state relative to now:
//
//   - Draft if any element carries a "draft" class token (e.g.
//     <body class="body draft">), so it's never built;
//   - Scheduled if it has a <time class="published"> (or, failing that, the
//     <time class="feed"> feed already reads) whose instant is after now;
//   - Expired if it has a <time class="expires"> whose instant is at or before
//     now;
//   - Published otherwise, including when it carries no dates at all.
//
// now is supplied by the caller and never read from the clock here, so builds
// are deterministic and testable. Resolve it once, at the edge of your program.
func PublishStatus(n *html.Node, now time.Time) Status {
	if hasClass(n, "draft") {
		return Status{State: Draft}
	}

	published := findTime(n, "published")
	if published == nil {
		published = findTime(n, "feed") // the same date feed sorts and dates entries by
	}
	if published != nil {
		if when, err := ParseDate(html.Attr(published, "datetime")); err == nil && when.After(now) {
			return Status{State: Scheduled, When: when}
		}
	}

	if expires := findTime(n, "expires"); expires != nil {
		if when, err := ParseDate(html.Attr(expires, "datetime")); err == nil && !when.After(now) {
			return Status{State: Expired, When: when}
		}
	}

	return Status{State: Published}
}

// IsPublished reports whether a document should be built at the reference time.
func IsPublished(n *html.Node, now time.Time) bool {
	return PublishStatus(n, now).State == Published
}

// findTime finds the first <time class="..."> element with the exact class,
// matching how feed finds <time class="feed">.
func findTime(n *html.Node, class string) *html.Node {
	return html.Find(n, html.All(
		html.IsAtom(atom.Time),
		html.HasAttr("class", class),
	))
}

// hasClass reports whether any element in n carries the given class token,
// splitting the class attribute on whitespace so "body draft" matches "draft".
func hasClass(n *html.Node, token string) bool {
	return html.Find(n, func(node *html.Node) bool {
		for _, field := range strings.Fields(html.Attr(node, "class")) {
			if field == token {
				return true
			}
		}
		return false
	}) != nil
}
