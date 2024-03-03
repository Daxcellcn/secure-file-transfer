package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func encryptFile(key []byte, inputFile string, outputFile string) error {
	plainText, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return ioutil.WriteFile(outputFile, cipherText, 0644)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: secure-file-transfer <input-file> <output-file>")
		os.Exit(1)
	}

	// Get the encryption key (you can modify this to suit your needs)
	key := sha256.Sum256([]byte("thisisaverysecurekey"))

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := encryptFile(key[:], inputFile, outputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File transfer completed successfully!")
}
