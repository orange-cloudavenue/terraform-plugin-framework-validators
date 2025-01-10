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

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	casesTypes "github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/cases"
)

var _ validator.String = casesValidator{}

const (
	CasesDisallowUpper  CasesValidatorType = "disallow_upper"
	CasesDisallowNumber CasesValidatorType = "disallow_number"
	CasesDisallowSpace  CasesValidatorType = "disallow_space"
	CasesDisallowLower  CasesValidatorType = "disallow_lower"
)

var casesTypesFunc = map[CasesValidatorType]func() validator.String{
	CasesDisallowUpper:  casesTypes.DisallowUpper,
	CasesDisallowNumber: casesTypes.DisallowNumber,
	CasesDisallowSpace:  casesTypes.DisallowSpace,
	CasesDisallowLower:  casesTypes.DisallowLower,
}

type CasesValidatorType string

type casesValidator struct {
	CasesTypes []CasesValidatorType
}

// Description describes the validation in plain text formatting.
func (validatorCase casesValidator) Description(_ context.Context) string {
	description := ""

	if len(validatorCase.CasesTypes) == 0 {
		description += "invalid configuration"
	}

	switch {
	case len(validatorCase.CasesTypes) > 1:
		description += "The value must respect the following rules : "
	case len(validatorCase.CasesTypes) == 1:
		description += "The value must respect the following rule : "
	}

	for i, networkType := range validatorCase.CasesTypes {
		if i == len(validatorCase.CasesTypes)-1 {
			description += casesTypesFunc[networkType]().Description(context.Background())
		} else {
			description += fmt.Sprintf("%s, ", casesTypesFunc[networkType]().Description(context.Background()))
		}
	}
	return description
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validatorCase casesValidator) MarkdownDescription(ctx context.Context) string {
	return validatorCase.Description(ctx)
}

// Validate performs the validation.
func (validatorCase casesValidator) ValidateString(
	ctx context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	if len(validatorCase.CasesTypes) == 0 {
		response.Diagnostics.AddError(
			fmt.Sprintf("Invalid configuration for attribute %s", request.Path),
			"Set at least one case type",
		)
		return
	}

	for _, networkType := range validatorCase.CasesTypes {
		if _, ok := casesTypesFunc[networkType]; !ok {
			response.Diagnostics.AddError(
				"Invalid case type",
				fmt.Sprintf("invalid case type: %s", networkType),
			)
			continue
		}

		resp := new(validator.StringResponse)
		casesTypesFunc[networkType]().ValidateString(ctx, request, resp)

		if resp.Diagnostics.HasError() {
			response.Diagnostics.Append(resp.Diagnostics...)
		}
	}
}

// Cases returns a new string validator that checks if the string matches any of the specified case types.
//
// Parameters:
//   - casesTypes: A slice of CasesValidatorType that specifies the types of cases to validate against.
//
// Returns:
//   - validator.String: A string validator that validates the string against the specified case types.
func Cases(casesTypes []CasesValidatorType) validator.String {
	return &casesValidator{
		CasesTypes: casesTypes,
	}
}
