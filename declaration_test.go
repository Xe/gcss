package gcss

import (
	"io"
	"testing"
)

func Test_declaration_WriteTo(t *testing.T) {
	ln := newLine(1, "font-size: 12px")

	dec, err := newDeclaration(ln, nil)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}

	_, err = dec.WriteTo(io.Discard)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}
}

func Test_newDeclaration_semicolonSuffixErr(t *testing.T) {
	ln := newLine(1, "color: blue;")

	_, err := newDeclaration(ln, nil)

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected := "declaration must not end with \";\" [line: 1]"; expected != err.Error() {
		t.Errorf("err should be %q [actual: %q]", expected, err.Error())
		return
	}
}

func Test_newDeclaration(t *testing.T) {
	ln := newLine(1, "html")

	_, err := newDeclaration(ln, nil)

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected := "declaration's property and value should be divided by a space [line: 1]"; expected != err.Error() {
		t.Errorf("err should be %q [actual: %q]", expected, err.Error())
		return
	}
}
