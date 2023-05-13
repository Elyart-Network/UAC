package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

func RsaGenKey() (private, public []byte, err error) {
	// Generate RSA Keys
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}

	// Validate Private Key
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	private = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey

	// Validate Public Key
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	public = pem.EncodeToMemory(block)
	return
}

func RsaSignWithSha256(data []byte, keyBytes []byte) ([]byte, error) {
	// Decode private key
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("private key error")
	}

	// ParsePKCS1PrivateKey parses an RSA private key in PKCS#1, ASN.1 DER form.
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Sign the hash
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature, nil
}

func RsaVerySignWithSha256(data, signData, keyBytes []byte) (bool, error) {
	// Decode public key
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return false, errors.New("public key error")
	}

	// ParsePKIXPublicKey parses a DER encoded public key.
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	// Verify the signature
	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RsaEncrypt(data, keyBytes []byte) ([]byte, error) {
	// Decode public key
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// ParsePKIXPublicKey parses a DER encoded public key.
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// type assert
	pub := pubInterface.(*rsa.PublicKey)

	// EncryptOAEP encrypts the given message with RSA-OAEP.
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func RsaDecrypt(ciphertext, keyBytes []byte) ([]byte, error) {
	// Decode private key
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("private key error")
	}

	// ParsePKCS1PrivateKey parses an RSA private key in PKCS#1, ASN.1 DER form.
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS#1 v1.5.
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return nil, err
	}
	return data, nil
}
