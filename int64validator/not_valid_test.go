/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

package int64validator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	hint64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/int64validator"
)

func TestNotValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.Int64
		expectError bool
	}
	tests := map[string]testCase{
		"invalid": {
			val:         types.Int64Value(15),
			expectError: true,
		},
		"valid": {
			val: types.Int64Value(25),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := validator.Int64Request{
				ConfigValue: test.val,
			}
			response := validator.Int64Response{}
			int64validator.Not(hint64validator.Between(10, 20)).ValidateInt64(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
