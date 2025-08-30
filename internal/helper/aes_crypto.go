package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

// func EncryptData(loan db.Loans) (db.Loans, error) {
// 	encryptedTexts := make(map[string]string)
// 	for key, text := range texts {
// 		encryptedText, err := encrypt(text)
// 		if err != nil {
// 			fmt.Printf("テキスト '%s' の暗号化エラー: %s\n", key, err)
// 			continue
// 		}
// 		encryptedTexts[key] = encryptedText
// 		fmt.Printf("テキスト '%s' が暗号化されました: %s\n", key, encryptedText)
// 	}
// 	return loan, nil
// }

type Crypto struct {
	Key string
}

func SetupCrypto(k string) Crypto {
	crypto := Crypto{
		Key: k,
	}
	return crypto
}

func (c Crypto) EncryptString(text string) (string, error) {
	fmt.Println(c.Key)
	if c.Key == "" {
		return "", errors.New("key is empty")
	}
	key := []byte(c.Key)

	plaintext := []byte(text)

	// ハッシュ関数を使ってテキストをハッシュ化し、初期化ベクトルとして使用
	hash := sha256.Sum256(key)
	iv := hash[:aes.BlockSize]

	// 暗号化アルゴリズムのブロックを作成
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 暗号化
	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func (c Crypto) DecryptString(ciphertext string) (string, error) {
	if c.Key == "" {
		return "", errors.New("key is empty")
	}

	key := []byte(c.Key)

	decoded, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// ハッシュ関数を使って初期化ベクトルを生成
	hash := sha256.Sum256(key)
	iv := hash[:aes.BlockSize]

	// 暗号化アルゴリズムのブロックを作成
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 復号化
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decoded, decoded)

	return string(decoded), nil
}
