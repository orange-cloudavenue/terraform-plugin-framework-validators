package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/common"
)

// PrefixContainsValidator is a validator which ensures that the configured attribute
// value contains the specified prefix.
//
// Null (unconfigured) and unknown (known after apply) values are skipped.
func PrefixContains(prefix string) validator.String {
	return &common.PrefixValidator{
		Prefix: prefix,
	}
}
