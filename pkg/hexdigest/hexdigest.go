package hexdigest

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gophers0/users/pkg/errs"
)

// HexDigest generate hex string of rnd 32 byte value.
func HexDigest() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", errs.NewStack(err)
	}

	return hex.EncodeToString(bytes), nil
}
