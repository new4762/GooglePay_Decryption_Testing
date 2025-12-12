# 🔐 Google Pay Payload Decryption & Formatter

This project provides tools for working with Google Pay encrypted payloads:
- Decrypt payloads using Google's ECv2 protocol in Java.
- Format payloads into Omise-style JSON using a Go helper.

---

## 📁 Project Structure

```
.
├── src/
│   └── main/
│       └── java/
│           └── com/
│               └── example/
│                   └── DecryptApp.java    # Java source code following Maven conventions
├── google_pay_json_helper.go        # Go helper to format Google Pay JSON payload
├── go.mod                          # Go module definition file
├── go.sum                          # Go dependency checksum file
├── pom.xml                         # Maven build configuration file
├── Makefile                        # Makefile to simplify build and run commands
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
  "public_key": "googletest",
  "data": "{...original payload JSON...}"
}
```

---

## 🛠 Development Notes

- Decryption uses [Tink](https://github.com/google/tink) (Google's crypto library).
- The Go formatter pulls the payload from the `GOOGLE_PAY_PAYLOAD` env variable.
