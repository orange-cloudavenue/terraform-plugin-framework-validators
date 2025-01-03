/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/common"
)

// PrefixContainsValidator is a validator which ensures that the configured attribute
// value contains the specified prefix.
//
// Null (unconfigured) and unknown (known after apply) values are skipped.
func PrefixContains(prefix string) validator.String {
	return &common.PrefixValidator{
		Prefix: prefix,
	}
}
