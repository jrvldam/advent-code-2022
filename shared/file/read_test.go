package file

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestGetFileContents(t *testing.T) {
  fs := fstest.MapFS{
    "input": { Data: []byte("input-contents") },
  }

  fileReader := NewReader()
  got, err := fileReader.GetContents(fs, "input")

  if err != nil {
    t.Fatalf("Reading file %v", err)
  }

  want := []string{"input-contents"}

  if !reflect.DeepEqual(got, want) {
    t.Errorf("want %+v, got %+v", want, got)
  }
}
