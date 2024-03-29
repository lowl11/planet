package param

import (
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type Param struct {
	value string
}

func New(value string) *Param {
	return &Param{
		value: value,
	}
}

func (param Param) String() string {
	return param.value
}

func (param Param) Strings() []string {
	return strings.Split(param.value, ",")
}

func (param Param) Int() (int, error) {
	intValue, err := strconv.Atoi(param.value)
	if err != nil {
		return 0, err
	}

	return intValue, nil
}

func (param Param) MustInt() int {
	intValue, err := strconv.Atoi(param.value)
	if err != nil {
		return 0
	}

	return intValue
}

func (param Param) Bool() bool {
	return param.value == "true"
}

func (param Param) UUID() (uuid.UUID, error) {
	uuidValue, err := uuid.Parse(param.value)
	if err != nil {
		return uuid.UUID{}, err
	}

	return uuidValue, nil
}

func (param Param) MustUUID() uuid.UUID {
	uuidValue, err := uuid.Parse(param.value)
	if err != nil {
		return uuid.UUID{}
	}

	return uuidValue
}
