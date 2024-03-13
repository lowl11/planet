package errors

import "encoding/json"

type output struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Context map[string]any `json:"context,omitempty"`
}

func (e *planetError) Output() []byte {
	var innerErrorMessage string
	if e.inner != nil {
		innerErrorMessage = ": " + e.inner.Error()
	}

	errorInBytes, err := json.Marshal(output{
		Message: e.message + innerErrorMessage,
		Code:    e.httpCode,
		Context: e.ctx,
	})
	if err != nil {
		return nil
	}

	return errorInBytes
}
