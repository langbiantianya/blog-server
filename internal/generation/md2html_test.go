package generation

import (
	"testing"
)

func TestMd2html(t *testing.T) {
	md := Md2html("test title", "# Hello, Goldmark!\n\nThis *is* a test document.")
	t.Log(string(md))
}
