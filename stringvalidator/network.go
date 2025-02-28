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
	IPV4WithNetmask NetworkValidatorType = "ipv4_with_netmask"
	IPV4Range       NetworkValidatorType = "ipv4_range"
	RFC1918         NetworkValidatorType = "rfc1918"

	TCPUDPPortRange NetworkValidatorType = "tcpudp_port_range"
)

var networkValidatorTypes = map[NetworkValidatorType]validator.String{
	IPV4:            networkTypes.IsIPV4(),
	IPV4WithCIDR:    networkTypes.IsIPV4WithCIDR(),
	IPV4WithNetmask: networkTypes.IsIPV4WithNetmask(),
	IPV4Range:       networkTypes.IsIPV4Range(),
	RFC1918:         networkTypes.IsRFC1918(),

	TCPUDPPortRange: networkTypes.IsTCPUDPPortRange(),
}

type (
	NetworkValidatorType string

	networkValidator struct {
		NetworkTypes []NetworkValidatorType
		ComparatorOR bool
	}
)

// Description describes the validation in plain text formatting.
func (validatorNet networkValidator) Description(ctx context.Context) string {
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
		for k, v := range networkValidatorTypes {
			if networkType == k {
				description += fmt.Sprintf("%s, ", v.Description(ctx))
			}
		}
	}

	if len(validatorNet.NetworkTypes) > 1 {
		description = description[:len(description)-2]
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
		for k, v := range networkValidatorTypes {
			if networkType == k {
				markdownDescription += computeDescription(v.MarkdownDescription(ctx), i)
			}
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
		// skip validation if value is null or unknown
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
		d := new(validator.StringResponse)
		if _, ok := networkValidatorTypes[networkType]; !ok {
			response.Diagnostics.AddError(
				"Invalid network type",
				fmt.Sprintf("invalid network type: %s", networkType),
			)
			return
		}

		networkValidatorTypes[networkType].ValidateString(ctx, request, d)
		diags.Append(d.Diagnostics...)
	}

	if validatorNet.ComparatorOR && diags.ErrorsCount() == len(validatorNet.NetworkTypes) {
		response.Diagnostics.AddError(
			fmt.Sprintf("Invalid configuration for attribute %s", request.Path),
			"Set at least one valid network type",
		)
	}

	if !validatorNet.ComparatorOR {
		response.Diagnostics.Append(diags...)
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
