package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(in string) string {
	c := sha1.New()
	c.Write([]byte(in))
	return hex.EncodeToString(c.Sum(nil))
}
