package controller

import validation "github.com/go-ozzo/ozzo-validation"

type PingRequest struct {
	Message string `json:"message"`
}

type PingResponse struct {
	Message string `json:"message"`
}

func (p PingRequest) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Message, validation.Required, validation.NilOrNotEmpty))
}
