/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package formatstypes

import "github.com/hashicorp/terraform-plugin-framework/schema/validator"

/*
IsBase64 returns a validator which ensures that the configured attribute
value is a valid base64 string with base64.StdEncoding package.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func IsBase64() validator.String {
	return &base64Validator{}
}

/*
IsUUIDv4 returns a validator which ensures that the configured attribute
value is a valid UUID v4 string.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func IsUUIDv4() validator.String {
	return &uuidValidator{}
}

/*
IsURN returns a validator which ensures that the configured attribute
value is a valid URN string.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func IsURN() validator.String {
	return &urnValidator{}
}
