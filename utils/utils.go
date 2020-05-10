package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(origin string) string {
	hasher := md5.New()
	hasher.Write([]byte(origin))
	return hex.EncodeToString(hasher.Sum(nil))
}
