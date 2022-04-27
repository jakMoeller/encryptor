package main

// #include <stdio.h>
// #include <stdlib.h>
// #include <unistd.h>
import "C"

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//export encrypt
func encrypt(k, v *C.char) *C.char {
	key, value := decodedString(k), decodedString(v)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, value, nil)
	return encodedString([]byte(fmt.Sprintf("%x", ciphertext)))
}

//export decrypt
func decrypt(k, v *C.char) *C.char {
	key, value := decodedString(k), decodedString(v)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := value[:nonceSize], value[nonceSize:]
	plaintext, decodeErr := aesGCM.Open(nil, nonce, ciphertext, nil)

	if decodeErr != nil {
		panic(err.Error())
	}

	return encodedString(plaintext)
}

func decodedString(in *C.char) []byte {
	out, err := hex.DecodeString(C.GoString(in))
	if err != nil {
		panic(err.Error())
	}
	return out
}

func encodedString(in []byte) *C.char {
	return C.CString(hex.EncodeToString(in))
}

func main() {
	// This is necessary for the compiler.
	// You can add something that will be executed when engaging your library to the interpreter.
}
