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
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type validatorDisallowNumber struct{}

// Description describes the validation in plain text formatting.
func (validator validatorDisallowNumber) Description(_ context.Context) string {
	return "disallow number characters"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator validatorDisallowNumber) MarkdownDescription(_ context.Context) string {
	return "disallow number characters"
}

// Validate performs the validation.
func (validator validatorDisallowNumber) ValidateString(
	_ context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	if regexp.MustCompile(`\d`).MatchString(request.ConfigValue.ValueString()) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"number characters are not allowed",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.ValueString()),
		)
		return
	}
}
