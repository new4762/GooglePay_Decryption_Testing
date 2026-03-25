package googlepay

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"testing"
)

func TestPublicKeyFromBase64PKCS8(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}

	privateKeyDER, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		t.Fatalf("marshal private key: %v", err)
	}

	got, err := PublicKeyFromBase64PKCS8(base64.StdEncoding.EncodeToString(privateKeyDER))
	if err != nil {
		t.Fatalf("extract public key: %v", err)
	}

	publicKeyBytes, err := privateKey.PublicKey.Bytes()
	if err != nil {
		t.Fatalf("encode public key: %v", err)
	}

	want := base64.StdEncoding.EncodeToString(publicKeyBytes)
	if got != want {
		t.Fatalf("unexpected public key: got %q want %q", got, want)
	}

	decoded, err := base64.StdEncoding.DecodeString(got)
	if err != nil {
		t.Fatalf("decode extracted public key: %v", err)
	}

	if len(decoded) != 65 {
		t.Fatalf("unexpected public key length: got %d want 65", len(decoded))
	}

	if decoded[0] != 0x04 {
		t.Fatalf("unexpected public key prefix: got 0x%02x want 0x04", decoded[0])
	}
}
