package runtime_tests

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/zuma206/rtml/runtime"
	"github.com/zuma206/rtml/stdlib"
)

const testdataPath = "testdata"

func TestRuntimeEndToEnd(t *testing.T) {
	entries, err := os.ReadDir(testdataPath)
	if err != nil {
		t.Errorf("failed to read %s directory", testdataPath)
		return
	}
	for _, entry := range entries {
		if err := runTestfile(t, path.Join(testdataPath, entry.Name())); err != nil {
			t.Errorf("error running test %s: %s", entry.Name(), err.Error())
		}
	}
}

func runTestfile(t *testing.T, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	parts := strings.Split(string(content), "\n==========\n")
	if len(parts) != 3 {
		return errors.New("testfile must have 3 parts")
	}
	code, expectedStream, expectedLog := []byte(parts[0]), (parts[1]), (parts[2])
	rt := runtime.New()
	stdlib.OpenStdlib(rt)
	var (
		stream bytes.Buffer
		log    bytes.Buffer
	)
	rt.Stream = &stream
	rt.Log = &log
	if err := rt.RunCode(bytes.NewBuffer(code)); err != nil {
		return err
	}
	compare(t, "stream", stream, expectedStream)
	compare(t, "log", log, expectedLog)
	return nil
}

func compare(t *testing.T, name string, actual bytes.Buffer, expected string) {
	actualString := actual.String()
	if actualString != expected {
		t.Logf("expected %s:\n%s\n", name, expected)
		t.Logf("actual %s:\n%s\n", name, actualString)
		t.Errorf("%s differed from expected", name)
	}
}
