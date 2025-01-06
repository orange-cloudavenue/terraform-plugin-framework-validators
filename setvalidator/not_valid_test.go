/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package setvalidator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	hsetvalidator "github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/setvalidator"
)

func TestNotValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         []attr.Value
		expectError bool
	}

	listValidValues := []attr.Value{types.StringValue("foo"), types.StringValue("baz"), types.StringValue("qux")}
	listInvalidValues := []attr.Value{types.StringValue("foo")}

	tests := map[string]testCase{
		"invalid": {
			val:         listInvalidValues,
			expectError: true,
		},
		"valid": {
			val:         listValidValues,
			expectError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.TODO()

			values := types.SetValueMust(types.StringType, test.val)
			request := validator.SetRequest{
				ConfigValue: values,
			}
			response := validator.SetResponse{}
			setvalidator.Not(hsetvalidator.SizeBetween(1, 2)).ValidateSet(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
