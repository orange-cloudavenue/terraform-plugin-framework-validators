package common

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = PrefixValidator{}

type PrefixValidator struct {
	Prefix string
}

// Description describes the validation in plain text formatting.
func (validator PrefixValidator) Description(_ context.Context) string {
	return fmt.Sprintf("must start with \"%s\"", validator.Prefix)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator PrefixValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("This value must start with `%s`.", validator.Prefix)
}

// Validate performs the validation.
func (validator PrefixValidator) ValidateString(_ context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	if !strings.HasPrefix(request.ConfigValue.ValueString(), validator.Prefix) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Does not start with prefix",
			fmt.Sprintf("The value %s does not start with the prefix \"%s\".", request.ConfigValue.String(), validator.Prefix),
		)
	}
}
