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
)

// This type of validator must satisfy all types.
var (
	_ validator.Float64 = OneOfWithDescriptionValidator{}
	_ validator.Int64   = OneOfWithDescriptionValidator{}
	_ validator.List    = OneOfWithDescriptionValidator{}
	_ validator.Map     = OneOfWithDescriptionValidator{}
	_ validator.Number  = OneOfWithDescriptionValidator{}
	_ validator.Set     = OneOfWithDescriptionValidator{}
	_ validator.String  = OneOfWithDescriptionValidator{}
)

type OneOfWithDescriptionIfAttributeIsOneOf struct {
	Value       attr.Value
	Description string
}

// OneOfWithDescriptionValidator validates that the value matches one of expected values.
type OneOfWithDescriptionIfAttributeIsOneOfValidator struct {
	PathExpression path.Expression
	Values         []OneOfWithDescriptionIfAttributeIsOneOf
	ExceptedValues []attr.Value
}

type OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
	Values         []OneOfWithDescriptionIfAttributeIsOneOf
	ExceptedValues []attr.Value
}

type OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse struct {
	Diagnostics diag.Diagnostics
}

func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) Description(_ context.Context) string {
	var expectedValueDescritpion string
	for i, expectedValue := range v.ExceptedValues {
		// remove the quotes around the string
		if i == len(v.ExceptedValues)-1 {
			expectedValueDescritpion += expectedValue.String()
			break
		}
		expectedValueDescritpion += fmt.Sprintf("%s, ", expectedValue.String())
	}

	var valuesDescription string
	for i, value := range v.Values {
		if i == len(v.Values)-1 {
			valuesDescription += fmt.Sprintf("%s (%s)", value.Value.String(), value.Description)
			break
		}
		valuesDescription += fmt.Sprintf("%s (%s), ", value.Value.String(), value.Description)
	}

	switch len(v.ExceptedValues) {
	case 1:
		return fmt.Sprintf("If the value of attribute %s is %s the allowed values are : %s", v.PathExpression.String(), expectedValueDescritpion, valuesDescription)
	default:
		return fmt.Sprintf("If the value of attribute %s is one of %s the allowed are : %s", v.PathExpression.String(), expectedValueDescritpion, valuesDescription)
	}
}

func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) MarkdownDescription(_ context.Context) string {
	var expectedValueDescritpion string
	for i, expectedValue := range v.ExceptedValues {
		// remove the quotes around the string
		x := strings.Trim(expectedValue.String(), "\"")

		switch i {
		case len(v.ExceptedValues) - 1:
			expectedValueDescritpion += fmt.Sprintf("`%s`", x)
		case len(v.ExceptedValues) - 2:
			expectedValueDescritpion += fmt.Sprintf("`%s` or ", x)
		default:
			expectedValueDescritpion += fmt.Sprintf("`%s`, ", x)
		}
	}

	valuesDescription := ""
	for _, value := range v.Values {
		valuesDescription += fmt.Sprintf("- `%s` - %s<br>", value.Value.String(), value.Description)
	}

	switch len(v.ExceptedValues) {
	case 1:
		return fmt.Sprintf("\n\n-> **If the value of the attribute [`%s`](#%s) is %s the value is one of** %s", v.PathExpression, v.PathExpression, expectedValueDescritpion, valuesDescription)
	default:
		return fmt.Sprintf("\n\n-> **If the value of the attribute [`%s`](#%s) is one of %s** : %s", v.PathExpression, v.PathExpression, expectedValueDescritpion, valuesDescription)
	}
}

func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) Validate(ctx context.Context, req OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest, res *OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse) {
	// Here attribute configuration is null or unknown, so we need to check if attribute in the path
	// is equal to one of the excepted values
	paths, diags := req.Config.PathMatches(ctx, req.PathExpression.Merge(v.PathExpression))
	if diags.HasError() {
		res.Diagnostics.Append(diags...)
		return
	}

	if len(paths) == 0 {
		res.Diagnostics.AddError(
			fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
			"Path must be set",
		)
		return
	}

	path := paths[0]

	// mpVal is the value of the attribute in the path
	var mpVal attr.Value
	res.Diagnostics.Append(req.Config.GetAttribute(ctx, path, &mpVal)...)
	if res.Diagnostics.HasError() {
		res.Diagnostics.AddError(
			fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
			fmt.Sprintf("Unable to retrieve attribute path: %q", path),
		)
		return
	}

	// If the target attribute configuration is unknown or null, there is nothing else to validate
	if mpVal.IsNull() || mpVal.IsUnknown() {
		return
	}

	for _, expectedValue := range v.ExceptedValues {
		// If the value of the target attribute is equal to one of the expected values, we need to validate the value of the current attribute
		if mpVal.Equal(expectedValue) || mpVal.String() == expectedValue.String() {
			if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
				res.Diagnostics.AddAttributeError(
					path,
					fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
					fmt.Sprintf("Value is empty. %s", v.Description(ctx)),
				)
				return
			}

			for _, value := range v.Values {
				if req.ConfigValue.Equal(value.Value) {
					// Ok the value is valid
					return
				}
			}

			// The value is not valid
			res.Diagnostics.AddAttributeError(
				path,
				fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
				fmt.Sprintf("Invalid value %s. %s", req.ConfigValue.String(), v.Description(ctx)),
			)
			return
		}
	}
}

func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Float64 validates that the value matches one of expected values.
func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Int64 validates that the value matches one of expected values.
func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Number validates that the value matches one of expected values.
func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// List validates that the value matches one of expected values.
func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Set validates that the value matches one of expected values.
func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

// Map validates that the value matches one of expected values.
func (v OneOfWithDescriptionIfAttributeIsOneOfValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
		Config:      req.Config,
		ConfigValue: req.ConfigValue,
		Path:        req.Path,
	}
	validateResp := &OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse{}

	v.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}
