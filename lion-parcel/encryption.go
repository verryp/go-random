package main

import (
	"crypto/aes"
	"crypto/cipher"
	cRand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func main() {
	apiKeyPlain := GenerateApiKey(32)
	encryptionKey := "cee2e0f11be53a9dc1c400f45b9b113d"

	fmt.Println("apiKeyPlain", apiKeyPlain)
	secret, _ := Encrypt(encryptionKey, apiKeyPlain)
	fmt.Println("secret", secret)

	decryptedApiKey, _ := Decrypt(encryptionKey, secret)
	fmt.Println("decryptedApiKey", decryptedApiKey)
}

func GenerateApiKey(n int) string {
	apiKey := randomString(n)
	return apiKey
}

// randomString ...
func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Encrypt string to base64 crypto using AES
func Encrypt(encryptionKey, text string) (string, error) {

	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the cipherText.
	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(cRand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// Decrypt from base64 to decrypted string
func Decrypt(encryptionKey, cryptoText string) (string, error) {

	cipherText, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the cipherText.
	if len(cipherText) < aes.BlockSize {
		err = fmt.Errorf("ciphertext too short")
		return "", err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	return fmt.Sprintf("%s", cipherText), nil
}
