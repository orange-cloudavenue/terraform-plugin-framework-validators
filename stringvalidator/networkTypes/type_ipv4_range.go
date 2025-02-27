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
	"bytes"
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type validatorIPV4Range struct{}

// Description describes the validation in plain text formatting.
func (validator validatorIPV4Range) Description(_ context.Context) string {
	return "a valid IPV4 address range (Ex: 192.168.0.1-192.168.0.100)"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator validatorIPV4Range) MarkdownDescription(_ context.Context) string {
	return "a valid IPV4 address range (Ex: `192.168.0.1-192.168.0.100`)"
}

// Validate performs the validation.
func (validator validatorIPV4Range) ValidateString(
	_ context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	// Split the string into two parts
	parts := strings.Split(request.ConfigValue.ValueString(), "-")
	if len(parts) != 2 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid IPV4 range",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}

	// Check if the first IP address is less than the second IP address
	firstIP := net.ParseIP(parts[0])
	secondIP := net.ParseIP(parts[1])
	if firstIP.To4() == nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse IPV4 address",
			fmt.Sprintf("the first part of the range is not a valid IPV4 address: %s", request.ConfigValue.String()),
		)
		return
	}

	if secondIP.To4() == nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse IPV4 address",
			fmt.Sprintf("the second part of the range is not a valid IPV4 address: %s", request.ConfigValue.String()),
		)
		return
	}

	if bytes.Compare(firstIP.To4(), secondIP.To4()) >= 0 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid IPV4 range",
			fmt.Sprintf("the first part of the range is not less than the second part: %s", request.ConfigValue.String()),
		)
		return
	}
}

func IsIPV4Range() validator.String {
	return &validatorIPV4Range{}
}
