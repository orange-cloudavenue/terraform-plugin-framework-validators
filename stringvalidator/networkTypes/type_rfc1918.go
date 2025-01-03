package networktypes

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type validatorRFC1918 struct{}

// Description describes the validation in plain text formatting.
func (validator validatorRFC1918) Description(_ context.Context) string {
	return "a valid IPV4 local address (RFC1918)."
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator validatorRFC1918) MarkdownDescription(_ context.Context) string {
	return "a valid IPV4 local address ([RFC1918](https://en.wikipedia.org/wiki/Private_network))."
}

// Validate performs the validation.
func (validator validatorRFC1918) ValidateString(
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

	if !net.ParseIP(request.ConfigValue.ValueString()).IsPrivate() {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"IP address is not RFC1918",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}
}

func IsRFC1918() validator.String {
	return &validatorRFC1918{}
}
