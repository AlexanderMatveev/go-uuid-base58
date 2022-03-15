package uuid58

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

func FromBase58(s string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(s))
}
