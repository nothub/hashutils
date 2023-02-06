package format

import (
	"testing"
)

func TestPHC(t *testing.T) {
	str := "$argon2id$v=19$m=65536,t=2,p=1$gZiV/M1gPc22ElAH/Jh1Hw$CWOrkoo7oJBQ/iyh7uJ0LO2aLEfrHwTWllSAxT0zRno"
	phc, err := ParsePHC(str)
	if err != nil {
		t.Fatal(err.Error())
	}
	if phc.Serialize() != str {
		t.Fatalf("wrong result: %s", phc.Serialize())
	}
}
