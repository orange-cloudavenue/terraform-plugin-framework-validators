/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package stringvalidator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator"
)

func TestFormatsValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		typesOfFormats []stringvalidator.FormatsValidatorType
		val            types.String
		expectError    bool
		ComparatorOR   bool
	}
	tests := map[string]testCase{
		"unknown": {
			val: types.StringUnknown(),
		},
		"null": {
			val: types.StringNull(),
		},
		"invalid-base64": {
			val: types.StringValue("invalidBase64"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsBase64,
			},
			expectError: true,
		},
		"valid-base64": {
			val: types.StringValue("dmFsaWQK"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsBase64,
			},
		},
		"invalid-uuid": {
			val: types.StringValue("urn:test:demo:4aeb40d8-038c-4e77-8181-a7054f583b12"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsUUIDv4,
			},
			expectError: true,
		},
		"valid-uuid": {
			val: types.StringValue("4aeb40d8-038c-4e77-8181-a7054f583b12"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsUUIDv4,
			},
		},
		"valid-urn": {
			val: types.StringValue("urn:test:demo:4aeb40d8-038c-4e77-8181-a7054f583b12"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsURN,
			},
		},
		"invalid-urn": {
			val: types.StringValue("4aeb40d8-038c-4e77-8181-a7054f583b12"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsURN,
			},
			expectError: true,
		},
		"no-validator": {
			val:            types.StringValue("dmFsaWQK"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{},
			expectError:    true,
		},
		"invalid-validator": {
			val: types.StringValue("dmFsaWQK"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				"invalid",
			},
			expectError: true,
		},
		"valid-uuid-or-urn": {
			val: types.StringValue("urn:test:demo:4aeb40d8-038c-4e77-8181-a7054f583b12"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsUUIDv4,
				stringvalidator.FormatsIsURN,
			},
			ComparatorOR: true,
			expectError:  false,
		},
		"invalid-uuid-or-urn": {
			val: types.StringValue("urn:dmFsaWQK-not-a-uuid"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsUUIDv4,
				stringvalidator.FormatsIsURN,
			},
			ComparatorOR: true,
			expectError:  true,
		},
		"invalid-uuid-and-urn": {
			val: types.StringValue("urn:dmFsaWQK-not-a-uuid"),
			typesOfFormats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsUUIDv4,
				stringvalidator.FormatsIsURN,
			},
			ComparatorOR: false,
			expectError:  true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := validator.StringRequest{
				ConfigValue: test.val,
			}
			response := validator.StringResponse{}
			stringvalidator.Formats(test.typesOfFormats, test.ComparatorOR).ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

func TestFormatsValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description  string
		Formats      []stringvalidator.FormatsValidatorType
		ComparatorOR bool
	}
	tests := map[string]testCase{
		"base64": {
			description: "The value must respect the following rule : must be a valid base64 string",
			Formats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsBase64,
			},
		},
		"uuid": {
			description: "The value must respect the following rule : must be a valid UUID v4",
			Formats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsUUIDv4,
			},
		},
		"urn": {
			description: "The value must respect the following rule : must be a valid URN",
			Formats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsURN,
			},
		},
		"uuid-or-urn": {
			description: "The value must respect at least one of the following rules :\nmust be a valid UUID v4, must be a valid URN",
			Formats: []stringvalidator.FormatsValidatorType{
				stringvalidator.FormatsIsUUIDv4,
				stringvalidator.FormatsIsURN,
			},
			ComparatorOR: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := stringvalidator.Formats(test.Formats, false)
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}

			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
