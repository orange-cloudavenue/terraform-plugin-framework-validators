/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package stringvalidator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/formatstypes"
)

var _ validator.String = formatsValidator{}

const (
	FormatsIsBase64 FormatsValidatorType = "is_base64"
	FormatsIsUUIDv4 FormatsValidatorType = "is_uuid_v4"
	FormatsIsURN    FormatsValidatorType = "is_urn"
)

var formatsTypesFunc = map[FormatsValidatorType]func() validator.String{
	FormatsIsBase64: formatstypes.IsBase64,
	FormatsIsUUIDv4: formatstypes.IsUUIDv4,
	FormatsIsURN:    formatstypes.IsURN,
}

type FormatsValidatorType string

type formatsValidator struct {
	FormatsTypes []FormatsValidatorType
	ComparatorOR bool
}

// Description describes the validation in plain text formatsting.
func (validatorFormats formatsValidator) Description(ctx context.Context) string {
	description := ""

	if len(validatorFormats.FormatsTypes) == 0 {
		description += "invalid configuration"
	}

	switch {
	case !validatorFormats.ComparatorOR && len(validatorFormats.FormatsTypes) > 1:
		description += "The value must respect at least one of the following rules :\n"
	case validatorFormats.ComparatorOR && len(validatorFormats.FormatsTypes) > 1:
		description += "The value must respect all of the following rules :\n"
	case len(validatorFormats.FormatsTypes) == 1:
		description += "The value must respect the following rule : "
	}

	for i, formatsType := range validatorFormats.FormatsTypes {
		if i == len(validatorFormats.FormatsTypes)-1 {
			description += formatsTypesFunc[formatsType]().Description(ctx)
		} else {
			description += formatsTypesFunc[formatsType]().Description(ctx) + ", "
		}
	}

	return description
}

// MarkdownDescription describes the validation in Markdown formatsting.
func (validatorFormats formatsValidator) MarkdownDescription(ctx context.Context) string {
	return validatorFormats.Description(ctx)
}

// Validate performs the validation.
func (validatorFormats formatsValidator) ValidateString(
	ctx context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	if len(validatorFormats.FormatsTypes) == 0 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"No configuration provided",
			fmt.Sprintf("invalid value: %s", request.ConfigValue.String()),
		)
		return
	}

	diags := diag.Diagnostics{}

	// Check if the formats type is valid
	// and call the validation function for each formats type
	// and append the diagnostics to the response
	// If the formats type is not valid, return an error
	for _, formatsType := range validatorFormats.FormatsTypes {
		d := new(validator.StringResponse)
		if _, ok := formatsTypesFunc[formatsType]; !ok {
			response.Diagnostics.AddError(
				"Invalid formats type",
				fmt.Sprintf("invalid formats type: %s", formatsType),
			)
			return
		}
		formatsTypesFunc[formatsType]().ValidateString(ctx, request, d)
		diags.Append(d.Diagnostics...)
	}

	// Check if the value is valid for at least one of the formats types if comparatorOR is true
	if validatorFormats.ComparatorOR && diags.ErrorsCount() == len(validatorFormats.FormatsTypes) {
		response.Diagnostics.AddError(
			fmt.Sprintf("Invalid configuration for attribute %s", request.Path),
			"Set at least one valid formats type",
		)
	}

	// If the value is valid for all of the formats types if comparatorOR is false (So AND logic)
	if !validatorFormats.ComparatorOR {
		response.Diagnostics.Append(diags...)
	}
}

/*
NewFormatsValidator creates a new FormatsValidator.
It takes a list of formats types and a boolean comparatorOR.
If comparatorOR is true, the value must be at least one of the formats types.
If comparatorOR is false, the value must be all of the formats types.
If no formats types are provided, the validator will return an error.
The formats types are defined in the formatsTypes package.
*/
func Formats(formatsTypes []FormatsValidatorType, comparatorOR bool) validator.String {
	return &formatsValidator{
		FormatsTypes: formatsTypes,
		ComparatorOR: comparatorOR,
	}
}
