package generation

import (
	"testing"
)

func TestMd2html(t *testing.T) {
	md, err := Md2html([]byte("# Hello, Goldmark!\n\nThis *is* a test document."))
	if err != nil {
		t.Error(err)
	}
	t.Log(string(md))
}
