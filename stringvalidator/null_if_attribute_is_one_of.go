package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// NullIfAttributeIsOneOf checks if the path.Path attribute contains
// one of the exceptedValue attr.Value.
func NullIfAttributeIsOneOf(path path.Expression, exceptedValue []attr.Value) validator.String {
	return internal.NullIfAttributeIsOneOf{
		PathExpression: path,
		ExceptedValues: exceptedValue,
	}
}
