package gcss

import (
	"errors"
	"os"
	"strings"
	"testing"
)

var errErrReader = errors.New("errReader error")

type errReader struct{}

func (r *errReader) Read(p []byte) (int, error) {
	return 0, errErrReader
}

func TestCompile_readAllErr(t *testing.T) {
	if _, err := Compile(os.Stdout, &errReader{}); err != errErrReader {
		t.Errorf("error should be %+v [actual: %+v]", errErrReader, err)
		return
	}
}

func TestCompile_compileBytesErr(t *testing.T) {
	r, err := os.Open("testdata/0015.gcss")

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}

	_, err = Compile(os.Stdout, r)

	if expected, actual := "indent is invalid [line: 5]", err.Error(); actual != expected {
		t.Errorf("error should be %+q [actual: %+q]", expected, actual)
		return
	}
}

func TestCompile(t *testing.T) {
	r, err := os.Open("testdata/0016.gcss")

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}

	if _, err := Compile(os.Stdout, r); err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}
}

func TestCompileFile_readFileErr(t *testing.T) {
	_, err := CompileFile("not_exist_file")

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected, actual := "open not_exist_file: ", err.Error(); !strings.HasPrefix(actual, expected) || !os.IsNotExist(err) {
		t.Errorf("err should be %q [actual: %q]", expected, actual)
		return
	}
}

func TestCompileFile_compileStringErr(t *testing.T) {
	_, err := CompileFile("testdata/0004.gcss")

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected, actual := "indent is invalid [line: 5]", err.Error(); expected != actual {
		t.Errorf("err should be %q [actual: %q]", expected, actual)
		return
	}
}

func TestCompileFile_writeErr(t *testing.T) {
	cssFileBack := cssFilePath

	cssFilePath = func(_ string) string {
		return "not_exist_dir/not_exist_file"
	}

	_, err := CompileFile("testdata/0003.gcss")

	if err == nil {
		t.Error("error should be occurred")
		return
	}

	if expected, actual := "open not_exist_dir/not_exist_file: ", err.Error(); !strings.HasPrefix(actual, expected) || !os.IsNotExist(err) {
		t.Errorf("err should be %q [actual: %q]", expected, actual)
		return
	}

	cssFilePath = cssFileBack
}

func TestCompileFile(t *testing.T) {
	path, err := CompileFile("testdata/0003.gcss")

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}

	if expected := "testdata/0003.css"; expected != path {
		t.Errorf("path should be %q [actual: %q]", expected, path)
		return
	}
}

func TestCompileFile_pattern2(t *testing.T) {
	gcssPath := "testdata/0007.gcss"

	path, err := CompileFile(gcssPath)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
		return
	}

	if expected := cssFilePath(path); expected != path {
		t.Errorf("path should be %q [actual: %q]", expected, path)
		return
	}
}

func Test_complieBytes(t *testing.T) {
	compileBytes([]byte(""))
}

func TestPath(t *testing.T) {
	path := "/test"

	if expected, actual := path+extGCSS, Path(path+extCSS); actual != expected {
		t.Errorf("returned value should be %q [actual: %q]", expected, actual)
		return
	}
}
