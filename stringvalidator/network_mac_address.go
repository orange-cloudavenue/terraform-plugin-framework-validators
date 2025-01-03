package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/common"
)

/*
IsMacAddress

returns a validator which ensures that the configured attribute
value is a valid MacAddress.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func IsMacAddress() validator.String {
	return &common.RegexValidator{
		Desc:         "must be a valid mac address",
		Regex:        `(?m)^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`,
		ErrorSummary: "Failed to parse mac address",
		ErrorDetail:  "This value is not a valid mac address",
	}
}
