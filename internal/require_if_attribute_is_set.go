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
	_ validator.Bool    = RequireIfAttributeIsSet{}
	_ validator.Float64 = RequireIfAttributeIsSet{}
	_ validator.Int64   = RequireIfAttributeIsSet{}
	_ validator.List    = RequireIfAttributeIsSet{}
	_ validator.Map     = RequireIfAttributeIsSet{}
	_ validator.Number  = RequireIfAttributeIsSet{}
	_ validator.Object  = RequireIfAttributeIsSet{}
	_ validator.Set     = RequireIfAttributeIsSet{}
	_ validator.String  = RequireIfAttributeIsSet{}
)

// RequireIfAttributeIsSet is the underlying struct implementing AlsoRequires.
type RequireIfAttributeIsSet struct {
	PathExpression path.Expression
}

type RequireIfAttributeIsSetRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type RequireIfAttributeIsSetResponse struct {
	Diagnostics diag.Diagnostics
}

func (av RequireIfAttributeIsSet) Description(_ context.Context) string {
	return fmt.Sprintf("If %s attribute is set this attribute is REQUIRED", av.PathExpression)
}

func (av RequireIfAttributeIsSet) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("If the [`%s`](#%s) attribute is set this attribute is **REQUIRED**", av.PathExpression, av.PathExpression)
}

func (av RequireIfAttributeIsSet) Validate(ctx context.Context, req RequireIfAttributeIsSetRequest, res *RequireIfAttributeIsSetResponse) {
	var diags diag.Diagnostics

	// If attribute configuration is not null, there is nothing else to validate
	if !req.ConfigValue.IsNull() {
		return
	}

	expression := req.PathExpression.Merge(av.PathExpression)

	// Here attribute configuration is null, so we need to check if attribute in the path
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

		// If the attribute configuration is null, there is nothing else to validate
		if mpVal.IsNull() {
			return
		}
		if req.ConfigValue.IsNull() {
			res.Diagnostics.AddAttributeError(
				path,
				fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
				av.Description(ctx),
			)
		}
	}
}

func (av RequireIfAttributeIsSet) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := RequireIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsSet) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := RequireIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsSet) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := RequireIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsSet) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := RequireIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsSet) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := RequireIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsSet) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := RequireIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsSet) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := RequireIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsSet) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := RequireIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av RequireIfAttributeIsSet) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := RequireIfAttributeIsSetRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &RequireIfAttributeIsSetResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}
