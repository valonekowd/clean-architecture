package playground

import (
	"context"

	"github.com/go-playground/validator"

	"github.com/valonekowd/clean-architecture/domain/entity"
	"github.com/valonekowd/clean-architecture/infrastructure/log"
	_validator "github.com/valonekowd/clean-architecture/infrastructure/validator"
	"github.com/valonekowd/clean-architecture/util/helper"
)

type playgroundValidator struct {
	logger    log.Logger
	validator *validator.Validate
}

func NewValidator(tagName string, logger log.Logger) _validator.Validator {
	v := &playgroundValidator{
		logger:    logger,
		validator: validator.New(),
	}

	v.setup(tagName)

	return v
}

func (p *playgroundValidator) setup(tagName string) {
	p.validator.SetTagName(tagName)

	// custom validation
	p.validator.RegisterValidation("is-valid-transaction-type", func(fl validator.FieldLevel) bool {
		return helper.StringInSlice(fl.Field().String(), entity.SupportedTransactionTypes)
	})

	p.validator.RegisterValidation("is-valid-bank", func(fl validator.FieldLevel) bool {
		return helper.StringInSlice(fl.Field().String(), entity.SupportedBanks)
	})

	p.validator.RegisterValidation("is-valid-gender", func(fl validator.FieldLevel) bool {
		return helper.StringInSlice(fl.Field().String(), entity.SupportedGenders)
	})
}

func (p *playgroundValidator) Struct(ctx context.Context, s interface{}) error {
	return p.validator.StructCtx(ctx, s)
}
