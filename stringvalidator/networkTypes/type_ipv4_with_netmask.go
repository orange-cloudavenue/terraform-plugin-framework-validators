package networktypes

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type validatorIPV4Netmask struct{}

// Description describes the validation in plain text formatting.
func (validator validatorIPV4Netmask) Description(_ context.Context) string {
	return "a valid IPV4 address with Netmask (192.168.0.1/255.255.255.0)."
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator validatorIPV4Netmask) MarkdownDescription(_ context.Context) string {
	return "a valid IPV4 address with Netmask (`192.168.0.1/255.255.255.0`)."
}

// Validate performs the validation.
func (validator validatorIPV4Netmask) ValidateString(
	_ context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	// Split the string into IP and Netmask
	ipMask := strings.Split(request.ConfigValue.ValueString(), "/")
	if len(ipMask) != 2 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse IPV4 address with Netmask",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}

	// Validate the IP
	if net.ParseIP(ipMask[0]) == nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse IPV4 address",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}

	// To4 : If ip is not an IPv4 address, To4 returns nil.
	if net.ParseIP(ipMask[0]).To4() == nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"IP address is not IPV4",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}

	// Validate the Netmask
	stringMask := net.IPMask(net.ParseIP(ipMask[1]).To4())
	if length, _ := stringMask.Size(); length == 0 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse Netmask",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}
}

func IsIPV4WithNetmask() validator.String {
	return &validatorIPV4Netmask{}
}
