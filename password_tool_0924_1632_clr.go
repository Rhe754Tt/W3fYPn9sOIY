// 代码生成时间: 2025-09-24 16:32:32
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "errors"
    "fmt"
    "log"
    "strings"
)

// PasswordTool provides functionality for password encryption and decryption
type PasswordTool struct {
    key []byte
}

// NewPasswordTool creates a new instance of PasswordTool with a given key
func NewPasswordTool(key string) (*PasswordTool, error) {
    if len(key) < 32 {
        return nil, errors.New("key must be at least 32 bytes")
    }
    return &PasswordTool{key: []byte(key)}, nil
}

// Encrypt encrypts the given plaintext password
func (pt *PasswordTool) Encrypt(plaintext string) (string, error) {
    block, err := aes.NewCipher(pt.key)
    if err != nil {
        return "", err
    }

    // PKCS7 padding
    plaintextBytes := PKCS7Padding([]byte(plaintext), aes.BlockSize)
    // CBC mode
    ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextBytes)

    // Base64 encoding for safe transport
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given ciphertext password
func (pt *PasswordTool) Decrypt(ciphertext string) (string, error) {
    decoded, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(pt.key)
    if err != nil {
        return "", err
    }

    if len(decoded) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }

    iv := decoded[:aes.BlockSize]
    ciphertext = decoded[aes.BlockSize:]
    stream := cipher.NewCFBDecrypter(block, iv)

    // XORKeyStream can work in-place if the two arguments are the same.
    stream.XORKeyStream(ciphertext, ciphertext)

    // Unpadding
    plaintextBytes := PKCS7UnPadding(ciphertext)
    return string(plaintextBytes), nil
}

// PKCS7Padding pads the plaintext with PKCS7 padding scheme
func PKCS7Padding(src []byte, blockSize int) []byte {
    padding := blockSize - len(src)%blockSize
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(src, padText...)
}

// PKCS7UnPadding removes the padding from the PKCS7 padded data
func PKCS7UnPadding(src []byte) []byte {
    length := len(src)
    unpadding := int(src[length-1])
    return src[:(length - unpadding)]
}

func main() {
    key := "your-very-very-very-very-very-very-very-very-very-very-long-secret-key"
    tool, err := NewPasswordTool(key)
    if err != nil {
        log.Fatal(err)
    }

    password := "mySecretPassword"
    encrypted, err := tool.Encrypt(password)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Encrypted: ", encrypted)

    decrypted, err := tool.Decrypt(encrypted)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Decrypted: ", decrypted)
}
