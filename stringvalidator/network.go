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

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	networkTypes "github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/networkTypes"
)

var _ validator.String = networkValidator{}

const (
	IPV4            NetworkValidatorType = "ipv4"
	IPV4WithCIDR    NetworkValidatorType = "ipv4_with_cidr"
	IPv4WithNetmask NetworkValidatorType = "ipv4_with_netmask"
	RFC1918         NetworkValidatorType = "rfc1918"
)

type NetworkValidatorType string

type networkValidator struct {
	NetworkTypes []NetworkValidatorType
	ComparatorOR bool
}

// Description describes the validation in plain text formatting.
func (validatorNet networkValidator) Description(_ context.Context) string {
	description := ""
	switch {
	case validatorNet.ComparatorOR && len(validatorNet.NetworkTypes) > 1:
		description += "The value must be at least one of the following :\n"
	case !validatorNet.ComparatorOR && len(validatorNet.NetworkTypes) > 1:
		description += "The value must be all of the following :\n"
	case len(validatorNet.NetworkTypes) == 1:
		description += "The value must be "
	}

	for _, networkType := range validatorNet.NetworkTypes {
		switch networkType {
		case IPV4:
			description += fmt.Sprintf("%s, ", networkTypes.IsIPV4().Description(context.Background()))
		case IPV4WithCIDR:
			description += fmt.Sprintf("%s, ", networkTypes.IsIPV4WithCIDR().Description(context.Background()))
		case IPv4WithNetmask:
			description += fmt.Sprintf("%s, ", networkTypes.IsIPV4WithNetmask().Description(context.Background()))
		case RFC1918:
			description += fmt.Sprintf("%s, ", networkTypes.IsRFC1918().Description(context.Background()))
		}
	}
	return description
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validatorNet networkValidator) MarkdownDescription(ctx context.Context) string {
	markdownDescription := ""
	enableAutoTab := len(validatorNet.NetworkTypes) > 1

	autoTab := func() string {
		if enableAutoTab {
			return "  - "
		}
		return ""
	}

	autoBackToLine := func(i int) string {
		if i == len(validatorNet.NetworkTypes)-1 {
			return ""
		}
		return "\n"
	}

	computeDescription := func(markdownDescription string, i int) string {
		return fmt.Sprintf("%s%s%s", autoTab(), markdownDescription, autoBackToLine(i))
	}

	switch {
	case validatorNet.ComparatorOR && len(validatorNet.NetworkTypes) > 1:
		markdownDescription += "The value must be at least one of the following :\n"
	case !validatorNet.ComparatorOR && len(validatorNet.NetworkTypes) > 1:
		markdownDescription += "The value must be all of the following :\n"
	case len(validatorNet.NetworkTypes) == 1:
		markdownDescription += "The value must be "
	}

	for i, networkType := range validatorNet.NetworkTypes {
		switch networkType {
		case IPV4:
			markdownDescription += computeDescription(networkTypes.IsIPV4().MarkdownDescription(ctx), i)
		case IPV4WithCIDR:
			markdownDescription += computeDescription(networkTypes.IsIPV4WithCIDR().MarkdownDescription(ctx), i)
		case IPv4WithNetmask:
			markdownDescription += computeDescription(networkTypes.IsIPV4WithNetmask().MarkdownDescription(ctx), i)
		case RFC1918:
			markdownDescription += computeDescription(networkTypes.IsRFC1918().MarkdownDescription(ctx), i)
		}
	}

	return markdownDescription
}

// Validate performs the validation.
func (validatorNet networkValidator) ValidateString(
	ctx context.Context,
	request validator.StringRequest,
	response *validator.StringResponse,
) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	if len(validatorNet.NetworkTypes) == 0 {
		response.Diagnostics.AddError(
			fmt.Sprintf("Invalid configuration for attribute %s", request.Path),
			"Set at least one network type",
		)
		return
	}

	diags := diag.Diagnostics{}

	for _, networkType := range validatorNet.NetworkTypes {
		switch networkType {
		case IPV4:
			d := new(validator.StringResponse)

			networkTypes.IsIPV4().ValidateString(ctx, request, d)
			if d.Diagnostics.HasError() && !validatorNet.ComparatorOR {
				response.Diagnostics.Append(d.Diagnostics...)
			} else if d.Diagnostics.HasError() && validatorNet.ComparatorOR {
				diags.Append(d.Diagnostics...)
			}
		case IPV4WithCIDR:
			d := new(validator.StringResponse)

			networkTypes.IsIPV4WithCIDR().ValidateString(ctx, request, d)
			if d.Diagnostics.HasError() && !validatorNet.ComparatorOR {
				response.Diagnostics.Append(d.Diagnostics...)
			} else if d.Diagnostics.HasError() && validatorNet.ComparatorOR {
				diags.Append(d.Diagnostics...)
			}
		case IPv4WithNetmask:
			d := new(validator.StringResponse)

			networkTypes.IsIPV4WithNetmask().ValidateString(ctx, request, d)
			if d.Diagnostics.HasError() && !validatorNet.ComparatorOR {
				response.Diagnostics.Append(d.Diagnostics...)
			} else if d.Diagnostics.HasError() && validatorNet.ComparatorOR {
				diags.Append(d.Diagnostics...)
			}
		case RFC1918:
			d := new(validator.StringResponse)

			networkTypes.IsRFC1918().ValidateString(ctx, request, d)
			if d.Diagnostics.HasError() && !validatorNet.ComparatorOR {
				response.Diagnostics.Append(d.Diagnostics...)
			} else if d.Diagnostics.HasError() && validatorNet.ComparatorOR {
				diags.Append(d.Diagnostics...)
			}
		default:
			response.Diagnostics.AddError(
				"Invalid network type",
				fmt.Sprintf("invalid network type: %s", networkType),
			)
		}
	}

	if validatorNet.ComparatorOR && diags.ErrorsCount() == len(validatorNet.NetworkTypes) {
		response.Diagnostics.AddError(
			fmt.Sprintf("Invalid configuration for attribute %s", request.Path),
			"Set at least one valid network type",
		)
	}
}

/*
IsNetwork returns a validator that validates the string value is a valid network.

Parameters:
  - networkTypes : The network types to validate.
  - comparatorOR : If true, the value must be at least one of the network types.
*/
func IsNetwork(networkTypes []NetworkValidatorType, comparatorOR bool) validator.String {
	return &networkValidator{
		NetworkTypes: networkTypes,
		ComparatorOR: comparatorOR,
	}
}
