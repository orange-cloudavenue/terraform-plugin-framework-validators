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

func TestValidHTTPCodeValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.String
		param       stringvalidator.HTTPCodeParams
		expectError bool
	}
	tests := map[string]testCase{
		"unknown": {
			val: types.StringUnknown(),
		},
		"null": {
			val: types.StringNull(),
		},
		"valid-100": {
			val: types.StringValue("100"),
			param: stringvalidator.HTTPCodeParams{
				Allow1xx: true,
			},
		},
		"valid-200": {
			val: types.StringValue("200"),
			param: stringvalidator.HTTPCodeParams{
				Allow2xx: true,
			},
		},
		"valid-300": {
			val: types.StringValue("302"),
			param: stringvalidator.HTTPCodeParams{
				Allow3xx: true,
			},
		},
		"valid-400": {
			val: types.StringValue("404"),
			param: stringvalidator.HTTPCodeParams{
				Allow4xx: true,
			},
		},
		"valid-500": {
			val: types.StringValue("500"),
			param: stringvalidator.HTTPCodeParams{
				Allow5xx: true,
			},
		},
		"valid-multi-range": {
			val: types.StringValue("404"),
			param: stringvalidator.HTTPCodeParams{
				Allow1xx: true,
				Allow2xx: true,
				Allow3xx: true,
				Allow4xx: true,
				Allow5xx: true,
			},
		},
		"invalid-400": {
			val:         types.StringValue("400"),
			param:       stringvalidator.HTTPCodeParams{},
			expectError: true,
		},
		"invalid-http-status-code": {
			val: types.StringValue("309"),
			param: stringvalidator.HTTPCodeParams{
				Allow3xx: true,
			},
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
			stringvalidator.HTTPCode(test.param).ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}

func TestValidHTTPCodeValidatorDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
		param       stringvalidator.HTTPCodeParams
	}
	tests := map[string]testCase{
		"only-100": {
			description: "The allowed HTTP status code pattern is 1xx",
			param: stringvalidator.HTTPCodeParams{
				Allow1xx: true,
			},
		},
		"only-200": {
			description: "The allowed HTTP status code pattern is 2xx",
			param: stringvalidator.HTTPCodeParams{
				Allow2xx: true,
			},
		},
		"only-300": {
			description: "The allowed HTTP status code pattern is 3xx",
			param: stringvalidator.HTTPCodeParams{
				Allow3xx: true,
			},
		},
		"only-400": {
			description: "The allowed HTTP status code pattern is 4xx",
			param: stringvalidator.HTTPCodeParams{
				Allow4xx: true,
			},
		},
		"only-500": {
			description: "The allowed HTTP status code pattern is 5xx",
			param: stringvalidator.HTTPCodeParams{
				Allow5xx: true,
			},
		},
		"multiple-ranges": {
			description: "The following HTTP status codes patterns are allowed: 1xx, 2xx, 3xx, 4xx, 5xx",
			param: stringvalidator.HTTPCodeParams{
				Allow1xx: true,
				Allow2xx: true,
				Allow3xx: true,
				Allow4xx: true,
				Allow5xx: true,
			},
		},
		"no-ranges": {
			description: "",
			param:       stringvalidator.HTTPCodeParams{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := stringvalidator.HTTPCode(test.param)
			if validator.Description(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.Description(context.Background()), test.description)
			}
		})
	}
}

func TestValidHTTPCodeValidatorMarkdownDescription(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description string
		param       stringvalidator.HTTPCodeParams
	}
	tests := map[string]testCase{
		"only-100": {
			description: "The allowed HTTP status code pattern is `1xx`",
			param: stringvalidator.HTTPCodeParams{
				Allow1xx: true,
			},
		},
		"only-200": {
			description: "The allowed HTTP status code pattern is `2xx`",
			param: stringvalidator.HTTPCodeParams{
				Allow2xx: true,
			},
		},
		"only-300": {
			description: "The allowed HTTP status code pattern is `3xx`",
			param: stringvalidator.HTTPCodeParams{
				Allow3xx: true,
			},
		},
		"only-400": {
			description: "The allowed HTTP status code pattern is `4xx`",
			param: stringvalidator.HTTPCodeParams{
				Allow4xx: true,
			},
		},
		"only-500": {
			description: "The allowed HTTP status code pattern is `5xx`",
			param: stringvalidator.HTTPCodeParams{
				Allow5xx: true,
			},
		},
		"multiple-ranges": {
			description: "The following HTTP status codes patterns are allowed: `1xx`, `2xx`, `3xx`, `4xx`, `5xx`",
			param: stringvalidator.HTTPCodeParams{
				Allow1xx: true,
				Allow2xx: true,
				Allow3xx: true,
				Allow4xx: true,
				Allow5xx: true,
			},
		},
		"no-ranges": {
			description: "",
			param:       stringvalidator.HTTPCodeParams{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := stringvalidator.HTTPCode(test.param)
			if validator.MarkdownDescription(context.Background()) != test.description {
				t.Fatalf("got unexpected description: %s != %s", validator.MarkdownDescription(context.Background()), test.description)
			}
		})
	}
}
