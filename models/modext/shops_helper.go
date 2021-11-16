// model helper methods
package modext

import (
	"fmt"
	"time"

	"github.com/akuruwasaki/concafe_map_mock/models"
	"github.com/go-playground/validator"
)

type Shop struct {
	Name      string `validate:"required,max=255"`
	Address   string `validate:"required,max=255"`
	Tel       string `validate:"required,max=45"`
	Content   string `validate:"required,max=1000"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ValidateShop(s *models.Shop) []string {
	var shop Shop
	var errorMessages []string // エラーメッセージ格納変数

	// set param to validator
	shop.Name = s.Name
	shop.Address = s.Address
	shop.Tel = s.Tel
	shop.Content = s.Content

	validate := validator.New()
	err := validate.Struct(shop)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			var errorMessage string
			fieldName := err.Field()
			var typ = err.Tag()

			switch typ {
			case "required":
				errorMessage = fmt.Sprintf("%sは必須項目です。", fieldName)
			case "max":
				errorMessage = fmt.Sprintf("%sは最大%s文字です。", fieldName, err.Param())
			default:
				errorMessage = fmt.Sprintf("error message for %s", fieldName)
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}
	return errorMessages
}
