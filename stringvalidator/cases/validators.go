/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package cases

import "github.com/hashicorp/terraform-plugin-framework/schema/validator"

func DisallowUpper() validator.String {
	return &validatorDisallowUpper{}
}

func DisallowLower() validator.String {
	return &validatorDisallowLower{}
}

func DisallowNumber() validator.String {
	return &validatorDisallowNumber{}
}

func DisallowSpace() validator.String {
	return &validatorDisallowSpace{}
}
