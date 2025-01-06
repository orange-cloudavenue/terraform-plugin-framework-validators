/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package int64validator

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

type OneOfWithDescriptionIfAttributeIsOneOfValues struct {
	Value       int64
	Description string
}

// OneOfWithDescriptionIfAttributeIsOneOf checks that the value is one of the expected values if the attribute is one of the exceptedValue.
// The description of the value is used to generate advanced
// Description and MarkdownDescription messages.
func OneOfWithDescriptionIfAttributeIsOneOf(path path.Expression, exceptedValue []attr.Value, values ...OneOfWithDescriptionIfAttributeIsOneOfValues) validator.String {
	frameworkValues := make([]internal.OneOfWithDescriptionIfAttributeIsOneOf, 0, len(values))

	for _, v := range values {
		frameworkValues = append(frameworkValues, internal.OneOfWithDescriptionIfAttributeIsOneOf{
			Value:       types.Int64Value(v.Value),
			Description: v.Description,
		})
	}

	return internal.OneOfWithDescriptionIfAttributeIsOneOfValidator{
		Values:         frameworkValues,
		ExceptedValues: exceptedValue,
		PathExpression: path,
	}
}
