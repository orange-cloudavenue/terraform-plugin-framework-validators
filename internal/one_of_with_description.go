/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
)

const oneOfWithDescriptionValidatorDescription = "Value must be one of:"

// This type of validator must satisfy all types.
var (
	_ validator.Float64 = OneOfWithDescriptionValidator{}
	_ validator.Int32   = OneOfWithDescriptionValidator{}
	_ validator.Int64   = OneOfWithDescriptionValidator{}
	_ validator.List    = OneOfWithDescriptionValidator{}
	_ validator.Map     = OneOfWithDescriptionValidator{}
	_ validator.Number  = OneOfWithDescriptionValidator{}
	_ validator.Set     = OneOfWithDescriptionValidator{}
	_ validator.String  = OneOfWithDescriptionValidator{}
)

type OneOfWithDescription struct {
	Value       attr.Value
	Description string
}

// OneOfWithDescriptionValidator validates that the value matches one of expected values.
type OneOfWithDescriptionValidator struct {
	Values []OneOfWithDescription
}

type OneOfWithDescriptionValidatorRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
	Values         []OneOfWithDescription
}

type OneOfWithDescriptionValidatorResponse struct {
	Diagnostics diag.Diagnostics
}

func (v OneOfWithDescriptionValidator) Description(_ context.Context) string {
	var valuesDescription string
	for i, value := range v.Values {
		if i == len(v.Values)-1 {
			valuesDescription += fmt.Sprintf("%s (%s)", value.Value.String(), value.Description)
			break
		}
		valuesDescription += fmt.Sprintf("%s (%s), ", value.Value.String(), value.Description)
	}
	return fmt.Sprintf("%s %s", oneOfWithDescriptionValidatorDescription, valuesDescription)
}

func (v OneOfWithDescriptionValidator) MarkdownDescription(_ context.Context) string {
	var valuesDescription string
	for i, value := range v.Values {
		x := strings.Trim(value.Value.String(), "\"")

		if i == len(v.Values)-1 {
			valuesDescription += fmt.Sprintf("`%s` (%s)", x, value.Description)
			break
		}
		valuesDescription += fmt.Sprintf("`%s` (%s), ", x, value.Description)
	}

	return fmt.Sprintf("%s %s", oneOfWithDescriptionValidatorDescription, valuesDescription)
}

func (v OneOfWithDescriptionValidator) Validate(ctx context.Context, req OneOfWithDescriptionValidatorRequest, res *OneOfWithDescriptionValidatorResponse) {
	// If attribute configuration is not null or unknown, there is nothing else to validate
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() || len(v.Values) == 0 {
		return
	}

	for _, value := range v.Values {
		if req.ConfigValue.Equal(value.Value) {
			return
		}
	}

	res.Diagnostics.Append(validatordiag.InvalidAttributeValueMatchDiagnostic(
		req.Path,
		v.Description(ctx),
		req.ConfigValue.String(),
	))
}

func (v OneOfWithDescriptionValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := OneOfWithDescriptionValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Float64 validates that the value matches one of expected values.
func (v OneOfWithDescriptionValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := OneOfWithDescriptionValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Int32 validates that the value matches one of expected values.
func (v OneOfWithDescriptionValidator) ValidateInt32(ctx context.Context, req validator.Int32Request, resp *validator.Int32Response) {
	validateReq := OneOfWithDescriptionValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Int64 validates that the value matches one of expected values.
func (v OneOfWithDescriptionValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := OneOfWithDescriptionValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Number validates that the value matches one of expected values.
func (v OneOfWithDescriptionValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := OneOfWithDescriptionValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// List validates that the value matches one of expected values.
func (v OneOfWithDescriptionValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := OneOfWithDescriptionValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Set validates that the value matches one of expected values.
func (v OneOfWithDescriptionValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := OneOfWithDescriptionValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Map validates that the value matches one of expected values.
func (v OneOfWithDescriptionValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := OneOfWithDescriptionValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}
