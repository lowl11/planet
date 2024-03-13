package errors

import (
	"fmt"
	"github.com/lowl11/planet/pkg/types"
	"net/http"
	"strings"
)

type Error interface {
	error
	fmt.Stringer

	SetError(err error) Error
	AddContext(key string, value any) Error
	SetContext(ctx map[string]any) Error
	SetHTTP(code int) Error

	Context() map[string]any
	HttpCode() int

	Output() []byte
}

type planetError struct {
	inner    error
	message  string
	ctx      map[string]any
	httpCode int
}

func New(message string) Error {
	return &planetError{
		message:  message,
		ctx:      make(map[string]any),
		httpCode: http.StatusInternalServerError,
	}
}

func (e *planetError) String() string {
	errorBuilder := strings.Builder{}

	errorBuilder.WriteString(e.message)
	if e.inner != nil {
		errorBuilder.WriteString(": ")
		errorBuilder.WriteString(e.inner.Error())
	}

	if e.ctx != nil && len(e.ctx) > 0 {
		errorBuilder.WriteString(". ")
		errorBuilder.WriteString(types.ToString(e.ctx))
	}

	return errorBuilder.String()
}

func (e *planetError) Error() string {
	return e.String()
}

func (e *planetError) SetHTTP(code int) Error {
	e.httpCode = code
	return e
}

func (e *planetError) HttpCode() int {
	return e.httpCode
}

func (e *planetError) SetError(err error) Error {
	e.inner = err
	return e
}

func (e *planetError) AddContext(key string, value any) Error {
	e.ctx[key] = value
	return e
}

func (e *planetError) SetContext(ctx map[string]any) Error {
	if ctx == nil {
		return e
	}

	for key, value := range ctx {
		e.ctx[key] = value
	}

	return e
}

func (e *planetError) Context() map[string]any {
	return e.ctx
}
