/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package common

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = RegexValidator{}

type RegexValidator struct {
	Desc  string
	Regex string

	ErrorSummary string
	ErrorDetail  string
}

// Description describes the validation in plain text formatting.
func (validator RegexValidator) Description(_ context.Context) string {
	return validator.Desc
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator RegexValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator RegexValidator) ValidateString(
	_ context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	re := regexp.MustCompile(validator.Regex)

	if !re.MatchString(request.ConfigValue.ValueString()) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			validator.ErrorSummary,
			validator.ErrorDetail,
		)
	}
}
