package gcss

import (
	"io"
	"testing"
)

func Test_variable_WriteTo(t *testing.T) {
	ln := newLine(1, "$test-name: test-value")

	v, err := newVariable(ln, nil)

	if err != nil {
		t.Errorf("error occurred [error: %v]", err.Error())
		return
	}

	v.WriteTo(io.Discard)
}

func Test_variable_WriteTo_fromFile(t *testing.T) {
	_, err := CompileFile("testdata/0013.gcss")

	if err != nil {
		t.Errorf("error occurred [error: %v]", err.Error())
		return
	}
}

func Test_variableNV_prefixDollarErr(t *testing.T) {
	ln := newLine(1, "test-name: test-value")

	_, _, err := variableNV(ln)

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected := "variable must start with \"$\" [line: 1]"; err.Error() != expected {
		t.Errorf("err should be %q [actual: %q]", expected, err.Error())
		return
	}
}

func Test_variableNV_lenErr(t *testing.T) {
	ln := newLine(1, "$test-name:test-value")

	_, _, err := variableNV(ln)

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected := "variable's name and value should be divided by a space [line: 1]"; err.Error() != expected {
		t.Errorf("err should be %q [actual: %q]", expected, err.Error())
		return
	}
}

func Test_variableNV_nameSuffixErr(t *testing.T) {
	ln := newLine(1, "$test-name test-value")

	_, _, err := variableNV(ln)

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected := "variable's name should end with a colon [line: 1]"; err.Error() != expected {
		t.Errorf("err should be %q [actual: %q]", expected, err.Error())
		return
	}
}

func Test_variableNV(t *testing.T) {
	ln := newLine(1, "$test-name: test-value")

	_, _, err := variableNV(ln)

	if err != nil {
		t.Errorf("error occurred [error: %v]", err.Error())
		return
	}
}

func Test_newVariable_err(t *testing.T) {
	ln := newLine(1, "test-name: test-value")

	_, err := newVariable(ln, nil)

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected := "variable must start with \"$\" [line: 1]"; err.Error() != expected {
		t.Errorf("err should be %q [actual: %q]", expected, err.Error())
		return
	}
}

func Test_newVariable_suffixErr(t *testing.T) {
	ln := newLine(1, "$test-name: test-value;")

	_, err := newVariable(ln, nil)

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected := "variable must not end with \";\""; err.Error() != expected {
		t.Errorf("err should be %q [actual: %q]", expected, err.Error())
		return
	}
}

func Test_newVariable(t *testing.T) {
	ln := newLine(1, "$test-name: test-value")

	_, err := newVariable(ln, nil)

	if err != nil {
		t.Errorf("error occurred [error: %v]", err.Error())
		return
	}
}
