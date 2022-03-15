package uuid58

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

func ToBase58(u uuid.UUID) (string, error) {
	b, err := u.MarshalBinary()
	if err != nil {
		return "", err
	}
	return base58.Encode(b), nil
}
