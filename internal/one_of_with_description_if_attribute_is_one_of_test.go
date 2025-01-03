/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

package internal_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

func TestOneOfWithDescriptionIfAttributeIsOneOfValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		req            internal.OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest
		in             path.Expression
		expectedValues []attr.Value
		expError       bool
	}

	testCases := map[string]testCase{
		// If attrOther is set and the value is one of ExceptedValues the value of attrToCheck is one of Values
		// This test case return an error because the value of attrOther is one of the
		// expected values and the value of attrToCheck is not one of the Values
		"baseString": {
			req: internal.OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
				ConfigValue:    types.StringValue("another value"),
				Path:           path.Root("attrToCheck"),
				PathExpression: path.MatchRoot("attrToCheck"),
				Values: []internal.OneOfWithDescriptionIfAttributeIsOneOf{
					{
						Value:       types.StringValue("expected value"),
						Description: "expected value",
					},
				},
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"attrToCheck": schema.StringAttribute{},
							"attrOther":   schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"attrToCheck": tftypes.String,
							"attrOther":   tftypes.String,
						},
					}, map[string]tftypes.Value{
						"attrToCheck": tftypes.NewValue(tftypes.String, "another value"),
						"attrOther":   tftypes.NewValue(tftypes.String, "value"),
					}),
				},
			},
			in: path.MatchRoot("attrOther"),
			expectedValues: []attr.Value{
				types.StringValue("value"),
			},
			expError: true,
		},
		"extendedString": {
			req: internal.OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
				ConfigValue:    types.StringValue("another value"),
				Path:           path.Root("foobar").AtListIndex(0).AtName("attrToCheck"),
				PathExpression: path.MatchRoot("foobar").AtListIndex(0).AtName("attrToCheck"),
				Values: []internal.OneOfWithDescriptionIfAttributeIsOneOf{
					{
						Value:       types.StringValue("expected value"),
						Description: "expected value",
					},
				},
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.StringAttribute{},
							"bar": schema.StringAttribute{},
							"foobar": schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"attrOther":   schema.StringAttribute{},
										"attrToCheck": schema.StringAttribute{},
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
										"attrOther":   tftypes.String,
										"attrToCheck": tftypes.String,
									},
								},
							},
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.String, "foo value"),
						"bar": tftypes.NewValue(tftypes.String, "bar value"),
						"foobar": tftypes.NewValue(tftypes.List{
							ElementType: tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"attrOther":   tftypes.String,
									"attrToCheck": tftypes.String,
								},
							},
						}, []tftypes.Value{
							tftypes.NewValue(tftypes.Object{
								AttributeTypes: map[string]tftypes.Type{
									"attrOther":   tftypes.String,
									"attrToCheck": tftypes.String,
								},
							}, map[string]tftypes.Value{
								"attrToCheck": tftypes.NewValue(tftypes.String, "another value"),
								"attrOther":   tftypes.NewValue(tftypes.String, "value"),
							}),
						},
						),
					}),
				},
			},
			in: path.MatchRoot("foobar").AtListIndex(0).AtName("attrOther"),
			expectedValues: []attr.Value{
				types.StringValue("value"),
			},
			expError: true,
		},
		"baseInt64": {
			req: internal.OneOfWithDescriptionIfAttributeIsOneOfValidatorRequest{
				ConfigValue:    types.StringValue("another value"),
				Path:           path.Root("attrToCheck"),
				PathExpression: path.MatchRoot("attrToCheck"),
				Values: []internal.OneOfWithDescriptionIfAttributeIsOneOf{
					{
						Value:       types.Int64Value(20),
						Description: "20 is better",
					},
				},
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"attrToCheck": schema.StringAttribute{},
							"attrOther":   schema.StringAttribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"attrToCheck": tftypes.Number,
							"attrOther":   tftypes.String,
						},
					}, map[string]tftypes.Value{
						"attrToCheck": tftypes.NewValue(tftypes.Number, int64(10)),
						"attrOther":   tftypes.NewValue(tftypes.String, "value"),
					}),
				},
			},
			in: path.MatchRoot("attrOther"),
			expectedValues: []attr.Value{
				types.StringValue("value"),
			},
			expError: true,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := &internal.OneOfWithDescriptionIfAttributeIsOneOfValidatorResponse{}

			internal.OneOfWithDescriptionIfAttributeIsOneOfValidator{
				PathExpression: test.in,
				ExceptedValues: test.expectedValues,
				Values:         test.req.Values,
			}.Validate(
				context.Background(),
				test.req,
				res,
			)

			if !test.expError && res.Diagnostics.HasError() {
				t.Fatalf("expected no error, got %v", res.Diagnostics)
			}

			if test.expError && !res.Diagnostics.HasError() {
				t.Fatalf("expected error, got none")
			}
		})
	}
}
