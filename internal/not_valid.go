package internal

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var (
	_ validator.String = NotValidator{}
	_ validator.Int64  = NotValidator{}
	_ validator.Set    = NotValidator{}
	_ validator.List   = NotValidator{}
	_ validator.Bool   = NotValidator{}
	_ validator.Map    = NotValidator{}
	_ validator.Object = NotValidator{}
)

// NotValidator validates that value does not validate against the value validator.
type NotValidator struct {
	// Required
	Desc validator.Describer

	// OneOf
	StringValidator validator.String
	Int64Validator  validator.Int64
	SetValidator    validator.Set
	ListValidator   validator.List
	BoolValidator   validator.Bool
	MapValidator    validator.Map
	ObjectValidator validator.Object
}

// Description describes the validation in plain text formatting.
func (v NotValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("Value must not satisfy the validation: %s.", v.Desc.Description(ctx))
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v NotValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
// The validator will pass if it encounters a value validator that returns no errors and will then return any warnings
// from the passing validator. Using All validator as value validators will pass if all the validators supplied in an
// All validator pass.
func (v NotValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	vResp := &validator.StringResponse{
		Diagnostics: diag.Diagnostics{},
	}

	v.StringValidator.ValidateString(ctx, req, vResp)

	// If there was an error then the not condition is true, simply return
	if vResp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid not condition",
		fmt.Sprintf("NOT %s", v.StringValidator.Description(ctx)),
	)
}

// Validate performs the validation.
// The validator will pass if it encounters a value validator that returns no errors and will then return any warnings
// from the passing validator. Using All validator as value validators will pass if all the validators supplied in an
// All validator pass.
func (v NotValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	vResp := &validator.Int64Response{
		Diagnostics: diag.Diagnostics{},
	}

	v.Int64Validator.ValidateInt64(ctx, req, vResp)

	// If there was an error then the not condition is true, simply return
	if vResp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid not condition",
		fmt.Sprintf("NOT %s", v.Int64Validator.Description(ctx)),
	)
}

// Validate performs the validation.
// The validator will pass if it encounters a value validator that returns no errors and will then return any warnings
// from the passing validator. Using All validator as value validators will pass if all the validators supplied in an
// All validator pass.
func (v NotValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	vResp := &validator.SetResponse{
		Diagnostics: diag.Diagnostics{},
	}

	v.SetValidator.ValidateSet(ctx, req, vResp)

	// If there was an error then the not condition is true, simply return
	if vResp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid not condition",
		fmt.Sprintf("NOT %s", v.SetValidator.Description(ctx)),
	)
}

// Validate performs the validation.
// The validator will pass if it encounters a value validator that returns no errors and will then return any warnings
// from the passing validator. Using All validator as value validators will pass if all the validators supplied in an
// All validator pass.
func (v NotValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	vResp := &validator.ListResponse{
		Diagnostics: diag.Diagnostics{},
	}

	v.ListValidator.ValidateList(ctx, req, vResp)

	// If there was an error then the not condition is true, simply return
	if vResp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid not condition",
		fmt.Sprintf("NOT %s", v.ListValidator.Description(ctx)),
	)
}

// Validate performs the validation.
// The validator will pass if it encounters a value validator that returns no errors and will then return any warnings
// from the passing validator. Using All validator as value validators will pass if all the validators supplied in an
// All validator pass.
func (v NotValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	vResp := &validator.BoolResponse{
		Diagnostics: diag.Diagnostics{},
	}

	v.BoolValidator.ValidateBool(ctx, req, vResp)

	// If there was an error then the not condition is true, simply return
	if vResp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid not condition",
		fmt.Sprintf("NOT %s", v.BoolValidator.Description(ctx)),
	)
}

// Validate performs the validation.
// The validator will pass if it encounters a value validator that returns no errors and will then return any warnings
// from the passing validator. Using All validator as value validators will pass if all the validators supplied in an
// All validator pass.
func (v NotValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	vResp := &validator.MapResponse{
		Diagnostics: diag.Diagnostics{},
	}

	v.MapValidator.ValidateMap(ctx, req, vResp)

	// If there was an error then the not condition is true, simply return
	if vResp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid not condition",
		fmt.Sprintf("NOT %s", v.MapValidator.Description(ctx)),
	)
}

// Validate performs the validation.
// The validator will pass if it encounters a value validator that returns no errors and will then return any warnings
// from the passing validator. Using All validator as value validators will pass if all the validators supplied in an
// All validator pass.
func (v NotValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	vResp := &validator.ObjectResponse{
		Diagnostics: diag.Diagnostics{},
	}

	v.ObjectValidator.ValidateObject(ctx, req, vResp)

	// If there was an error then the not condition is true, simply return
	if vResp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid not condition",
		fmt.Sprintf("NOT %s", v.ObjectValidator.Description(ctx)),
	)
}
