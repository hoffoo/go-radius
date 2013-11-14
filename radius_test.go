package radius

import (
	"testing"
)

func TestServerStart(t *testing.T) {
	c := make(chan Packet)
	s, err := Start(&c, LocalAuth, "win")

	if err != nil {
		t.Fatal(err)
	}

	err = s.Close()

	if err != nil {
		t.Fatal(err)
	}
}
