package chksum

import (
	"crypto/sha512"
	"github.com/nothub/hashutils/encoding"
	"strings"
	"testing"
)

func TestChecksum(t *testing.T) {
	testChksum(t, "hello", "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043", encoding.Hex)
	testChksum(t, "hello", "m3HSJL1i83hdltRq0+o9czGb+8KJDKra4t/3JRlnPKcjI8PZm6XBHXx6zG4UuMXaDEZjR1wuXDre9G9zvN7AQw", encoding.B64)
	testChksum(t, "foo!bar?baz ß_:)", "3b1ef87f6e39b543f6ff9c327285dfc3aa64e93f870fb6820b1b65a4a6825cf1b02857e0424ee7fd93db59d092981ec3dd8e32ea59fdb9e62a9b7a4502552879", encoding.Hex)
	testChksum(t, "foo!bar?baz ß_:)", "Ox74f245tUP2/5wycoXfw6pk6T+HD7aCCxtlpKaCXPGwKFfgQk7n/ZPbWdCSmB7D3Y4y6ln9ueYqm3pFAlUoeQ", encoding.B64)
}

func testChksum(t *testing.T, s string, h string, e encoding.Scheme) {
	r := strings.NewReader(s)
	c, err := Create(r, sha512.New(), e)
	if err != nil {
		t.Fatal(err.Error())
	}
	if c != h {
		t.Fatal("wrong result")
	}
}
