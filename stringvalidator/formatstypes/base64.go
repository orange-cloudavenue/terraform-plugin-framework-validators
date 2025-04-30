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
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = base64Validator{}

type base64Validator struct{}

// Description describes the validation in plain text formatting.
func (validator base64Validator) Description(_ context.Context) string {
	return "must be a valid base64 string"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator base64Validator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator base64Validator) ValidateString(
	_ context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	if _, err := base64.StdEncoding.DecodeString(request.ConfigValue.ValueString()); err != nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Failed to parse base64 string",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
	}
}
