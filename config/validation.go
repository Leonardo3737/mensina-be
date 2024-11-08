package config

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Validate = validator.New()
	tranl    ut.Translator
)
