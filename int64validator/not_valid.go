package int64validator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// Not returns a validator which ensures that the validators passed as arguments
// are not met.
func Not(valueValidator validator.Int64) validator.Int64 {
	return internal.NotValidator{
		Desc:           valueValidator,
		Int64Validator: valueValidator,
	}
}
