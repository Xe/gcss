package gcss

import (
	"io"
	"testing"
)

func Test_comment_WriteTo(t *testing.T) {
	ln := newLine(1, "// test")

	c := newComment(ln, nil)

	if _, err := c.WriteTo(io.Discard); err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}
}

func Test_newComment(t *testing.T) {
	ln := newLine(1, "// test")

	c := newComment(ln, nil)

	if c == nil || c.ln != ln {
		t.Error("c is invalid")
		return
	}
}
