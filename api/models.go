package api

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type WelcomeResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	Env     string    `json:"env"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateShitpostPayload struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

func (csp CreateShitpostPayload) Validate() error {
	return validation.ValidateStruct(&csp,
		validation.Field(&csp.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&csp.Author, validation.Required, validation.Length(1, 255)),
		validation.Field(&csp.Content, validation.Required, validation.Length(1, 500)),
	)
}

type DeleteShitpostPayload struct {
	Passcode string `json:"passcode"`
}

func (dsp DeleteShitpostPayload) Validate() error {
	return validation.ValidateStruct(&dsp,
		validation.Field(&dsp.Passcode, validation.Required, validation.Length(8, 8)),
	)
}
