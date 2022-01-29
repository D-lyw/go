package main

import (
	"fmt"
	"github.com/D-lyw/go/hmac"
)

func main() {

	md5Result := hmac.Md5("ABC")
	fmt.Println(md5Result)

	hmacResult := hmac.Hmac("key", "ABC")
	fmt.Println(hmacResult)

	sha1Result := hmac.Sha1("ABC")
	fmt.Println(sha1Result)
}
