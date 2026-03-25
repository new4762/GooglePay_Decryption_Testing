package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"example/internal/googlepay"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Failed to load .env file")
	}

	// Read environment variables
	privateKey := os.Getenv("GOOGLE_PAY_PRIVATE_KEY")
	raw := os.Getenv("GOOGLE_PAY_PAYLOAD")

	if privateKey == "" || raw == "" {
		log.Fatal("❌ Missing GOOGLE_PAY_PRIVATE_KEY or GOOGLE_PAY_PAYLOAD in environment")
	}

	publicKey, err := googlepay.PublicKeyFromBase64PKCS8(privateKey)
	if err != nil {
		log.Fatalf("❌ Failed to extract public key: %v", err)
	}

	type Wrapper struct {
		PublicKey string `json:"public_key"`
		Data      string `json:"data"`
	}

	// Marshal the wrapper
	out, err := json.Marshal(Wrapper{
		PublicKey: publicKey,
		Data:      raw,
	})
	if err != nil {
		log.Fatalf("❌ Failed to marshal: %v", err)
	}

	// Output the result
	fmt.Println(string(out))
}
