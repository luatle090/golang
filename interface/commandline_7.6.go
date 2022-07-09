package main

import (
	"crypto/sha512"
	"encoding/base64"
	"flag"
	"fmt"
)

type TextEncrypt string

type encryptFlag struct{ TextEncrypt }

func encryptSha512(plainText string) string {
	var sha512Hash = sha512.New()
	sha512Hash.Write([]byte(plainText))

	// Get the SHA-512 hashed password. Return []byte
	hashedPasswordBytes := sha512Hash.Sum(nil)

	// convert []byte to base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash
}

func (txt TextEncrypt) String() string {
	return fmt.Sprintf("after hash: %s", string(txt))
}

func (e *encryptFlag) Set(s string) error {
	var unit string
	var value string
	fmt.Sscanf(s, "%s%s", &value, &unit) // no error check needed
	switch unit {
	case "sha512":
		fmt.Print()
		e.TextEncrypt = TextEncrypt(encryptSha512(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)

}

func EncryptFlag(name string, value TextEncrypt, usage string) *TextEncrypt {
	e := encryptFlag{value}
	flag.CommandLine.Var(&e, name, usage)
	return &e.TextEncrypt
}

var encryptSha = EncryptFlag("sha512", TextEncrypt(encryptSha512("20")), "the sha512")

func main() {
	flag.Parse()
	fmt.Println(*encryptSha)
}
