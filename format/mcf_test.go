package format

import (
	"testing"
)

func TestMCF(t *testing.T) {
	str := "$foo$bar"
	mcf, err := ParseMCF(str)
	if err != nil {
		t.Fatal(err.Error())
	}
	if mcf.Serialize() != str {
		t.Fatalf("wrong result: %s", mcf.Serialize())
	}
}
