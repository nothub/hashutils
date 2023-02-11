package phc

import (
	"testing"
)

func TestPHC(t *testing.T) {
	phcStr := "$argon2id$v=19$m=65536,t=2,p=1$gZiV/M1gPc22ElAH/Jh1Hw$CWOrkoo7oJBQ/iyh7uJ0LO2aLEfrHwTWllSAxT0zRno"
	phc, err := Parse(phcStr)
	if err != nil {
		t.Fatal(err.Error())
	}
	if phc.Serialize() != phcStr {
		t.Fatalf("\nexpected: %s\nactual:   %s", phcStr, phc.Serialize())
	}

	for _, p := range phc.Params() {
		if p.K == "" || p.V == "" {
			t.Fatal("empty param field")
		}
	}

	mcfStr := "$argon2id$CWOrkoo7oJBQ/iyh7uJ0LO2aLEfrHwTWllSAxT0zRno"
	mcf, err := phc.ToMCF()
	if err != nil {
		t.Fatal(err.Error())
	}
	if mcf.Serialize() != mcfStr {
		t.Fatalf("\nexpected: %s\nactual:   %s", mcfStr, mcf.Serialize())
	}
}
