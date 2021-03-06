// written by "xeuus" aka aryan

package main

import (
	"github.com/xeuus/gocry/pkg"
	"fmt"
)

func main() {

	// this method is nothing but a shuffled version of base64 encoding,
	// but also, works in Javascript or Typescript or any other language,
	// implementation is easy and needs not effort.
	// make sure you share this cipher with other client.
	cipher := cry.Base64Cipher()
	fmt.Println("Cipher: ", string(cipher))
	fmt.Println(cry.Base64Scrabmle(cipher))

	// we have a message and we want to send it for someone!
	myMessage := "Hello World!"
	fmt.Println("Message: ", myMessage)


	// for performance, set cipher once
	cry.Base64Set(cipher)


	// first encrypt it, this method takes an []byte array
	// so message could be a blob of data, whatever
	enc := cry.Base64Encode([]byte(myMessage))

	fmt.Println("Encoded: ", string(enc))
	// assume we have received a message from other client,
	// which we know it's encrypted using the same cipher,
	// we just gonna decrypt it
	dec := cry.Base64Decode(enc)

	// here we get our message back.
	fmt.Println("Decoded: ", string(dec))

	// let's go and run it, go run $GOPATH/src/github.com/xeuus/gocry/examples/b64c.go
}
