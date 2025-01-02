package mapvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// NullIfAttributeIsSet checks if the path.Path attribute contains
// one of the exceptedValue attr.Value.
func NullIfAttributeIsSet(path path.Expression) validator.Map {
	return internal.NullIfAttributeIsSet{
		PathExpression: path,
	}
}