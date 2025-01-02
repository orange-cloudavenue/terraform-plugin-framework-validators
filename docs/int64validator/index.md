---
hide:
    - navigation
---

# Int64Validator

Int64 validator are used to validate the plan of a int64 attribute.
It will be used into the `Validators` field of the `schema.Int64Attribute` struct.

## How to use it

```go
import (
    fint64validator "github.com/orange-cloudavenue/terraform-plugin-framework-validators/int64validator"
)
```

## List of Validators

- [`RequireIfAttributeIsOneOf`](../common/require_if_attribute_is_one_of.md) - This validator is used to require the attribute if another attribute is one of the given values.
- [`RequireIfAttributeIsSet`](../common/require_if_attribute_is_set.md) - This validator is used to require the attribute if another attribute is set.
- [`NullIfAttributeIsOneOf`](../common/null_if_attribute_is_one_of.md) - This validator is used to verify the attribute value is null if another attribute is one of the given values.
- [`NullIfAttributeIsSet`](../common/null_if_attribute_is_set.md) - This validator is used to verify the attribute value is null if another attribute is set.
- [`OneOfWithDescription`](oneofwithdescription.md) - This validator is used to check if the string is one of the given values and format the description and the markdown description.
- [`OneOfWithDescriptionIfAttributeIsOneOf`](../common/oneofwithdescriptionifattributeisoneof.md) - This validator is used to check if the string is one of the given values if the attribute is one of and format the description and the markdown description.
- [`AttributeIsDivisibleByAnInteger`](attribute_is_divisible_by_an_integer.md) - This validator is used to validate that the attribute is divisible by an integer.
- [`ZeroRemainder`](zero_remainder.md) - This validator checks if the configured attribute is divisible by a specified integer X, and has zero remainder.

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
