package internal

import (
	"errors"
	"fmt"
)

func ErrFieldNotExist(fieldName string) error {
	return errors.New(fmt.Sprintf("orm: field %s not exists", fieldName))
}
