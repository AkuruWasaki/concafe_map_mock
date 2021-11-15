// model helper methods
package modext

import (
	"time"

	"github.com/akuruwasaki/concafe_map_mock/models"
	"github.com/go-playground/validator"
	"github.com/volatiletech/null/v8"
)

type Shop struct {
	Name      string      `validate:"required"`
	Address   null.String `validate:"max=255"`
	Tel       null.String `validate:"max=45"`
	Content   null.String `validate:"max=10"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ValidateShop(s *models.Shop) error {
	var shop Shop
	// set param to validator
	shop.Name = s.Name
	shop.Address = s.Address
	shop.Tel = s.Tel
	shop.Content = s.Content

	validate := validator.New()
	return validate.Struct(shop)
}
