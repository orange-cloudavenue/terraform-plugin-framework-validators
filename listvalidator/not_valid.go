package listvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// Not returns a validator which ensures that the validators passed as arguments
// are not met.
func Not(valueValidator validator.List) validator.List {
	return internal.NotValidator{
		Desc:          valueValidator,
		ListValidator: valueValidator,
	}
}
