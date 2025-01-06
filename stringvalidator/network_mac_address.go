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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/common"
)

/*
IsMacAddress

returns a validator which ensures that the configured attribute
value is a valid MacAddress.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func IsMacAddress() validator.String {
	return &common.RegexValidator{
		Desc:         "must be a valid mac address",
		Regex:        `(?m)^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`,
		ErrorSummary: "Failed to parse mac address",
		ErrorDetail:  "This value is not a valid mac address",
	}
}
