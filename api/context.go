package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/lowl11/planet/errors"
	"github.com/lowl11/planet/log"
)

func Ok(ctx *fiber.Ctx, body ...any) error {
	var requestBody any
	if len(body) > 0 {
		requestBody = body[0]
	}

	setCtJSON(ctx)
	return Bytes(ctx, requestBody)
}

func Error(ctx *fiber.Ctx, err error) error {
	setCtJSON(ctx)
	defer log.Error(err)

	if planetErr, ok := err.(errors.Error); ok {
		return Bytes(ctx, planetErr.Output())
	}

	return Bytes(ctx, errors.
		New("Untyped error").
		SetError(err))
}

func Bytes(ctx *fiber.Ctx, body any) error {
	if body == nil {
		bodyInBytes, _ := json.Marshal(justOK{
			Status: "OK",
		})
		return ctx.Send(bodyInBytes)
	}

	bodyInBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return ctx.Send(bodyInBytes)
}
