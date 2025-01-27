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

func TestNullIfAttributeIsSetOneOfValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		req             internal.NullIfAttributeIsSetRequest
		in              path.Expression
		inPath          path.Path
		expError        bool
		expErrorMessage string
	}

	testCases := map[string]testCase{
		"baseString": {
			req: internal.NullIfAttributeIsSetRequest{
				ConfigValue:    types.StringValue("value"),
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
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, "value"),
					}),
				},
			},
			in:              path.MatchRoot("foo"),
			inPath:          path.Root("foo"),
			expError:        true,
			expErrorMessage: "If foo attribute is set this attribute is NULL",
		},
		"extendedString": {
			req: internal.NullIfAttributeIsSetRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("foobar").AtListIndex(0).AtName("bar2"),
				PathExpression: path.MatchRoot("foobar").AtListIndex(0).AtName("bar2"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
							"foobar": schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"bar1": schema.StringAttribute{},
										"bar2": schema.StringAttribute{},
									},
								},
							},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.String,
							"bar": tftypes.String,
							"foobar": tftypes.List{
								ElementType: tftypes.Object{
									AttributeTypes: map[string]tftypes.Type{
										"bar1": tftypes.String,
										"bar2": tftypes.String,
									},
								},
							},
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
						"foobar": tftypes.NewValue(tftypes.List{
							ElementType: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"bar1": tftypes.String,
									"bar2": tftypes.String,
								},
							},
						}, []tftypes.Value{
							tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"bar1": tftypes.String,
									"bar2": tftypes.String,
								},
							}, map[string]tftypes.Value{
								"bar1": tftypes.NewValue(tftypes.String, "bar1 excepted value"),
								"bar2": tftypes.NewValue(tftypes.String, attr.NullValueString),
							}),
						},
						),
					}),
				},
			},
			in:              path.MatchRoot("foobar").AtListIndex(0).AtName("bar1"),
			inPath:          path.Root("foobar").AtListIndex(0).AtName("bar1"),
			expError:        true,
			expErrorMessage: "If foobar[0].bar1 attribute is set this attribute is NULL",
		},
		"baseInt32": {
			req: internal.NullIfAttributeIsSetRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.Int32Attribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.Number,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.Number, int32(10)),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in:              path.MatchRoot("foo"),
			inPath:          path.Root("foo"),
			expError:        true,
			expErrorMessage: "If foo attribute is set this attribute is NULL",
		},
		"baseInt64": {
			req: internal.NullIfAttributeIsSetRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.Int64Attribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.Number,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.Number, int64(10)),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in:              path.MatchRoot("foo"),
			inPath:          path.Root("foo"),
			expError:        true,
			expErrorMessage: "If foo attribute is set this attribute is NULL",
		},
		"baseBool": {
			req: internal.NullIfAttributeIsSetRequest{
				ConfigValue:    types.StringValue("bar value"),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.BoolAttribute{},
							"bar": schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.Bool,
							"bar": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.Bool, true),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in:              path.MatchRoot("foo"),
			inPath:          path.Root("foo"),
			expError:        true,
			expErrorMessage: "If foo attribute is set this attribute is NULL",
		},
		"path-attribute-is-null": {
			req: internal.NullIfAttributeIsSetRequest{
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
						"foo": tftypes.NewValue(tftypes.String, attr.NullValueString),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in:       path.MatchRoot("foo"),
			inPath:   path.Root("foo"),
			expError: false,
		},
		"config-attribute-is-null": {
			req: internal.NullIfAttributeIsSetRequest{
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
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
					}),
				},
			},
			in:       path.MatchRoot("foo"),
			inPath:   path.Root("foo"),
			expError: false,
		},
		"config-attribute-is-null-and-path-attribute-not-match": {
			req: internal.NullIfAttributeIsSetRequest{
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
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, attr.NullValueString),
					}),
				},
			},
			in:       path.MatchRoot("foo"),
			inPath:   path.Root("foo"),
			expError: false,
		},
		"unknown": {
			req: internal.NullIfAttributeIsSetRequest{
				ConfigValue:    types.StringUnknown(),
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
						"foo": tftypes.NewValue(tftypes.String, "excepted value"),
						"bar": tftypes.NewValue(tftypes.String, attr.UnknownValueString),
					}),
				},
			},
			in:              path.MatchRoot("foo"),
			inPath:          path.Root("foo"),
			expError:        true,
			expErrorMessage: "If foo attribute is set this attribute is NULL",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := &internal.NullIfAttributeIsSetResponse{}

			internal.NullIfAttributeIsSet{
				PathExpression: test.in,
			}.Validate(context.TODO(), test.req, res)

			if test.expError && res.Diagnostics.HasError() {
				if !res.Diagnostics.Contains(diag.NewAttributeErrorDiagnostic(
					test.inPath,
					fmt.Sprintf("Invalid configuration for attribute %s", test.req.Path),
					test.expErrorMessage,
				)) {
					t.Fatal(fmt.Sprintf("expected error(s) to contain (%s), got none. Error message is : (%s)", test.expErrorMessage, res.Diagnostics.Errors())) //nolint:gosimple
				}
			}

			if !test.expError && res.Diagnostics.HasError() {
				t.Fatalf("unexpected error(s): %s", res.Diagnostics)
			}

			if test.expError && !res.Diagnostics.HasError() {
				t.Fatal("expected error(s), got none")
			}
		})
	}
}
