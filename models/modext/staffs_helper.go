// model helper methods
package modext

import (
	"fmt"
	"time"

	"github.com/akuruwasaki/concafe_map_mock/models"
	"github.com/go-playground/validator"
)

type Staff struct {
	Name      string `validate:"required,max=255"`
	Age       int    `validate:"required"`
	ShopID    int    `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ValidateStaff(s *models.Staff) []string {
	var staff Staff
	var errorMessages []string // エラーメッセージ格納変数

	// set param to validator
	staff.Name = s.Name
	staff.Age = s.Age
	staff.ShopID = s.ShopID

	validate := validator.New()
	err := validate.Struct(staff)

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
