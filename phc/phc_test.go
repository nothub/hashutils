package phc

import (
	"testing"
)

func TestPHC(t *testing.T) {
	phcStr := "$argon2id$v=19$m=65536,t=2,p=1$gZiV/M1gPc22ElAH/Jh1Hw$CWOrkoo7oJBQ/iyh7uJ0LO2aLEfrHwTWllSAxT0zRno"
	mcfStr := "$argon2id$CWOrkoo7oJBQ/iyh7uJ0LO2aLEfrHwTWllSAxT0zRno"

	phc, err := Parse(phcStr)
	if err != nil {
		t.Fatal(err.Error())
	}
	if phc.Serialize() != phcStr {
		t.Fatalf("wrong result: %s", phc.Serialize())
	}

	mcf, err := phc.ToMCF()
	if err != nil {
		t.Fatal(err.Error())
	}
	if mcf.Serialize() != mcfStr {
		t.Fatalf("wrong result: %s", phc.Serialize())
	}
}
