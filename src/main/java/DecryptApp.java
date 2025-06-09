package com.example;

import com.google.crypto.tink.apps.paymentmethodtoken.GooglePaymentsPublicKeysManager;
import com.google.crypto.tink.apps.paymentmethodtoken.PaymentMethodTokenRecipient;
import com.google.crypto.tink.subtle.EllipticCurves;
import io.github.cdimascio.dotenv.Dotenv;

import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.interfaces.ECPrivateKey;
import java.util.Base64;

public class DecryptApp {
    public static void main(String[] args) throws Exception {
        System.out.println("🔐 Google Pay Decryption: ECv2");

        // Load environment config from .env file
        Dotenv dotenv = Dotenv.configure().filename(".env").ignoreIfMissing().load();

        String env = dotenv.get("GOOGLE_PAY_ENV").toLowerCase();
        String merchantId = dotenv.get("GOOGLE_PAY_MERCHANT_ID");
        String privateKey = dotenv.get("GOOGLE_PAY_PRIVATE_KEY");
        String payload = dotenv.get("GOOGLE_PAY_PAYLOAD");

        if (payload == null || payload.isBlank()) {
            throw new IllegalArgumentException("❌ Missing GOOGLE_PAY_PAYLOAD in .env file");
        }

        // Load private key (PKCS#8, Base64-encoded)
        String privateKeyBase64 = privateKey.trim();
        ECPrivateKey ecPrivateKey = parseECPrivateKey(privateKeyBase64);

        // Choose correct Google public key manager
        GooglePaymentsPublicKeysManager keysManager = switch (env) {
            case "prod", "production" -> {
                System.out.println("🔧 Using GooglePaymentsPublicKeysManager: PRODUCTION");
                yield GooglePaymentsPublicKeysManager.INSTANCE_PRODUCTION;
            }
            default -> {
                System.out.println("🔧 Using GooglePaymentsPublicKeysManager: TEST");
                yield GooglePaymentsPublicKeysManager.INSTANCE_TEST;
            }
        };

        // Decrypt the payload
        String decryptedMessage = new PaymentMethodTokenRecipient.Builder()
                .fetchSenderVerifyingKeysWith(keysManager)
                .recipientId(merchantId)
                .protocolVersion("ECv2")
                .addRecipientPrivateKey(ecPrivateKey)
                .build()
                .unseal(payload);

        System.out.println("✅ Decrypted message:\n" + decryptedMessage);
    }

    private static ECPrivateKey parseECPrivateKey(String base64Pkcs8) throws Exception {
        byte[] decoded = Base64.getDecoder().decode(base64Pkcs8);
        return EllipticCurves.getEcPrivateKey(decoded);
    }
}
