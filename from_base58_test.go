package id

import (
	"github.com/google/uuid"
	"testing"
)

var fromBase58Correct = []struct {
	in  string
	out uuid.UUID
}{
	{"AnUw9CFxaSCF81BJrbbfiJ", uuid.Must(uuid.Parse("4f3beea4-d4a4-4842-bc96-57e818879f03"))},
}

var fromBase58Incorrect = []struct {
	in  string
	out uuid.UUID
}{
	{"INCORRECT_BASE58_STRING", uuid.Must(uuid.Parse("4f3beea4-d4a4-4842-bc96-57e818879f03"))},
}

func TestFromBase58(t *testing.T) {
	for x, test := range fromBase58Correct {
		u, err := FromBase58(test.in)
		if err != nil {
			t.Errorf("FromBase58 test #%d failed: got error: %v", x, err)
			continue
		}
		if u != test.out {
			t.Errorf("FromBase58 failed #%d: got: %s want: %s", x, u, test.out)
			continue
		}
	}
	for x, test := range fromBase58Incorrect {
		_, err := FromBase58(test.in)
		if err == nil {
			t.Errorf("FromBase58 test #%d failed: must be error", x)
			continue
		}
	}
}
