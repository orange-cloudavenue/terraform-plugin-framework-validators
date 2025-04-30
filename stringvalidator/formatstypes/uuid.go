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

var _ validator.String = uuidValidator{}

type uuidValidator struct{}

// Description describes the validation in plain text formatting.
func (validator uuidValidator) Description(_ context.Context) string {
	return "must be a valid UUID v4"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator uuidValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator uuidValidator) ValidateString(
	ctx context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	// Use the common regex validator to validate the UUID format
	// and add the error message if it doesn't match
	// the expected UUID format.
	regexValidator := common.RegexValidator{
		Desc: "must be a valid UUID",
		// UUID v4 regex pattern
		// https://www.ietf.org/rfc/rfc9562.txt
		// explain: https://www.bortzmeyer.org/9562.html
		Regex:        `(?m)^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`,
		ErrorSummary: "Failed to parse UUID",
		ErrorDetail:  "This value is not a valid (v4) UUID",
	}
	regexValidator.ValidateString(ctx, request, response)
}
