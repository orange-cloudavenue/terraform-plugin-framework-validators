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

package internal_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

func stringMatchPredicate() func(context.Context, attr.Value) (bool, diag.Diagnostics) {
	return func(_ context.Context, v attr.Value) (bool, diag.Diagnostics) {
		s, ok := v.(types.String)
		if !ok {
			return false, diag.Diagnostics{
				diag.NewErrorDiagnostic("predicate type error", fmt.Sprintf("expected types.String, got %T", v)),
			}
		}

		return s.ValueString() == "match", nil
	}
}

func TestNullIfAttributeMatchesValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		req                     internal.NullIfAttributeMatchesRequest
		in                      path.Expression
		inPath                  path.Path
		predicate               func(context.Context, attr.Value) (bool, diag.Diagnostics)
		descriptionText         string
		markdownDescriptionText string
		expError                bool
		expErrorMessage         string
		expConfigError          bool
		expConfigErrorMessage   string
		expPredicateError       bool
		expPredicateWarning     bool
	}

	testCases := map[string]testCase{
		"current-attribute-is-null": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringNull(),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "match"),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in:        path.MatchRoot("foo"),
			inPath:    path.Root("foo"),
			predicate: stringMatchPredicate(),
			expError:  false,
		},
		"target-attribute-is-null": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, attr.NullValueString),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:        path.MatchRoot("foo"),
			inPath:    path.Root("foo"),
			predicate: stringMatchPredicate(),
			expError:  false,
		},
		"target-attribute-is-unknown": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:        path.MatchRoot("foo"),
			inPath:    path.Root("foo"),
			predicate: stringMatchPredicate(),
			expError:  false,
		},
		"predicate-returns-true": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "match"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:              path.MatchRoot("foo"),
			inPath:          path.Root("foo"),
			predicate:       stringMatchPredicate(),
			expError:        true,
			expErrorMessage: "If foo attribute matches condition this attribute is NULL",
		},
		"predicate-returns-false": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "no match"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:        path.MatchRoot("foo"),
			inPath:    path.Root("foo"),
			predicate: stringMatchPredicate(),
			expError:  false,
		},
		"predicate-returns-warning": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "no match"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:     path.MatchRoot("foo"),
			inPath: path.Root("foo"),
			predicate: func(_ context.Context, v attr.Value) (bool, diag.Diagnostics) {
				return false, diag.Diagnostics{
					diag.NewWarningDiagnostic("predicate warning", "predicate returned a warning"),
				}
			},
			expError:            false,
			expPredicateWarning: true,
		},
		"predicate-returns-error": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "match"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:     path.MatchRoot("foo"),
			inPath: path.Root("foo"),
			predicate: func(_ context.Context, v attr.Value) (bool, diag.Diagnostics) {
				_ = v
				return false, diag.Diagnostics{
					diag.NewErrorDiagnostic("predicate error", "predicate returned an error"),
				}
			},
			expError:          false,
			expPredicateError: true,
		},
		"multi-path-one-null-one-matches": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.ListAttribute{
								ElementType: types.StringType,
							},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.List{
								ElementType: tftypes.String,
							},
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.List{
							ElementType: tftypes.String,
						}, []tftypes.Value{
							tftypes.NewValue(tftypes.String, attr.NullValueString),
							tftypes.NewValue(tftypes.String, "match"),
						}),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:              path.MatchRoot("foo").AtAnyListIndex(),
			inPath:          path.Root("foo").AtListIndex(1),
			predicate:       stringMatchPredicate(),
			expError:        true,
			expErrorMessage: "If foo[*] attribute matches condition this attribute is NULL",
		},
		"nil-predicate": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "match"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:                    path.MatchRoot("foo"),
			inPath:                path.Root("foo"),
			expConfigError:        true,
			expConfigErrorMessage: "Predicate must be set",
		},
		"custom-description": {
			req: internal.NullIfAttributeMatchesRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "match"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:              path.MatchRoot("foo"),
			inPath:          path.Root("foo"),
			predicate:       stringMatchPredicate(),
			descriptionText: "custom description",
			expError:        true,
			expErrorMessage: "custom description",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := &internal.NullIfAttributeMatchesResponse{}

			internal.NullIfAttributeMatches{
				PathExpression:          test.in,
				Predicate:               test.predicate,
				DescriptionText:         test.descriptionText,
				MarkdownDescriptionText: test.markdownDescriptionText,
			}.Validate(context.TODO(), test.req, res)

			if test.expConfigError {
				if !res.Diagnostics.Contains(diag.NewErrorDiagnostic(
					"Invalid validator configuration",
					test.expConfigErrorMessage,
				)) {
					t.Fatalf("expected configuration error (%s), got: %s", test.expConfigErrorMessage, res.Diagnostics)
				}
				return
			}

			if test.expError && res.Diagnostics.HasError() {
				if !res.Diagnostics.Contains(diag.NewAttributeErrorDiagnostic(
					test.inPath,
					fmt.Sprintf("Invalid configuration for attribute %s", test.req.Path),
					test.expErrorMessage,
				)) {
					t.Fatalf("expected error(s) to contain (%s), got none. Error message is : (%s)", test.expErrorMessage, res.Diagnostics.Errors())
				}
			}

			if !test.expError && res.Diagnostics.HasError() && !test.expPredicateError {
				t.Fatalf("unexpected error(s): %s", res.Diagnostics)
			}

			if test.expError && !res.Diagnostics.HasError() {
				t.Fatal("expected error(s), got none")
			}

			if test.expPredicateWarning && res.Diagnostics.WarningsCount() != 1 {
				t.Fatalf("expected 1 warning, got %d", res.Diagnostics.WarningsCount())
			}

			if test.expPredicateError {
				if res.Diagnostics.ErrorsCount() != 1 {
					t.Fatalf("expected 1 predicate error, got %d", res.Diagnostics.ErrorsCount())
				}

				if res.Diagnostics.Contains(diag.NewAttributeErrorDiagnostic(
					test.inPath,
					fmt.Sprintf("Invalid configuration for attribute %s", test.req.Path),
					"If foo attribute matches condition this attribute is NULL",
				)) {
					t.Fatal("expected no match error when predicate returns an error")
				}
			}
		})
	}
}

func TestNullIfAttributeMatchesValidator_Description(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	defaultValidator := internal.NullIfAttributeMatches{
		PathExpression: path.MatchRoot("foo"),
	}
	if got, want := defaultValidator.Description(ctx), "If foo attribute matches condition this attribute is NULL"; got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
	if got, want := defaultValidator.MarkdownDescription(ctx), "If the [`foo`](#foo) attribute matches condition this attribute is **NULL**"; got != want {
		t.Errorf("expected %q, got %q", want, got)
	}

	customValidator := internal.NullIfAttributeMatches{
		PathExpression:          path.MatchRoot("foo"),
		DescriptionText:         "custom plain",
		MarkdownDescriptionText: "custom markdown",
	}
	if got, want := customValidator.Description(ctx), "custom plain"; got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
	if got, want := customValidator.MarkdownDescription(ctx), "custom markdown"; got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
