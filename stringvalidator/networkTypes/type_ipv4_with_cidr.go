/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

package networktypes

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type validatorIPV4CIDR struct{}

// Description describes the validation in plain text formatting.
func (validator validatorIPV4CIDR) Description(_ context.Context) string {
	return "a valid IPV4 address with CIDR (192.168.0.1/24)."
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator validatorIPV4CIDR) MarkdownDescription(_ context.Context) string {
	return "a valid IPV4 address with CIDR (`192.168.0.1/24`)."
}

// Validate performs the validation.
func (validator validatorIPV4CIDR) ValidateString(
	_ context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	netIP, _, err := net.ParseCIDR(request.ConfigValue.ValueString())
	if err != nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse IPV4 address with CIDR",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}

	// To4 : If ip is not an IPv4 address, To4 returns nil.
	if netIP.To4() == nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"IP address is not IPV4",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}
}

func IsIPV4WithCIDR() validator.String {
	return &validatorIPV4CIDR{}
}
