package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/lowl11/planet/errors"
	"github.com/lowl11/planet/log"
	"net/http"
)

type Context struct {
	inner *fiber.Ctx
}

func New(ctx *fiber.Ctx) *Context {
	return &Context{
		inner: ctx,
	}
}

func (ctx *Context) With(with func(ctx *fiber.Ctx)) {
	with(ctx.inner)
}

func (ctx *Context) Header(key, value string) *Context {
	ctx.inner.Response().Header.Set(key, value)
	return ctx
}

func (ctx *Context) ContentType(contentType string) *Context {
	ctx.Header("Content-Type", contentType)
	return ctx
}

func (ctx *Context) Authorization(token string) *Context {
	ctx.Header("Authorization", token)
	return ctx
}

func (ctx *Context) Status(status int) *Context {
	ctx.inner.Response().SetStatusCode(status)
	return ctx
}

func (ctx *Context) Ok(body ...any) error {
	ctx.ContentType("application/json")

	var requestBody any
	if len(body) > 0 {
		requestBody = body[0]
	}

	if requestBody == nil {
		bodyInBytes, _ := json.Marshal(justOK{
			Status: "OK",
		})
		return ctx.inner.Send(bodyInBytes)
	}

	bodyInBytes, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	return ctx.Bytes(bodyInBytes)
}

func (ctx *Context) Created(id ...any) error {
	ctx.Status(http.StatusCreated)

	if len(id) > 0 {
		return ctx.Ok(createdWithID{
			ID: id[0],
		})
	}

	return ctx.Bytes(nil)
}

func (ctx *Context) Error(err error) error {
	ctx.ContentType("application/json")

	code := ctx.inner.Response().StatusCode()

	var planetErr errors.Error
	defer func() {
		log.Error(planetErr)
	}()

	planetErr, ok := err.(errors.Error)
	if ok {
		return ctx.
			Status(planetErr.HttpCode()).inner.
			Send(planetErr.Output())
	}

	if planetErr == nil {
		planetErr = errors.
			New("Untyped error").
			SetHTTP(code).
			SetError(err)
	}

	return ctx.
		Status(code).inner.
		Send(planetErr.Output())
}

func (ctx *Context) Bytes(body []byte) error {
	return ctx.inner.Send(body)
}
