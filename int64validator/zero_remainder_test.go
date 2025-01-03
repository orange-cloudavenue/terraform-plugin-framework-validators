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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/int64validator"
)

func TestZeroRemainderValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.Int64
		divider     int64
		expectError bool
	}
	tests := map[string]testCase{
		"unknown Int64": {
			val:     types.Int64Unknown(),
			divider: 2,
		},
		"null Int64": {
			val:     types.Int64Null(),
			divider: 2,
		},
		"4/2 => OK": {
			val:     types.Int64Value(4),
			divider: 2,
		},
		"4/3 => KO": {
			val:         types.Int64Value(4),
			divider:     3,
			expectError: true,
		},
		"5/10 => KO": {
			val:         types.Int64Value(5),
			divider:     10,
			expectError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := validator.Int64Request{
				Path:           path.Root("test"),
				PathExpression: path.MatchRoot("test"),
				ConfigValue:    test.val,
			}
			response := validator.Int64Response{}
			int64validator.ZeroRemainder(test.divider).ValidateInt64(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
