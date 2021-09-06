package rsakey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// 生成, 拿到密钥(在KeyPath中)
func GenerateRsaKey(keySize int) (string, string) {
	/***************** PRIVATE **********************/
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derText,
	}
	privKey := string(pem.EncodeToMemory(block))
	/***************** PUBLIC **********************/
	publicKey := privateKey.PublicKey
	derpText, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derpText,
	}
	pubKey := string(pem.EncodeToMemory(block))
	return privKey, pubKey
}

// RSADecrypt : decode
func RSADecrypt(cipherText string, privKeyS string) (string, error) {
	block, _ := pem.Decode([]byte(privKeyS))
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return "", err
	}
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, []byte(cipherText))
	if err != nil {
		return "", err
	}
	return string(plainText), err
}
