package middleware

import (
	"be-client/util"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

type Source string

const (
	HeaderName            = "X-Csrf-Token"
	SourceCookie   Source = "cookie"
	SourceHeader   Source = "header"
	SourceURLQuery Source = "query"
)

func CorsFilter() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Language",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           86400,
	})
}

func CSRFFilter() fiber.Handler {
	return csrf.New(csrf.Config{
		KeyLookup:      fmt.Sprintf("%s:%s", SourceHeader, HeaderName),
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     30 * time.Minute,
		KeyGenerator:   util.UUIDFunc(),
	})
}
