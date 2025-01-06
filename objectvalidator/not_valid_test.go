/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package objectvalidator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	hobjectvalidator "github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/objectvalidator"
)

func TestNotValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.Object
		expectError bool
	}

	tests := map[string]testCase{
		"valid-null": {
			val: types.ObjectNull(
				map[string]attr.Type{
					"field1": types.StringType,
				},
			),
			expectError: false,
		},
		"invalid-unknown": {
			val: types.ObjectUnknown(
				map[string]attr.Type{
					"field1": types.StringType,
				},
			),
			expectError: true,
		},
		"invalid-empty": {
			val: types.ObjectValueMust(
				map[string]attr.Type{
					"field1": types.StringType,
				},
				map[string]attr.Value{
					"field1": types.StringNull(),
				},
			),
			expectError: true,
		},
		"invalid-elements": {
			val: types.ObjectValueMust(
				map[string]attr.Type{
					"field1": types.StringType,
				},
				map[string]attr.Value{
					"field1": types.StringValue("value1"),
				},
			),
			expectError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.TODO()

			request := validator.ObjectRequest{
				Path:           path.Root("test"),
				PathExpression: path.MatchRoot("test"),
				ConfigValue:    test.val,
			}
			response := validator.ObjectResponse{}
			objectvalidator.Not(hobjectvalidator.IsRequired()).ValidateObject(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
