package int64validator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Int64 = zeroRemainderValidator{}

type zeroRemainderValidator struct {
	Divider int64
}

// Description describes the validation in plain text formatting.
func (validator zeroRemainderValidator) Description(_ context.Context) string {
	return fmt.Sprintf("This attribute needs to be divisible by %d with zero remainder.", validator.Divider)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator zeroRemainderValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator zeroRemainderValidator) ValidateInt64(
	ctx context.Context,
	req validator.Int64Request,
	res *validator.Int64Response,
) {
	// If attribute configuration is not null or unknown, there is nothing else to validate
	if req.ConfigValue.IsNull() && req.ConfigValue.IsUnknown() {
		return
	}

	if req.ConfigValue.ValueInt64()%validator.Divider != 0 {
		res.Diagnostics.AddAttributeError(
			req.Path,
			fmt.Sprintf("This value has a non-zero remainder when divided by %d.", validator.Divider),
			validator.Description(ctx),
		)
		return
	}
}

/*
ZeroRemainder returns a validator that checks if the configured attribute is divisible by a specified integer X, and has zero remainder.
If the attribute is not divisible by X or has a non-zero remainder, the validation will fail.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func ZeroRemainder(divider int64) validator.Int64 {
	return &zeroRemainderValidator{
		Divider: divider,
	}
}
