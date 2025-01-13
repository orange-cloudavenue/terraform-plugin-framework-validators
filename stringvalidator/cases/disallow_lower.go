/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package cases

import (
	"context"
	"fmt"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type validatorDisallowLower struct{}

// Description describes the validation in plain text formatting.
func (validator validatorDisallowLower) Description(_ context.Context) string {
	return "disallow lowercase characters"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator validatorDisallowLower) MarkdownDescription(_ context.Context) string {
	return "disallow lowercase characters"
}

// Validate performs the validation.
func (validator validatorDisallowLower) ValidateString(
	_ context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	for _, r := range request.ConfigValue.ValueString() {
		if unicode.IsLower(r) {
			response.Diagnostics.AddAttributeError(
				request.Path,
				"lowercase characters are not allowed",
				fmt.Sprintf("invalid value: %s", request.ConfigValue.ValueString()),
			)
			return
		}
	}
}
