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

func TestValidTCPUDPRangeValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		value       types.String
		expectError bool
	}

	tests := map[string]testCase{
		"unknown": {
			value: types.StringUnknown(),
		},
		"null": {
			value: types.StringNull(),
		},
		"valid": {
			value:       types.StringValue("10000-10200"),
			expectError: false,
		},
		"invalid": {
			value:       types.StringValue("10200-10000"),
			expectError: true,
		},
		"invalid-first-part": {
			value:       types.StringValue("notPort-10200"),
			expectError: true,
		},
		"invalid-second-part": {
			value:       types.StringValue("10000-notPort"),
			expectError: true,
		},
		"invalid-not-range": {
			value:       types.StringValue("ImNotARange"),
			expectError: true,
		},
		"invalid-range-1": {
			value:       types.StringValue("66000-10200"),
			expectError: true,
		},
		"invalid-range-2": {
			value:       types.StringValue("10000-66200"),
			expectError: true,
		},
		"ipv6": {
			value:       types.StringValue("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
			expectError: true,
		},
		"multiple byte characters": {
			// Rightwards Arrow Over Leftwards Arrow (U+21C4; 3 bytes)
			value:       types.StringValue("⇄"),
			expectError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := validator.StringRequest{
				ConfigValue: test.value,
			}
			response := validator.StringResponse{}
			networktypes.IsTCPUDPPortRange().ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

// TestValidTCPUDPRangeValidatorDescription.
func TestValidTCPUDPRangeValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}

	tests := map[string]testCase{
		"description": {
			description: "a valid TCP/UDP port range (Ex: 1-65535)",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsTCPUDPPortRange()
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}
		})
	}
}

// TestValidTCPUDPRangeValidatorMarkdownDescription.
func TestValidTCPUDPRangeValidatorMarkdownDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}

	tests := map[string]testCase{
		"description": {
			description: "a valid TCP/UDP port range (Ex: `1-65535`)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsTCPUDPPortRange()
			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
