package stringvalidator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator"
)

func TestCasesValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		typesOfCases []stringvalidator.CasesValidatorType
		val          types.String
		ComparatorOR bool
		expectError  bool
	}
	tests := map[string]testCase{
		"unknown": {
			val: types.StringUnknown(),
		},
		"null": {
			val: types.StringNull(),
		},
		"disallow-upper": {
			val: types.StringValue("lowerAndUPPER"),
			typesOfCases: []stringvalidator.CasesValidatorType{
				stringvalidator.CasesDisallowUpper,
			},
			expectError: true,
		},
		"disallow-number": {
			val: types.StringValue("123"),
			typesOfCases: []stringvalidator.CasesValidatorType{
				stringvalidator.CasesDisallowNumber,
			},
			expectError: true,
		},
		"disallow-upper-and-number": {
			val: types.StringValue("lowerAnd123"),
			typesOfCases: []stringvalidator.CasesValidatorType{
				stringvalidator.CasesDisallowUpper,
				stringvalidator.CasesDisallowNumber,
			},
			expectError: true,
		},
		"valid": {
			val: types.StringValue("onlylower"),
			typesOfCases: []stringvalidator.CasesValidatorType{
				stringvalidator.CasesDisallowUpper,
				stringvalidator.CasesDisallowNumber,
			},
		},
		"no-validator": {
			val:          types.StringValue("lowerAndUPPER"),
			typesOfCases: []stringvalidator.CasesValidatorType{},
			expectError:  true,
		},
		"invalid-validator": {
			val: types.StringValue("lowerAndUPPER"),
			typesOfCases: []stringvalidator.CasesValidatorType{
				"invalid",
			},
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
			stringvalidator.Cases(test.typesOfCases).ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

func TestCasesValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
		cases       []stringvalidator.CasesValidatorType
	}
	tests := map[string]testCase{
		"disallow-upper": {
			description: "The value must respect the following rule : disallow uppercase characters",
			cases: []stringvalidator.CasesValidatorType{
				stringvalidator.CasesDisallowUpper,
			},
		},
		"disallow-number": {
			description: "The value must respect the following rule : disallow number characters",
			cases: []stringvalidator.CasesValidatorType{
				stringvalidator.CasesDisallowNumber,
			},
		},
		"disallow-upper-and-number": {
			description: "The value must respect the following rules : disallow uppercase characters, disallow number characters",
			cases: []stringvalidator.CasesValidatorType{
				stringvalidator.CasesDisallowUpper,
				stringvalidator.CasesDisallowNumber,
			},
		},
		"no-validator": {
			description: "invalid configuration",
			cases:       []stringvalidator.CasesValidatorType{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := stringvalidator.Cases(test.cases)
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}

			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
