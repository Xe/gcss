package gcss

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func Test_e2e(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 13; i++ {
		idx := strconv.Itoa(i)

		gcssPath := "testdata/e2e/actual/" + strings.Repeat("0", 4-len(idx)) + idx + ".gcss"

		wg.Add(1)

		go func() {
			defer wg.Done()

			actualCSSPath, err := CompileFile(gcssPath)

			if err != nil {
				t.Errorf("error occurred [error: %q]", err.Error())
				return
			}

			expectedCSSPath := strings.Replace(actualCSSPath, "actual", "expected", -1)

			actualB, err := os.ReadFile(actualCSSPath)

			if err != nil {
				t.Errorf("error occurred [error: %q]", err.Error())
				return
			}

			expectedB, err := os.ReadFile(expectedCSSPath)

			if err != nil {
				t.Errorf("error occurred [error: %q]", err.Error())
				return
			}

			if strings.TrimSpace(string(actualB)) != strings.TrimSpace(string(expectedB)) {
				t.Errorf("actual result does not match the expected result [path: %q]", gcssPath)
				return
			}
		}()

		wg.Wait()
	}
}
