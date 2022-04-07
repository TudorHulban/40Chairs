package http

import "github.com/gofiber/fiber/v2"

const (
	urlAnnounce = "myidis"
	urlPing     = "ping"
	urlRanges   = "ranges"
)

type server struct {
	engine *fiber.App
}
