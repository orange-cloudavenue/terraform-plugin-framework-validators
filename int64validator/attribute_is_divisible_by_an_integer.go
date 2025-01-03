package int64validator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Int64 = attributeIsDivisibleByAnInteger{}

type attributeIsDivisibleByAnInteger struct {
	PathExpression path.Expression
}

// Description describes the validation in plain text formatting.
func (validator attributeIsDivisibleByAnInteger) Description(_ context.Context) string {
	return fmt.Sprintf("All the possibilities of dividing the value of attribute %s by an integer.", validator.PathExpression.String())
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator attributeIsDivisibleByAnInteger) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator attributeIsDivisibleByAnInteger) ValidateInt64(
	ctx context.Context,
	req validator.Int64Request,
	res *validator.Int64Response,
) {
	// If attribute configuration is not null or unknown, there is nothing else to validate
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	paths, diags := req.Config.PathMatches(ctx, req.PathExpression.Merge(validator.PathExpression))
	res.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	if len(paths) == 0 {
		res.Diagnostics.AddAttributeError(
			paths[0],
			"Invalid configuration",
			"Path must be set",
		)
		return
	}

	if len(paths) > 1 {
		res.Diagnostics.AddAttributeError(
			paths[0],
			"Invalid configuration",
			"Path must be unique",
		)
		return
	}

	var attributeValue types.Int64

	diags = req.Config.GetAttribute(ctx, paths[0], &attributeValue)
	if diags.HasError() {
		res.Diagnostics.AddAttributeError(
			paths[0],
			"Invalid configuration",
			"Unable to retrieve attribute path",
		)
	}

	if attributeValue.IsNull() || attributeValue.IsUnknown() {
		return
	}

	// Logic to validate the value.
	// If the valueFromAttribute is divisible by configValue and the remainder is 0, then return true.
	// Examples:
	// - valueFromAttribute = 10, configValue = 2, return true
	// - valueFromAttribute = 10, configValue = 3, return false

	if attributeValue.ValueInt64()%req.ConfigValue.ValueInt64() != 0 {
		res.Diagnostics.AddAttributeError(
			paths[0],
			"Invalid configuration",
			validator.Description(ctx),
		)
		return
	}
}

/*
AttributeIsDivisibleByAnInteger returns a validator that ensures the configured attribute is divisible by an integer.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func AttributeIsDivisibleByAnInteger(path path.Expression) validator.Int64 {
	return &attributeIsDivisibleByAnInteger{
		PathExpression: path,
	}
}
