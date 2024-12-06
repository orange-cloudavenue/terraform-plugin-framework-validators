package setvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// Not returns a validator which ensures that the validators passed as arguments
// are not met.
func Not(valueValidator validator.Set) validator.Set {
	return internal.NotValidator{
		Desc:         valueValidator,
		SetValidator: valueValidator,
	}
}
