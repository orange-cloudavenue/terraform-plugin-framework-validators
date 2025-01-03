package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/common"
)

/*
IsUUID
returns a validator which ensures that the configured attribute
value is a valid (v4) UUID.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func IsUUID() validator.String {
	return &common.RegexValidator{
		Desc:         "must be a valid UUID",
		Regex:        `(?m)^\w{8}-\w{4}-\w{4}-\w{4}-\w{12}$`,
		ErrorSummary: "Failed to parse UUID",
		ErrorDetail:  "This value is not a valid (v4) UUID",
	}
}
