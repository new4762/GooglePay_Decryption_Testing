# 🔐 Google Pay Payload Decryption & Formatter

This project provides tools for working with Google Pay encrypted payloads:
- Decrypt payloads using Google's ECv2 protocol in Java.
- Format payloads into Omise-style JSON using a Go helper.
- Extract the base64 public key derived from your Google Pay private key.

---

## 📁 Project Structure

```
.
├── cmd/
│   └── google_pay_public_key_helper/
│       └── main.go                # Go helper to derive the public key from GOOGLE_PAY_PRIVATE_KEY
├── internal/
│   └── googlepay/
│       ├── keyutil.go             # Shared EC key conversion utilities
│       └── keyutil_test.go        # Tests for public-key derivation
├── src/
│   └── main/
│       └── java/
│           └── DecryptApp.java    # Java decryption entrypoint
├── google_pay_json_helper.go      # Go helper to format Google Pay JSON payload
├── go.mod                         # Go module definition file
├── go.sum                         # Go dependency checksum file
├── pom.xml                        # Maven build configuration file
├── Makefile                       # Makefile to simplify build and run commands
├── .env                           # Environment variables for configuration
└── README.md                      # Project overview and usage instructions
```

---

## 🔧 Setup

### 1. Install dependencies

```bash
# Java / Maven
brew install maven   # or use apt/choco depending on your OS

# Go (if not installed)
brew install go
```

---

### 2. Environment Variables

Create a `.env` file in the project root with the following:

```dotenv
GOOGLE_PAY_ENV=test # or prod
GOOGLE_PAY_MERCHANT_ID=gateway:your_gateway_name
GOOGLE_PAY_PRIVATE_KEY=YOUR_PKCS8_BASE64_KEY
GOOGLE_PAY_PAYLOAD={"signature":"...","protocolVersion":"ECv2",...}
```

`GOOGLE_PAY_PRIVATE_KEY` is expected to be a base64-encoded PKCS#8 EC private key. The Go helpers derive the matching public key from this value automatically in Google Pay's uncompressed point format.

---

## ▶️ Usage

### Decrypt payload (Java)

```bash
make decrypt
```

### Format payload as Omise JSON

```bash
make payload
```

Output:
```json
{
  "public_key": "<base64 uncompressed EC public key derived from GOOGLE_PAY_PRIVATE_KEY>",
  "data": "{...original payload JSON...}"
}
```

### Extract the derived public key

```bash
make public-key
```

Output:

```text
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE...
```

---

## 🛠 Development Notes

- Decryption uses [Tink](https://github.com/google/tink) (Google's crypto library).
- The Go helpers expect `GOOGLE_PAY_PRIVATE_KEY` to be a base64-encoded PKCS#8 EC private key.
- The Go formatter pulls the payload from the `GOOGLE_PAY_PAYLOAD` env variable and derives `public_key` from the private key in uncompressed point format.
