package main

import (
	"strings"
	"testing"
)

func TestOrbitCount(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
`
	expected := struct {
		direct, indirect int
	}{11, 31}

	uom, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Parsing error")
	}
	direct, indirect := uom.OrbitCount()
	if direct != expected.direct {
		t.Errorf("got %v direct orbits; expected %v", direct, expected.direct)
	}
	if indirect != expected.indirect {
		t.Errorf("got %v inddirect orbits; expected %v", indirect, expected.indirect)
	}
}

func TestOrbitalTransfers(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN
`
	expected := 4

	uom, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Parsing error: %v", err)
	}
	you, ok := uom["YOU"]
	if !ok {
		t.Fatalf("Parsing error: YOU can not be found")
	}
	san, ok := uom["SAN"]
	if !ok {
		t.Fatalf("Parsing error: SAN can not be found")
	}
	if n := you.OrbitalTransfers(san); n != expected {
		t.Errorf("got %v orbital transfers; expected %v", n, expected)
	}
}
