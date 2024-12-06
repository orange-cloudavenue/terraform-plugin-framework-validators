package listvalidator_test

import (
	"context"
	"testing"

	hlistvalidator "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/listvalidator"
)

func TestNotValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         []attr.Value
		expectError bool
	}

	listValidValues := []attr.Value{types.StringValue("foo"), types.StringValue("baz"), types.StringValue("qux")}
	listInvalidValues := []attr.Value{types.StringValue("foo")}

	tests := map[string]testCase{
		"invalid": {
			val:         listInvalidValues,
			expectError: true,
		},
		"valid": {
			val:         listValidValues,
			expectError: false,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.TODO()

			values := types.ListValueMust(types.StringType, test.val)
			request := validator.ListRequest{
				ConfigValue: values,
			}
			response := validator.ListResponse{}
			listvalidator.Not(hlistvalidator.SizeBetween(1, 2)).ValidateList(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
