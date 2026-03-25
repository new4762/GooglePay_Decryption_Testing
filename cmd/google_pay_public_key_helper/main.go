package main

import (
	"fmt"
	"log"
	"os"

	"example/internal/googlepay"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Failed to load .env file")
	}

	privateKey := os.Getenv("GOOGLE_PAY_PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("❌ Missing GOOGLE_PAY_PRIVATE_KEY in environment")
	}

	publicKey, err := googlepay.PublicKeyFromBase64PKCS8(privateKey)
	if err != nil {
		log.Fatalf("❌ Failed to extract public key: %v", err)
	}

	fmt.Println(publicKey)
}
