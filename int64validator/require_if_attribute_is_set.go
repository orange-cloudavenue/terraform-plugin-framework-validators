package int64validator

import (
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// RequireIfAttributeIsSet checks if the path.Path attribute is set.
func RequireIfAttributeIsSet(path path.Expression) validator.Int64 {
	return internal.RequireIfAttributeIsSet{
		PathExpression: path,
	}
}
