package exception

import (
	"errors"
	"fmt"
)

func Try(action func() error) (err error) {
	defer func() {
		panicError := Catch(recover())
		if panicError != nil {
			err = panicError
		}
	}()

	if err = action(); err != nil {
		return err
	}

	return nil
}

func Catch(panicError any) error {
	if panicError == nil {
		return nil
	}

	if err, ok := panicError.(error); ok {
		return err
	}

	if errText, ok := panicError.(string); ok {
		return errors.New(errText)
	}

	return fmt.Errorf("%s", panicError)
}
