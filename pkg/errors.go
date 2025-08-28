package pkg

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse はエラーレスポンスを生成する関数です。
func ErrorResponse(err error) fiber.Map {
	switch err.Error() {
	case "sql: no rows in result set":
		return fiber.Map{"error": "データがありません"}
	// case `pq: duplicate key value violates unique constraint \"loans_customer_id_key\"`:
	// 	return gin.H{"error": "重複登録です"}
	default:
		return fiber.Map{"error": err.Error()}
	}

}
