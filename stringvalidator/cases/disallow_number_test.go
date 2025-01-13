/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package cases_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	cases "github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/cases"
)

func TestValidDisallowNumberValidator(t *testing.T) {
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
			val: types.StringValue("onlytext"),
		},
		"invalid-lower-number": {
			val:         types.StringValue("lowerAnd123"),
			expectError: true,
		},
		"invalid-number": {
			val:         types.StringValue("123"),
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
			cases.DisallowNumber().ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

func TestValidDisallowNumberValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}
	tests := map[string]testCase{
		"description": {
			description: "disallow number characters",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := cases.DisallowNumber()
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}
		})
	}
}

func TestValidDisallowNumberValidatorMarkdownDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}
	tests := map[string]testCase{
		"description": {
			description: "disallow number characters",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := cases.DisallowNumber()
			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
