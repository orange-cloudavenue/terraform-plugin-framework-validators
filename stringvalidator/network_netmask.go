package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/common"
)

/*
IsNetmask

returns a validator which ensures that the configured attribute
value is a valid Netmask.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func IsNetmask() validator.String {
	return &common.RegexValidator{
		Desc:         "must be a valid netmask",
		Regex:        `(?m)^(((255\.){3}(255|254|252|248|240|224|192|128|0+))|((255\.){2}(255|254|252|248|240|224|192|128|0+)\.0)|((255\.)(255|254|252|248|240|224|192|128|0+)(\.0+){2})|((255|254|252|248|240|224|192|128|0+)(\.0+){3}))$`,
		ErrorSummary: "Failed to parse netmask",
		ErrorDetail:  "This value is not a valid netmask",
	}
}
