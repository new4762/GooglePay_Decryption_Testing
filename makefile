# Load environment variables from .env file
include .env
export

# Default target
.PHONY: all
all: help

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make payload     Format Google Pay payload to Omise JSON format"
	@echo "  make decrypt     Run the Java decryption app"

# Run Go JSON formatter
.PHONY: payload
payload:
	go run google_pay_json_helper.go

# Compile and run Java decryption app
.PHONY: decrypt
decrypt:
	mvn compile
	mvn exec:java -Dexec.mainClass="com.example.DecryptApp"
