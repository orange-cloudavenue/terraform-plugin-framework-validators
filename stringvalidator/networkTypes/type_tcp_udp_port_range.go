/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package networktypes

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type validatorTCPUDPPortRange struct{}

// Description describes the validation in plain text formatting.
func (validator validatorTCPUDPPortRange) Description(_ context.Context) string {
	return "a valid TCP/UDP port range (Ex: 1-65535)"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator validatorTCPUDPPortRange) MarkdownDescription(_ context.Context) string {
	return "a valid TCP/UDP port range (Ex: `1-65535`)"
}

// Validate performs the validation.
func (validator validatorTCPUDPPortRange) ValidateString(_ context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	// format of the string is "1-65535"
	// split the string into two parts
	p := strings.Split(request.ConfigValue.ValueString(), "-")

	if len(p) != 2 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid TCP/UDP port range",
			"the value must be in the format of `1-65535`",
		)
		return
	}

	// convert the strings to integers
	start, err := strconv.Atoi(p[0])
	if err != nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid TCP/UDP port",
			fmt.Sprintf("the first part of the range is not a valid TCP/UDP port: %s", request.ConfigValue.String()),
		)
		return
	}

	end, err := strconv.Atoi(p[1])
	if err != nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid TCP/UDP port",
			fmt.Sprintf("the second part of the range is not a valid TCP/UDP port: %s", request.ConfigValue.String()),
		)
		return
	}

	// check if the integers are valid TCP or UDP ports
	if start <= 0 || start > 65535 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid TCP/UDP port",
			fmt.Sprintf("the port must be between 1 and 65535: %s", request.ConfigValue.String()),
		)
		return
	}

	if end <= 0 || end > 65535 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid TCP/UDP port",
			fmt.Sprintf("the port must be between 1 and 65535: %s", request.ConfigValue.String()),
		)
		return
	}

	// check if the start is less than the end
	if start >= end {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid TCP/UDP port range",
			fmt.Sprintf("the first part of the range is not less than the second part: %s", request.ConfigValue.String()),
		)
		return
	}
}

func IsTCPUDPPortRange() validator.String {
	return &validatorTCPUDPPortRange{}
}
