package networktypes

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type validatorIPV4 struct{}

// Description describes the validation in plain text formatting.
func (validator validatorIPV4) Description(_ context.Context) string {
	return "a valid IPV4 address (192.168.0.1)."
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator validatorIPV4) MarkdownDescription(_ context.Context) string {
	return "a valid IPV4 address (`192.168.0.1`)."
}

// Validate performs the validation.
func (validator validatorIPV4) ValidateString(
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
			"Failed to parse IPV4 address",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}

	// To4 : If ip is not an IPv4 address, To4 returns nil.
	if net.ParseIP(request.ConfigValue.ValueString()).To4() == nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"IP address is not IPV4",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}
}

func IsIPV4() validator.String {
	return &validatorIPV4{}
}
