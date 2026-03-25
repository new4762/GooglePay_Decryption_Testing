package googlepay

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"strings"
)

// PublicKeyFromBase64PKCS8 derives a base64-encoded public key in uncompressed
// point format from a base64-encoded PKCS#8 EC private key.
func PublicKeyFromBase64PKCS8(base64PKCS8 string) (string, error) {
	privateKeyDER, err := base64.StdEncoding.DecodeString(strings.TrimSpace(base64PKCS8))
	if err != nil {
		return "", fmt.Errorf("decode base64 private key: %w", err)
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyDER)
	if err != nil {
		return "", fmt.Errorf("parse PKCS#8 private key: %w", err)
	}

	ecPrivateKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("private key is %T, expected *ecdsa.PrivateKey", privateKey)
	}

	publicKeyBytes, err := ecPrivateKey.PublicKey.Bytes()
	if err != nil {
		return "", fmt.Errorf("encode public key: %w", err)
	}

	return base64.StdEncoding.EncodeToString(publicKeyBytes), nil
}
