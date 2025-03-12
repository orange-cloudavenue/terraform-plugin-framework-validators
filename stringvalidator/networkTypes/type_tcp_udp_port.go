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

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type validatorTCPUDPPort struct{}

// Description describes the validation in plain text formatting.
func (validator validatorTCPUDPPort) Description(_ context.Context) string {
	return "a valid TCP/UDP port (Ex: 8080)"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator validatorTCPUDPPort) MarkdownDescription(_ context.Context) string {
	return "a valid TCP/UDP port (Ex: `8080`)"
}

// Validate performs the validation.
func (validator validatorTCPUDPPort) ValidateString(_ context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	// convert the strings to integers
	port, err := strconv.Atoi(request.ConfigValue.ValueString())
	if err != nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid TCP/UDP port",
			fmt.Sprintf("the first part of the range is not a valid TCP/UDP port: %s", request.ConfigValue.String()),
		)
		return
	}

	// check if the integers are valid TCP or UDP ports
	if port <= 0 || port > 65535 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid TCP/UDP port",
			fmt.Sprintf("the port must be between 1 and 65535: %s", request.ConfigValue.String()),
		)
		return
	}
}

func IsTCPUDPPort() validator.String {
	return &validatorTCPUDPPort{}
}
