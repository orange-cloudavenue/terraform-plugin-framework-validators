/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package int32validator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	hint32validator "github.com/hashicorp/terraform-plugin-framework-validators/int32validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/int32validator"
)

func TestNotValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.Int32
		expectError bool
	}
	tests := map[string]testCase{
		"invalid": {
			val:         types.Int32Value(15),
			expectError: true,
		},
		"valid": {
			val: types.Int32Value(25),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := validator.Int32Request{
				ConfigValue: test.val,
			}
			response := validator.Int32Response{}
			int32validator.Not(hint32validator.Between(10, 20)).ValidateInt32(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
