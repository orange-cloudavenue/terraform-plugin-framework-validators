/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package formatstypes_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/formatstypes"
)

func TestBase64Validator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.String
		expectError bool
	}
	tests := map[string]testCase{
		"unknown": {
			val: types.StringUnknown(),
		},
		"null": {
			val: types.StringNull(),
		},
		"valid": {
			val: types.StringValue("dGVzdA=="),
		},
		"invalid": {
			val:         types.StringValue("dGVzdA"),
			expectError: true,
		},
		"invalid_with_numeric": {
			val:         types.StringValue("invalidBase64"),
			expectError: true,
		},
		"invalid_with_space": {
			val:         types.StringValue("AVec eSPace"),
			expectError: true,
		},
		"invalid_with_special_char": {
			val:         types.StringValue("AVecSPecialCh@r"),
			expectError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := validator.StringRequest{
				ConfigValue: test.val,
			}
			response := validator.StringResponse{}
			formatstypes.IsBase64().ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}
			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatal("expected no error, got error")
			}
		})
	}
}
