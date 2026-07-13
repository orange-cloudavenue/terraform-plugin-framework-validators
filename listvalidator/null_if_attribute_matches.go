/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package listvalidator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/internal"
)

// NullIfAttributeMatches checks if the path.Path attribute matches the predicate condition.
func NullIfAttributeMatches(path path.Expression, predicate func(ctx context.Context, value attr.Value) (bool, diag.Diagnostics)) validator.List {
	return internal.NullIfAttributeMatches{
		PathExpression: path,
		Predicate:      predicate,
	}
}

// NullIfAttributeMatchesDescription is a struct to hold the description and markdown description for NullIfAttributeMatchesWithDescription.
type NullIfAttributeMatchesDescription struct {
	Description         string
	MarkdownDescription string
}

// NullIfAttributeMatchesWithDescription checks if the path.Path attribute matches the predicate condition and allows overriding the validator description.
func NullIfAttributeMatchesWithDescription(
	path path.Expression,
	predicate func(ctx context.Context, value attr.Value) (bool, diag.Diagnostics),
	description NullIfAttributeMatchesDescription,
) validator.List {
	return internal.NullIfAttributeMatches{
		PathExpression:          path,
		Predicate:               predicate,
		DescriptionText:         description.Description,
		MarkdownDescriptionText: description.MarkdownDescription,
	}
}
