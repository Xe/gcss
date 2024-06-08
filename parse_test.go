package gcss

import (
	"os"
	"strings"
	"testing"
)

func Test_parse_topNewElementErr(t *testing.T) {
	data, err := os.ReadFile("./testdata/0011.gcss")

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}

	elemc, errc := parse(strings.Split(formatLF(string(data)), lf))

	select {
	case <-elemc:
		t.Error("error should be occurred")
		return
	case err := <-errc:
		if expected, actual := "selector must not end with \"{\" [line: 1]", err.Error(); actual != expected {
			t.Errorf("err should be %q [actual: %q]", expected, actual)
			return
		}
	}
}

func Test_parse_AppendChildrenNewElementErr(t *testing.T) {
	data, err := os.ReadFile("./testdata/0012.gcss")

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}

	elemc, errc := parse(strings.Split(formatLF(string(data)), lf))

	select {
	case <-elemc:
		t.Error("error should be occurred")
		return
	case err := <-errc:
		if expected, actual := "declaration must not end with \";\" [line: 2]", err.Error(); actual != expected {
			t.Errorf("err should be %q [actual: %q]", expected, actual)
			return
		}
	}
}

func Test_parse_appendChildrenErr(t *testing.T) {
	data, err := os.ReadFile("./testdata/0002.gcss")

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}

	elemc, errc := parse(strings.Split(formatLF(string(data)), lf))

	select {
	case <-elemc:
		t.Error("error should be occurred")
		return
	case err := <-errc:
		if expected, actual := "indent is invalid [line: 5]", err.Error(); actual != expected {
			t.Errorf("err should be %q [actual: %q]", expected, actual)
			return
		}
	}
}

func Test_parse(t *testing.T) {
	data, err := os.ReadFile("./testdata/0001.gcss")
	if err != nil {
		t.Errorf("error occurred [error: %s]", err.Error())
		return
	}

	elemc, errc := parse(strings.Split(formatLF(string(data)), lf))

	select {
	case <-elemc:
	case err := <-errc:
		t.Errorf("error occurred [error: %s]", err.Error())
		return
	}
}

func Test_formatLF(t *testing.T) {
	s := cr + crlf + lf + "a" + crlf + lf + cr + "b" + lf + cr + crlf
	expectedS := lf + lf + lf + "a" + lf + lf + lf + "b" + lf + lf + lf

	if formatLF(s) != expectedS {
		t.Errorf("formatLF(s) should be %s [actual: %s]", expectedS, formatLF(s))
		return
	}
}
