/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package stringvalidator

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = httpCode{}

type httpCode struct {
	allowedRanges []httpCodeParam
}

// Description describes the validation in plain text formatting.
func (validator httpCode) Description(_ context.Context) string {
	if len(validator.allowedRanges) == 0 {
		return ""
	}

	description := ""

	if len(validator.allowedRanges) == 1 {
		description += fmt.Sprintf("The allowed HTTP status code pattern is %s", validator.allowedRanges[0].format)
	} else {
		description += "The following HTTP status codes patterns are allowed: "
		for _, r := range validator.allowedRanges {
			description += fmt.Sprintf("%s, ", r.format)
		}
		description = description[:len(description)-2] // remove the last comma and space
	}

	return description
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator httpCode) MarkdownDescription(ctx context.Context) string {
	if len(validator.allowedRanges) == 0 {
		return ""
	}

	description := ""

	if len(validator.allowedRanges) == 1 {
		description += fmt.Sprintf("The allowed HTTP status code pattern is `%s`", validator.allowedRanges[0].format)
	} else {
		description += "The following HTTP status codes patterns are allowed: "
		for _, r := range validator.allowedRanges {
			description += fmt.Sprintf("`%s`, ", r.format)
		}
		description = description[:len(description)-2] // remove the last comma and space
	}
	return description
}

// Validate performs the validation.
func (validator httpCode) ValidateString(
	_ context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	// Check if the value is a valid HTTP status code in the allowed ranges
	httpCode := request.ConfigValue.ValueString()

	// convert the string to an integer
	code, err := strconv.Atoi(httpCode)
	if err != nil {
		response.Diagnostics.AddError(
			"Invalid HTTP code",
			fmt.Sprintf("The value %s is not a valid HTTP status code", httpCode),
		)
		return
	}

	// if the status text is empty, it means the code is not valid
	if v := http.StatusText(code); v == "" {
		response.Diagnostics.AddError(
			"Invalid HTTP code",
			fmt.Sprintf("The value %s is not a valid HTTP status code defined by the HTTP RFC9110", httpCode),
		)
		return
	}

	// Check if the code is in the allowed ranges
	for _, r := range validator.allowedRanges {
		if code >= r.start && code <= r.end {
			return
		}
	}

	response.Diagnostics.AddError(
		"Invalid HTTP code",
		fmt.Sprintf("The value %s is not a valid HTTP status code in the allowed ranges", httpCode),
	)
}

type HTTPCodeParams struct {
	Allow1xx bool
	Allow2xx bool
	Allow3xx bool
	Allow4xx bool
	Allow5xx bool
}

type httpCodeParam struct {
	format string
	start  int
	end    int
}

// HTTPCode validates that a string represents a valid HTTP status code.
//
// Parameters:
//   - settings: HTTPCodeParams containing the configuration for the validator.
//
// Returns:
//   - validator.String: A validator that checks if the string is a valid HTTP status code.
func HTTPCode(settings HTTPCodeParams) validator.String {
	return &httpCode{
		allowedRanges: func() (ranges []httpCodeParam) {
			if settings.Allow1xx {
				ranges = append(ranges, httpCodeParam{start: 100, end: 199, format: "1xx"})
			}
			if settings.Allow2xx {
				ranges = append(ranges, httpCodeParam{start: 200, end: 299, format: "2xx"})
			}
			if settings.Allow3xx {
				ranges = append(ranges, httpCodeParam{start: 300, end: 399, format: "3xx"})
			}
			if settings.Allow4xx {
				ranges = append(ranges, httpCodeParam{start: 400, end: 499, format: "4xx"})
			}
			if settings.Allow5xx {
				ranges = append(ranges, httpCodeParam{start: 500, end: 599, format: "5xx"})
			}
			return ranges
		}(),
	}
}
