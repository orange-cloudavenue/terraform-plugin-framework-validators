package cases

import "github.com/hashicorp/terraform-plugin-framework/schema/validator"

type Validator interface {
	validator.String
}
