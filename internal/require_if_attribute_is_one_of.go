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
	_ validator.Bool    = RequireIfAttributeIsOneOf{}
	_ validator.Float64 = RequireIfAttributeIsOneOf{}
	_ validator.Int64   = RequireIfAttributeIsOneOf{}
	_ validator.List    = RequireIfAttributeIsOneOf{}
	_ validator.Map     = RequireIfAttributeIsOneOf{}
	_ validator.Number  = RequireIfAttributeIsOneOf{}
	_ validator.Object  = RequireIfAttributeIsOneOf{}
	_ validator.Set     = RequireIfAttributeIsOneOf{}
	_ validator.String  = RequireIfAttributeIsOneOf{}
)

// RequireIfAttributeIsOneOf is the underlying struct implementing AlsoRequires.
type RequireIfAttributeIsOneOf struct {
	PathExpression path.Expression
	ExceptedValues []attr.Value
}

type RequireIfAttributeIsOneOfRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
	ExceptedValues []attr.Value
}

type RequireIfAttributeIsOneOfResponse struct {
	Diagnostics diag.Diagnostics
}

func (av RequireIfAttributeIsOneOf) Description(_ context.Context) string {
	var expectedValueDescritpion string
	for i, expectedValue := range av.ExceptedValues {
		if i == len(av.ExceptedValues)-1 {
			expectedValueDescritpion += expectedValue.String()
			break
		}
		expectedValueDescritpion += fmt.Sprintf("%s, ", expectedValue.String())
	}

	if len(av.ExceptedValues) == 1 {
		return fmt.Sprintf("If %s attribute is set and the value is %s this attribute is REQUIRED", av.PathExpression, expectedValueDescritpion)
	}
	return fmt.Sprintf("If %s attribute is set and the value is one of %s this attribute is REQUIRED", av.PathExpression, expectedValueDescritpion)
}

func (av RequireIfAttributeIsOneOf) MarkdownDescription(_ context.Context) string {
	var expectedValueDescritpion string
	for i, expectedValue := range av.ExceptedValues {
		// remove the quotes around the string
		v := strings.Trim(expectedValue.String(), "\"")

		switch i {
		case len(av.ExceptedValues) - 1:
			expectedValueDescritpion += fmt.Sprintf("`%s`", v)
		case len(av.ExceptedValues) - 2:
			expectedValueDescritpion += fmt.Sprintf("`%s` or ", v)
		default:
			expectedValueDescritpion += fmt.Sprintf("`%s`, ", v)
		}
	}

	switch len(av.ExceptedValues) {
	case 1:
		return fmt.Sprintf("If the value of [`%s`](#%s) attribute is %s this attribute is **REQUIRED**", av.PathExpression, av.PathExpression, expectedValueDescritpion)
	default:
		return fmt.Sprintf("If the value of [`%s`](#%s) attribute is one of %s this attribute is **REQUIRED**", av.PathExpression, av.PathExpression, expectedValueDescritpion)
	}
}

func (av RequireIfAttributeIsOneOf) Validate(ctx context.Context, req RequireIfAttributeIsOneOfRequest, res *RequireIfAttributeIsOneOfResponse) {
	var diags diag.Diagnostics

	// If attribute configuration is not null or unknown, there is nothing else to validate
	if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
		return
	}

	expression := req.PathExpression.Merge(av.PathExpression)

	// Here attribute configuration is null or unknown, so we need to check if attribute in the path
	// is equal to one of the excepted values
	paths, diags := req.Config.PathMatches(ctx, expression)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	if len(paths) == 0 {
		res.Diagnostics.AddError(
			fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
			"Path must be set",
		)
		return
	}

	for _, path := range paths {
		var mpVal attr.Value
		diags = req.Config.GetAttribute(ctx, path, &mpVal)
		if diags.HasError() {
			res.Diagnostics.AddError(
				fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
				fmt.Sprintf("Unable to retrieve attribute path: %q", path),
			)
			return
		}

		// If the attribute configuration is unknown, there is nothing else to validate
		if mpVal.IsNull() || mpVal.IsUnknown() {
			return
		}

		for _, expectedValue := range av.ExceptedValues {
			if mpVal.Equal(expectedValue) {
				if req.ConfigValue.IsNull() {
					res.Diagnostics.AddAttributeError(
						path,
						fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
						av.Description(ctx),
					)
				}
				return
			}
		}
	}
}

func (av RequireIfAttributeIsOneOf) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := RequireIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsOneOf) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := RequireIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsOneOf) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := RequireIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsOneOf) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := RequireIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsOneOf) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := RequireIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsOneOf) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := RequireIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsOneOf) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := RequireIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsOneOf) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := RequireIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsOneOf) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := RequireIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}
