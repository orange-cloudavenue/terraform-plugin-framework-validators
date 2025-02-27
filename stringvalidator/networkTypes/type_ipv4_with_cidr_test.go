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

func TestValidIPV4WithCIDRValidator(t *testing.T) {
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
		"valid-ip-valid-cidr": {
			val: types.StringValue("192.168.1.1/24"),
		},
		"invalid-ip-valid-cidr": {
			val:         types.StringValue("192.168.1/24"),
			expectError: true,
		},
		"valid-ip-invalid-cidr": {
			val:         types.StringValue("192.168.1.1/33"),
			expectError: true,
		},
		"invalid-ipv4-valid-netmask": {
			val:         types.StringValue("2001:0db8:85a3:0000:0000:8a2e:0370:7334/32"),
			expectError: true,
		},
		"invalid-ip-no-cidr": {
			val:         types.StringValue("192.168.1"),
			expectError: true,
		},
		"valid-ip-no-cidr": {
			val:         types.StringValue("192.168.1.1"),
			expectError: true,
		},
		"ipv6": {
			val:         types.StringValue("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
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
			networktypes.IsIPV4WithCIDR().ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

// TestValidIPV4WithCIDRValidatorDescription.
func TestValidIPV4WithCIDRValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}
	tests := map[string]testCase{
		"description": {
			description: "a valid IPV4 address with CIDR (Ex: 192.168.0.1/24)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsIPV4WithCIDR()
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}
		})
	}
}

// TestValidIPV4WithCIDRValidatorMarkdownDescription.
func TestValidIPV4WithCIDRValidatorMarkdownDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}
	tests := map[string]testCase{
		"description": {
			description: "a valid IPV4 address with CIDR (Ex: `192.168.0.1/24`)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsIPV4WithCIDR()
			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
