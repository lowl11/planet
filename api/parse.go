package api

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gofiber/fiber/v2"
	"github.com/lowl11/planet/errors"
	"github.com/lowl11/planet/log"
	"github.com/lowl11/planet/pkg/types"
	"reflect"
	"regexp/syntax"
	"sync"
)

var (
	_validator = sync.Pool{
		New: func() any {
			validator, err := newValidator()
			if err != nil {
				log.Error("Create validator error: ", err)
				return nil
			}

			return validator
		},
	}
)

func Parse(ctx *fiber.Ctx, export any) error {
	contentType := types.ToString(ctx.Request().Header.Peek("Content-Type"))

	if reflect.ValueOf(export).Kind() != reflect.Ptr {
		return errors.New("Pointer required")
	}

	v := _validator.Get().(*validator)
	switch contentType {
	case "application/json":
		if err := json.Unmarshal(ctx.Body(), &export); err != nil {
			return errors.
				New("Parse request body for validation error").
				SetError(err)
		}

		if err := v.Struct(export); err != nil {
			return err
		}

		return nil
	case "application/xml":
		if err := xml.Unmarshal(ctx.Body(), &export); err != nil {
			return errors.
				New("Parse request body for validation error").
				SetError(err)
		}

		if err := v.Struct(export); err != nil {
			return err
		}

		return nil
	}

	return errors.
		New("Unknown content-type").
		AddContext("Content-Type", syntax.OpAnyCharNotNL)
}
