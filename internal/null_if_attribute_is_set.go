package internal

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// This type of validator must satisfy all types.
var (
	_ validator.Bool    = NullIfAttributeIsSet{}
	_ validator.Float64 = NullIfAttributeIsSet{}
	_ validator.Int64   = NullIfAttributeIsSet{}
	_ validator.List    = NullIfAttributeIsSet{}
	_ validator.Map     = NullIfAttributeIsSet{}
	_ validator.Number  = NullIfAttributeIsSet{}
	_ validator.Object  = NullIfAttributeIsSet{}
	_ validator.Set     = NullIfAttributeIsSet{}
	_ validator.String  = NullIfAttributeIsSet{}
)

// NullIfAttributeIsSet is the underlying struct implementing AlsoRequires.
type NullIfAttributeIsSet struct {
	PathExpression path.Expression
}

type NullIfAttributeIsSetRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type NullIfAttributeIsSetResponse struct {
	Diagnostics diag.Diagnostics
}

func (av NullIfAttributeIsSet) Description(_ context.Context) string {
	return fmt.Sprintf("If %s attribute is set this attribute is NULL", av.PathExpression)
}

func (av NullIfAttributeIsSet) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("If the [`%s`](#%s) attribute is set this attribute is **NULL**", av.PathExpression, av.PathExpression)
}

func (av NullIfAttributeIsSet) Validate(ctx context.Context, req NullIfAttributeIsSetRequest, res *NullIfAttributeIsSetResponse) {
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

		// if the attribute is null, we don't need to check the value
		if mpVal.IsNull() {
			return
		}

		res.Diagnostics.AddAttributeError(
			path,
			fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
			av.Description(ctx),
		)
	}
}

func (av NullIfAttributeIsSet) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := NullIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsSet) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := NullIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsSet) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := NullIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsSet) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := NullIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsSet) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := NullIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsSet) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := NullIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsSet) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := NullIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsSet) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := NullIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeIsSet) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := NullIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}
