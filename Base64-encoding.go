package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	data := "abc123!?$*&()'-=@~"

	p := fmt.Println
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	p(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	p(string(sDec))
	p("\n")
	
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	p(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	p(string(uDec))
}