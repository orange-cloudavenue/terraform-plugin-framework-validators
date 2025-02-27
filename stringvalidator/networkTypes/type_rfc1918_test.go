/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package networktypes_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	networktypes "github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/networkTypes"
)

func TestValidRFC1918Validator(t *testing.T) {
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
		"valid-prefix-192": {
			val: types.StringValue("192.168.1.1"),
		},
		"valid-prefix-172": {
			val: types.StringValue("172.16.0.1"),
		},
		"valid-prefix-10": {
			val: types.StringValue("10.0.0.1"),
		},
		"invalid-public-ip": {
			val:         types.StringValue("1.1.1.1"),
			expectError: true,
		},
		"ipv6": {
			val:         types.StringValue("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
			expectError: true,
		},
		"invalid": {
			val:         types.StringValue("192.168.1"),
			expectError: true,
		},
		"multiple byte characters": {
			// Rightwards Arrow Over Leftwards Arrow (U+21C4; 3 bytes)
			val:         types.StringValue("⇄"),
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
			networktypes.IsRFC1918().ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

// TestValidRFC1918ValidatorDescription.
func TestValidRFC1918ValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}
	tests := map[string]testCase{
		"description": {
			description: "a valid IPV4 local address (RFC1918)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsRFC1918()
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}
		})
	}
}

// TestValidRFC1918ValidatorMarkdownDescription.
func TestValidRFC1918ValidatorMarkdownDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}
	tests := map[string]testCase{
		"description": {
			description: "a valid IPV4 local address ([RFC1918](https://en.wikipedia.org/wiki/Private_network))",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsRFC1918()
			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
