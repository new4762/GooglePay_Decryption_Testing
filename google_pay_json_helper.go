package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Failed to load .env file")
	}

	// Read environment variables
	publicKey := "googletest"
	raw := os.Getenv("GOOGLE_PAY_PAYLOAD")

	if publicKey == "" || raw == "" {
		log.Fatal("❌ Missing GOOGLE_PAY_MERCHANT_ID or GOOGLE_PAY_PAYLOAD in environment")
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
