/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package int32validator

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestAttributeIsDivisibleByAnIntegerValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		req      validator.Int32Request
		in       path.Expression
		expError bool
	}

	testCases := map[string]testCase{
		"work": {
			req: validator.Int32Request{
				ConfigValue:    types.Int32Value(6),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.Int32Attribute{},
							"bar": schema.Int32Attribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.Number,
							"bar": tftypes.Number,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.Number, 12),
						"bar": tftypes.NewValue(tftypes.Number, 6),
					}),
				},
			},
			in:       path.MatchRoot("foo"),
			expError: false,
		},
		"not work": {
			req: validator.Int32Request{
				ConfigValue:    types.Int32Value(6),
				Path:           path.Root("bar"),
				PathExpression: path.MatchRoot("bar"),
				Config: tfsdk.Config{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"foo": schema.Int32Attribute{},
							"bar": schema.Int32Attribute{},
						},
					},
					Raw: tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"foo": tftypes.Number,
							"bar": tftypes.Number,
						},
					}, map[string]tftypes.Value{
						"foo": tftypes.NewValue(tftypes.Number, 13),
						"bar": tftypes.NewValue(tftypes.Number, 6),
					}),
				},
			},
			in:       path.MatchRoot("foo"),
			expError: true,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := &validator.Int32Response{}

			x := attributeIsDivisibleByAnInteger{
				PathExpression: test.in,
			}

			x.ValidateInt32(context.TODO(), test.req, res)

			if test.expError && !res.Diagnostics.HasError() {
				t.Fatal("expected error(s), got none")
			}

			if !test.expError && res.Diagnostics.HasError() {
				t.Fatalf("unexpected error(s): %s", res.Diagnostics)
			}
		})
	}
}
