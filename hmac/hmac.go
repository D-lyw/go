package hmac

import (
	"crypto/hmac"
	md52 "crypto/md5"
	sha12 "crypto/sha1"
	"encoding/hex"
)

func Md5(data string) string {
	md5 := md52.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(""))

	return hex.EncodeToString(md5Data)
}

func Hmac(key, data string) string {
	h := hmac.New(md52.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(data string) string {
	sha1 := sha12.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum(nil))
}
