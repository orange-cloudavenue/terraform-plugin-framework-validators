package stringvalidator

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = netIPValidator{}

type netIPValidator struct{}

// Description describes the validation in plain text formatting.
func (validator netIPValidator) Description(_ context.Context) string {
	return "must be a valid IP with net.ParseIP"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator netIPValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator netIPValidator) ValidateString(
	_ context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	if net.ParseIP(request.ConfigValue.ValueString()) == nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse IP address",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
	}
}

/*
IsIP returns a validator which ensures that the configured attribute
value is a valid IP address with net.ParseIP package.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func IsIP() validator.String {
	return &netIPValidator{}
}
