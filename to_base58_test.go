package id

import (
	"github.com/google/uuid"
	"testing"
)

var toBase58Correct = []struct {
	in  uuid.UUID
	out string
}{
	{uuid.Must(uuid.Parse("4f3beea4-d4a4-4842-bc96-57e818879f03")), "AnUw9CFxaSCF81BJrbbfiJ"},
}

func TestToBase58(t *testing.T) {
	for x, test := range toBase58Correct {
		u, err := ToBase58(test.in)
		if err != nil {
			t.Errorf("ToBase58 test #%d failed: got error: %v", x, err)
			continue
		}
		if u != test.out {
			t.Errorf("ToBase58 failed #%d: got: %s want: %s", x, u, test.out)
			continue
		}
	}
}
