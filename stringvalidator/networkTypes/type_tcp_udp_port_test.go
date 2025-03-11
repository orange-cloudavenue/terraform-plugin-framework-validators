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

func TestValidTCPUDPValidator(t *testing.T) {
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
			value:       types.StringValue("8080"),
			expectError: false,
		},
		"invalid": {
			value:       types.StringValue("10200-10000"),
			expectError: true,
		},
		"invalid-out": {
			value:       types.StringValue("80000"),
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
			networktypes.IsTCPUDPPort().ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

func TestValidTCPUDPValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}

	tests := map[string]testCase{
		"description": {
			description: "a valid TCP/UDP port (Ex: 8080)",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsTCPUDPPort()
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}
		})
	}
}

func TestValidTCPUDPValidatorMarkdownDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}

	tests := map[string]testCase{
		"description": {
			description: "a valid TCP/UDP port (Ex: `8080`)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsTCPUDPPort()
			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
