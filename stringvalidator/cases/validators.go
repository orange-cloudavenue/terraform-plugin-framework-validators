package cases

import "github.com/hashicorp/terraform-plugin-framework/schema/validator"

func DisallowUpper() validator.String {
	return &validatorDisallowUpper{}
}

func DisallowLower() validator.String {
	return &validatorDisallowLower{}
}

func DisallowNumber() validator.String {
	return &validatorDisallowNumber{}
}

func DisallowSpace() validator.String {
	return &validatorDisallowSpace{}
}
