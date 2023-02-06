package b64

import "testing"

func TestB64(t *testing.T) {
	test(t, "foo")
	test(t, "äöü")
	test(t, "123")
	test(t, ":^)")
	test(t, "   ")
	test(t, "🍯🐻🤓")
	test(t, "《孟子·离娄上》")
	test(t, "\x00\x00\x00")
}

func test(t *testing.T, str string) {
	enc := Encode([]byte(str))
	dec, err := Decode(enc)
	if err != nil {
		t.Fatal(err.Error())
	}
	if str != string(dec) {
		t.Fatal("original != decoded")
	}
}
