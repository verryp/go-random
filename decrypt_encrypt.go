package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type (
	Crypto interface {
		Encrypt(ctx context.Context, text string) (string, error)
		Decrypt(ctx context.Context, cryptoText string) (string, error)
	}

	crypto struct {
		encryptionKey string
	}
)

func NewCrypto(encryptionKey string) Crypto {
	return &crypto{
		encryptionKey: encryptionKey,
	}
}

// Encrypt string to base64 crypto using AES
func (c *crypto) Encrypt(ctx context.Context, text string) (string, error) {
	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(c.encryptionKey))
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the cipherText.
	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// Decrypt from base64 to decrypted string
func (c *crypto) Decrypt(ctx context.Context, cryptoText string) (string, error) {
	cipherText, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher([]byte(c.encryptionKey))
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the cipherText.
	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	return fmt.Sprintf("%s", cipherText), nil
}

func main() {
	fmt.Println("ngateee", ngateEncryptor())
	crypto := NewCrypto("cee2e0f75be53a9dc1c400f45b9b113d")

	//encrypted, _ := crypto.Encrypt(context.Background(), "lxgUsUIZO1UdnjtckADexMOTx0xUIAmT5KknS0k3v04LO4I3jXZH9jhbA28LQtojTPDUJOXsMEV0iVyujEkBY4Cdy48I6OXwbPGEUP0gAS0=")
	//decrypted, _ := crypto.Decrypt(context.Background(), encrypted)
	decryptedd, err := crypto.Decrypt(context.Background(), "lxgUsUIZO1UdnjtckADexMOTx0xUIAmT5KknS0k3v04LO4I3jXZH9jhbA28LQtojTPDUJOXsMEV0iVyujEkBY4Cdy48I6OXwbPGEUP0gAS0=")
	if err != nil {
		fmt.Println("err", err)
	}

	//fmt.Println("encrypted", encrypted)
	//fmt.Println(decrypted)
	fmt.Println(decryptedd)

	nc := NewCrypto("fb50063eab734ca5be314b0205899e7e")
	dd, err := nc.Encrypt(context.Background(), decryptedd)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("dd", dd)

	mas := NewCrypto("fb50063eab734ca5be314b0205899e7e")
	masDD, _ := mas.Decrypt(context.Background(), "UU2AJiGKYyDUBSaMo8ejIlnWR7hXNkPIK5JvSWaXIJL_iE6sqoY693TNP2FJm7e3y5WF5uG-yxV77QSfStsITTW4fNDHhmorDEQT-T8J3w4=")
	fmt.Println("masddd", masDD)

	mm := NewCrypto("cee2e0f75be53a9dc1c400f45b9b113d")
	ktl, _ := mm.Decrypt(context.Background(), "lxgUsUIZO1UdnjtckADexMOTx0xUIAmT5KknS0k3v04LO4I3jXZH9jhbA28LQtojTPDUJOXsMEV0iVyujEkBY4Cdy48I6OXwbPGEUP0gAS0=")
	fmt.Println("k", ktl)
}

func ngateEncryptor() string {
	c := NewCrypto("cee2e0f75be53a9dc1c400f45b9b113d")

	decr, _ := c.Decrypt(context.Background(), "lxgUsUIZO1UdnjtckADexMOTx0xUIAmT5KknS0k3v04LO4I3jXZH9jhbA28LQtojTPDUJOXsMEV0iVyujEkBY4Cdy48I6OXwbPGEUP0gAS0=")

	return decr
}
