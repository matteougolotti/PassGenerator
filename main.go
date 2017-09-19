package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	var pass string
	_, err := fmt.Scanln(&pass)
	if err != nil {
		panic(err)
	}

	hasher := sha1.New()
	hasher.Write([]byte(pass))
	sha := hasher.Sum(nil)

	res := hex.EncodeToString(sha)
	fmt.Println(res)
}
