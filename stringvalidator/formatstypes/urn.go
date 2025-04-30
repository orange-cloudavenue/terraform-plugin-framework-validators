/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package formatstypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/common"
)

var _ validator.String = &urnValidator{}

type urnValidator struct{}

// Description describes the validation in plain text formatting.
func (validator urnValidator) Description(_ context.Context) string {
	return "must be a valid URN"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator urnValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator urnValidator) ValidateString(
	ctx context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	// Use the common regex validator to validate the URN format
	// and add the error message if it doesn't match
	// the expected URN format.
	regexValidator := common.RegexValidator{
		Desc:         "must be a valid URN",
		Regex:        `(?m)urn:[A-Za-z0-9][A-Za-z0-9-]{0,31}:([A-Za-z0-9()+,\-.:=@;$_!*']|%[0-9A-Fa-f]{2})+`,
		ErrorSummary: "Failed to parse URN",
		ErrorDetail:  "This value is not a valid URN",
	}
	regexValidator.ValidateString(ctx, request, response)
	if response.Diagnostics.HasError() {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse URN",
			"This value is not a valid URN",
		)
	}
}
