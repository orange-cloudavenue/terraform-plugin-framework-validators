/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package int32validator_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/int32validator"
)

func TestOneOfWithDescriptionValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		in        types.Int32
		validator validator.Int32
		expErrors int
	}

	testCases := map[string]testCase{
		"simple-match": {
			in: types.Int32Value(10),
			validator: int32validator.OneOfWithDescription(
				int32validator.OneOfWithDescriptionValues{
					Value:       10,
					Description: "10 description",
				},
				int32validator.OneOfWithDescriptionValues{
					Value:       20,
					Description: "20 description",
				},
				int32validator.OneOfWithDescriptionValues{
					Value:       30,
					Description: "30 description",
				},
			),
			expErrors: 0,
		},

		"simple-mismatch": {
			in: types.Int32Value(11),
			validator: int32validator.OneOfWithDescription(
				int32validator.OneOfWithDescriptionValues{
					Value:       10,
					Description: "10 description",
				},
				int32validator.OneOfWithDescriptionValues{
					Value:       20,
					Description: "20 description",
				},
				int32validator.OneOfWithDescriptionValues{
					Value:       30,
					Description: "30 description",
				},
			),
			expErrors: 1,
		},
		"skip-validation-on-null": {
			in: types.Int32Null(),
			validator: int32validator.OneOfWithDescription(
				int32validator.OneOfWithDescriptionValues{
					Value:       10,
					Description: "10 description",
				},
				int32validator.OneOfWithDescriptionValues{
					Value:       20,
					Description: "20 description",
				},
				int32validator.OneOfWithDescriptionValues{
					Value:       30,
					Description: "30 description",
				},
			),
			expErrors: 0,
		},
		"skip-validation-on-unknown": {
			in: types.Int32Unknown(),
			validator: int32validator.OneOfWithDescription(
				int32validator.OneOfWithDescriptionValues{
					Value:       10,
					Description: "10 description",
				},
				int32validator.OneOfWithDescriptionValues{
					Value:       20,
					Description: "20 description",
				},
				int32validator.OneOfWithDescriptionValues{
					Value:       30,
					Description: "30 description",
				},
			),
			expErrors: 0,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			req := validator.Int32Request{
				ConfigValue: test.in,
			}
			res := validator.Int32Response{}
			test.validator.ValidateInt32(context.TODO(), req, &res)

			if test.expErrors > 0 && !res.Diagnostics.HasError() {
				t.Fatalf("expected %d error(s), got none", test.expErrors)
			}

			if test.expErrors > 0 && test.expErrors != res.Diagnostics.ErrorsCount() {
				t.Fatalf("expected %d error(s), got %d: %v", test.expErrors, res.Diagnostics.ErrorsCount(), res.Diagnostics)
			}

			if test.expErrors == 0 && res.Diagnostics.HasError() {
				t.Fatalf("expected no error(s), got %d: %v", res.Diagnostics.ErrorsCount(), res.Diagnostics)
			}
		})
	}
}
