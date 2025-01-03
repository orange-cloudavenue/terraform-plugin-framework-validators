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
	_ validator.Bool    = NullIfAttributeIsOneOf{}
	_ validator.Float64 = NullIfAttributeIsOneOf{}
	_ validator.Int64   = NullIfAttributeIsOneOf{}
	_ validator.List    = NullIfAttributeIsOneOf{}
	_ validator.Map     = NullIfAttributeIsOneOf{}
	_ validator.Number  = NullIfAttributeIsOneOf{}
	_ validator.Object  = NullIfAttributeIsOneOf{}
	_ validator.Set     = NullIfAttributeIsOneOf{}
	_ validator.String  = NullIfAttributeIsOneOf{}
)

// NullIfAttributeIsOneOf is the underlying struct implementing AlsoRequires.
type NullIfAttributeIsOneOf struct {
	PathExpression path.Expression
	ExceptedValues []attr.Value
}

type NullIfAttributeIsOneOfRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
	ExceptedValues []attr.Value
}

type NullIfAttributeIsOneOfResponse struct {
	Diagnostics diag.Diagnostics
}

func (av NullIfAttributeIsOneOf) Description(_ context.Context) string {
	var expectedValueDescritpion string
	for i, expectedValue := range av.ExceptedValues {
		// remove the quotes around the string
		if i == len(av.ExceptedValues)-1 {
			expectedValueDescritpion += expectedValue.String()
			break
		}
		expectedValueDescritpion += fmt.Sprintf("%s, ", expectedValue.String())
	}

	if len(av.ExceptedValues) == 1 {
		return fmt.Sprintf("If %s attribute is set and the value is %s this attribute is NULL", av.PathExpression, expectedValueDescritpion)
	}
	return fmt.Sprintf("If %s attribute is set and the value is one of %s this attribute is NULL", av.PathExpression, expectedValueDescritpion)
}

func (av NullIfAttributeIsOneOf) MarkdownDescription(_ context.Context) string {
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
		return fmt.Sprintf("If the value of [`%s`](#%s) attribute is %s this attribute is **NULL**", av.PathExpression, av.PathExpression, expectedValueDescritpion)
	default:
		return fmt.Sprintf("If the value of [`%s`](#%s) attribute is one of %s this attribute is **NULL**", av.PathExpression, av.PathExpression, expectedValueDescritpion)
	}
}

func (av NullIfAttributeIsOneOf) Validate(ctx context.Context, req NullIfAttributeIsOneOfRequest, res *NullIfAttributeIsOneOfResponse) {
	var diags diag.Diagnostics

	// If attribute configuration is null, there is nothing else to validate
	if req.ConfigValue.IsNull() {
		return
	}

	// Here attribute configuration is null or unknown, so we need to check if attribute in the path
	// is equal to one of the excepted values
	paths, diags := req.Config.PathMatches(ctx, req.PathExpression.Merge(av.PathExpression))
	res.Diagnostics.Append(diags...)
	if diags.HasError() {
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
		// TODO
		if mpVal.IsNull() || mpVal.IsUnknown() {
			return
		}

		for _, expectedValue := range av.ExceptedValues {
			if mpVal.Equal(expectedValue) {
				res.Diagnostics.AddAttributeError(
					path,
					fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
					av.Description(ctx),
				)
				return
			}
		}
	}
}

func (av NullIfAttributeIsOneOf) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := NullIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsOneOf) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := NullIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsOneOf) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := NullIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsOneOf) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := NullIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsOneOf) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := NullIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsOneOf) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := NullIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsOneOf) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := NullIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsOneOf) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := NullIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsOneOf) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := NullIfAttributeIsOneOfRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsOneOfResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}
