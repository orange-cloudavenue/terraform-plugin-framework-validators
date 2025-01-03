package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// Not returns a validator which ensures that the validators passed as arguments
// are not met.
func Not(valueValidator validator.String) validator.String {
	return internal.NotValidator{
		Desc:            valueValidator,
		StringValidator: valueValidator,
	}
}
