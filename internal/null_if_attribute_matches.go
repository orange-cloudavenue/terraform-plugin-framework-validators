/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// This type of validator must satisfy all types.
var (
	_ validator.Bool    = NullIfAttributeMatches{}
	_ validator.Float64 = NullIfAttributeMatches{}
	_ validator.Int32   = NullIfAttributeMatches{}
	_ validator.Int64   = NullIfAttributeMatches{}
	_ validator.List    = NullIfAttributeMatches{}
	_ validator.Map     = NullIfAttributeMatches{}
	_ validator.Number  = NullIfAttributeMatches{}
	_ validator.Object  = NullIfAttributeMatches{}
	_ validator.Set     = NullIfAttributeMatches{}
	_ validator.String  = NullIfAttributeMatches{}
)

// NullIfAttributeMatches requires the current attribute to be null when the target attribute satisfies the configured predicate.
type NullIfAttributeMatches struct {
	PathExpression          path.Expression
	Predicate               func(ctx context.Context, value attr.Value) (bool, diag.Diagnostics)
	DescriptionText         string
	MarkdownDescriptionText string
}

type NullIfAttributeMatchesRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type NullIfAttributeMatchesResponse struct {
	Diagnostics diag.Diagnostics
}

func (av NullIfAttributeMatches) Description(_ context.Context) string {
	if av.DescriptionText != "" {
		return av.DescriptionText
	}

	return fmt.Sprintf("If %s attribute matches condition this attribute is NULL", av.PathExpression)
}

func (av NullIfAttributeMatches) MarkdownDescription(_ context.Context) string {
	if av.MarkdownDescriptionText != "" {
		return av.MarkdownDescriptionText
	}

	return fmt.Sprintf("If the [`%s`](#%s) attribute matches condition this attribute is **NULL**", av.PathExpression, av.PathExpression)
}

func (av NullIfAttributeMatches) Validate(ctx context.Context, req NullIfAttributeMatchesRequest, res *NullIfAttributeMatchesResponse) {
	var diags diag.Diagnostics

	if av.Predicate == nil {
		res.Diagnostics.AddError(
			"Invalid validator configuration",
			"Predicate must be set",
		)
		return
	}

	// If attribute configuration is null, there is nothing else to validate
	if req.ConfigValue.IsNull() {
		return
	}

	// Here attribute configuration is not null, so we need to check if attribute in the path
	// matches the predicate condition
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

		// if the attribute is null or unknown, we don't need to check the value
		if mpVal.IsNull() || mpVal.IsUnknown() {
			continue
		}

		match, predicateDiags := av.Predicate(ctx, mpVal)
		res.Diagnostics.Append(predicateDiags...)
		if predicateDiags.HasError() {
			return
		}

		if match {
			res.Diagnostics.AddAttributeError(
				path,
				fmt.Sprintf("Invalid configuration for attribute %s", req.Path),
				av.Description(ctx),
			)
		}
	}
}

func (av NullIfAttributeMatches) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeMatches) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeMatches) ValidateInt32(ctx context.Context, req validator.Int32Request, resp *validator.Int32Response) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeMatches) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeMatches) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeMatches) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeMatches) ValidateNumber(ctx context.Context, req validator.NumberRequest, resp *validator.NumberResponse) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeMatches) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeMatches) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func (av NullIfAttributeMatches) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := NullIfAttributeMatchesRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &NullIfAttributeMatchesResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}
