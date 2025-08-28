package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomStringNum(digits int) string {
	numbers := []rune("0123456789")
	rand.Seed(time.Now().UnixNano())

	var result string
	for i := 0; i < digits; i++ {
		result += string(numbers[rand.Intn(len(numbers))])
	}

	formattedNumber := fmt.Sprintf("%s%s%s", result[:3], result[3:6], result[6:])
	return formattedNumber
}

func RandomMail() string {
	// mail := RandomString(6) + "@gmail.com"
	// return mail
	return fmt.Sprintf("%semail.com", RandomString(6))
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomStatus() string {
	currencies := []string{"current", "pending", "wakai"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomCollectionStatus() string {
	currencies := []string{"unprocessing", "processing", "done", "noresponse", "noconnect", "success"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomDate(year int) time.Time {
	// 指定された年の開始日と終了日を設定
	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year, time.December, 31, 23, 59, 59, 999999999, time.UTC)

	// 開始日から終了日までの間でランダムな日付を生成
	rand.Seed(time.Now().UnixNano())
	days := int(endDate.Sub(startDate).Hours() / 24)
	randomDays := rand.Intn(days + 1)

	// ランダムな日付を返す
	return startDate.AddDate(0, 0, randomDays)
}

func RandomTransactionDate(year int, month time.Month) time.Time {
	// 指定された年の開始日と終了日を設定
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year, month, 28, 23, 59, 59, 999999999, time.UTC)

	// 開始日から終了日までの間でランダムな日付を生成
	rand.Seed(time.Now().UnixNano())
	days := int(endDate.Sub(startDate).Hours() / 24)
	randomDays := rand.Intn(days + 1)

	// ランダムな日付を返す
	return startDate.AddDate(0, 0, randomDays)
}
