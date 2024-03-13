package api

import "github.com/gofiber/fiber/v2"

func SetHeader(ctx *fiber.Ctx, key, value string) *fiber.Ctx {
	ctx.Response().Header.Set(key, value)
	return ctx
}

func setCtJSON(ctx *fiber.Ctx) {
	SetHeader(ctx, "Content-Type", "application/json")
}
