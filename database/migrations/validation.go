package migrations

import (
	"github.com/go-ozzo/ozzo-validation"
	"regexp"
)

func (p Product) ValidateProduct() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Match(regexp.MustCompile("^[A-Za-z0-9- ]+$"))),
		validation.Field(&p.CategoryID, validation.Required),
	)
}

func (c Category) ValidateCategory() error {
	return validation.ValidateStruct(&c,
		// aphanum with space
		validation.Field(&c.Name, validation.Required, validation.Match(regexp.MustCompile("^[A-Za-z0-9- ]+$"))),
	)
}
