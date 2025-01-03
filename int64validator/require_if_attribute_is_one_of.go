/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

package int64validator

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// RequireIfAttributeIsOneOf checks if the path.Path attribute contains
// one of the exceptedValue attr.Value.
func RequireIfAttributeIsOneOf(path path.Expression, exceptedValue []attr.Value) validator.Int64 {
	return internal.RequireIfAttributeIsOneOf{
		PathExpression: path,
		ExceptedValues: exceptedValue,
	}
}
