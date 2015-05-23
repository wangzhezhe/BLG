package main

import (
	//b64 is the shortcut writing for the "encoding/base64"
	b64 "encoding/base64"
	"fmt"
)

func main() {

	fmt.Println("using stdencoding:")
	data := "abc123!?$*&()'-=@~"
	//加密
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	//解密
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	fmt.Println("using urlencoding:")
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	//actually the encode and decode process is really easy
	//there are some different using two kind of encoding ways
	//attention to  which way is needed in practical

}
