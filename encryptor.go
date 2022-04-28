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
	"sync"
)

var cryptLock = sync.Mutex{}

//export encrypt
func encrypt(k, v *C.char) *C.char {
	cryptLock.Lock()
	defer cryptLock.Unlock()
	key, err := hex.DecodeString(C.GoString(k))
	if err != nil {
		println(err.Error())
		return C.CString("DECODE_FAILURE")
	}

	value := []byte(C.GoString(v))
	block, err := aes.NewCipher(key)
	if err != nil {
		println(err.Error())
		return v
	}

	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		println(gcmErr.Error())
		return v
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		println(err.Error())
		return v
	}

	ciphertext := aesGCM.Seal(nonce, nonce, value, nil)

	println(fmt.Sprintf("Encryption Finished: %s => ENCRYPTED", value))

	return C.CString(string(ciphertext))
}

//export decrypt
func decrypt(k, v *C.char) *C.char {
	cryptLock.Lock()
	defer cryptLock.Unlock()
	key, err := hex.DecodeString(C.GoString(k))
	if err != nil {
		println(err.Error())
		return C.CString("DECODE_FAILURE")
	}

	value := []byte(C.GoString(v))
	block, err := aes.NewCipher(key)
	if err != nil {
		println(err.Error())
		return C.CString("DECODE_FAILURE")
	}

	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		println(gcmErr.Error())
		return C.CString("DECODE_FAILURE")
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := value[:nonceSize], value[nonceSize:]

	plaintext, decodeErr := aesGCM.Open(nil, nonce, ciphertext, nil)

	if decodeErr != nil {
		println(decodeErr.Error())
		return C.CString("DECODE_FAILURE")
	}

	println(fmt.Sprintf("Decryption Finished: ENCRYPTED => %s", plaintext))

	return C.CString(string(plaintext))
}

func main() {
	// This is necessary for the compiler.
	// You can add something that will be executed when engaging your library to the interpreter.
}
