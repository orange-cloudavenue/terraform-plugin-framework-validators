/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package mapvalidator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	hmapvalidator "github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/mapvalidator"
)

func TestNotValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         map[string]attr.Value
		expectError bool
	}

	mapValidValues := map[string]attr.Value{"foo": types.StringValue("foo"), "baz": types.StringValue("baz"), "qux": types.StringValue("qux")}
	mapInvalidValues := map[string]attr.Value{"foo": types.StringValue("foo")}

	tests := map[string]testCase{
		"invalid": {
			val:         mapInvalidValues,
			expectError: true,
		},
		"valid": {
			val:         mapValidValues,
			expectError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.TODO()

			values := types.MapValueMust(types.StringType, test.val)
			request := validator.MapRequest{
				ConfigValue: values,
			}
			response := validator.MapResponse{}
			mapvalidator.Not(hmapvalidator.SizeBetween(1, 2)).ValidateMap(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
