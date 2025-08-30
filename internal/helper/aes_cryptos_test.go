package helper

import (
	"testing"
)

func (c Crypto) TestEncryptDecrypt(t *testing.T) {

	// Setup crypto with test key
	crypto := SetupCrypto(c.Key)

	// Test data
	testData := "テスト岡野"

	// Encrypt
	encryptedData, err := crypto.EncryptString(testData)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}
	if encryptedData == "" {
		t.Fatal("Encrypted data is empty")
	}
}
