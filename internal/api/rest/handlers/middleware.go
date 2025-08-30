package handlers

import (
	"gifma-backend/pkg"
	"gifma-backend/token"

	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorizationHeader := ctx.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.Status(http.StatusUnauthorized).JSON(pkg.ErrorResponse(err))
			return nil
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New(" authorization header format")
			ctx.Status(http.StatusUnauthorized).JSON(pkg.ErrorResponse(err))
			return nil
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.Status(http.StatusUnauthorized).JSON(pkg.ErrorResponse(err))
			return nil
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.Status(http.StatusUnauthorized).JSON(pkg.ErrorResponse(err))
			return nil
		}

		ctx.Locals(authorizationPayloadKey, payload)
		ctx.Next()
		return nil
	}

}
