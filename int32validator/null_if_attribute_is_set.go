/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package int32validator

import (
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// NullIfAttributeIsSet checks if the path.Path attribute contains
// one of the exceptedValue attr.Value.
func NullIfAttributeIsSet(path path.Expression) validator.Int32 {
	return internal.NullIfAttributeIsSet{
		PathExpression: path,
	}
}
