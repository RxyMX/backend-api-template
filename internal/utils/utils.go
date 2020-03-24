package utils

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
)

func UnmarshalAndValidate(data []byte, v validation.Validatable) error {
	err := json.Unmarshal(data, v)

	if err == nil {
		err = v.Validate()
	}

	return err
}
