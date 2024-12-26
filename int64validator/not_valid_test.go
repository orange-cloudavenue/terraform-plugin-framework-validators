package int64validator_test

import (
	"context"
	"testing"

	hint64validator "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/int64validator"
)

func TestNotValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.Int64
		expectError bool
	}
	tests := map[string]testCase{
		"invalid": {
			val:         types.Int64Value(15),
			expectError: true,
		},
		"valid": {
			val: types.Int64Value(25),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := validator.Int64Request{
				ConfigValue: test.val,
			}
			response := validator.Int64Response{}
			int64validator.Not(hint64validator.Between(10, 20)).ValidateInt64(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
