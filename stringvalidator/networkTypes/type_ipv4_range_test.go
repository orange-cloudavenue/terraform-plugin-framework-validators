package networktypes_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	networktypes "github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/networkTypes"
)

func TestValidIPV4RangeValidator(t *testing.T) {
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
			value:       types.StringValue("192.168.0.1-192.168.0.10"),
			expectError: false,
		},
		"invalid": {
			value:       types.StringValue("192.168.0.100-192.168.0.10"),
			expectError: true,
		},
		"invalid-first-part": {
			value:       types.StringValue("192.168.0-192.168.0.10"),
			expectError: true,
		},
		"invalid-second-part": {
			value:       types.StringValue("192.168.0.10-notIP"),
			expectError: true,
		},
		"invalid-not-range": {
			value:       types.StringValue("ImNotARange"),
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
			networktypes.IsIPV4Range().ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

// TestValidIPV4RangeValidatorDescription.
func TestValidIPV4RangeValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}

	tests := map[string]testCase{
		"description": {
			description: "a valid IPV4 address range (Ex: 192.168.0.1-192.168.0.100)",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsIPV4Range()
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}
		})
	}
}

// TestValidIPV4RangeValidatorMarkdownDescription.
func TestValidIPV4RangeValidatorMarkdownDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
	}

	tests := map[string]testCase{
		"description": {
			description: "a valid IPV4 address range (Ex: `192.168.0.1-192.168.0.100`)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := networktypes.IsIPV4Range()
			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
