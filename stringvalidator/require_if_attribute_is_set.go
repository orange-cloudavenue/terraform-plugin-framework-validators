package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// RequireIfAttributeIsSet checks if the path.Path attribute is set.
func RequireIfAttributeIsSet(path path.Expression) validator.String {
	return internal.RequireIfAttributeIsSet{
		PathExpression: path,
	}
}
