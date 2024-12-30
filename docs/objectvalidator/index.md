# ObjectValidator

Object validator are used to validate the plan of an object attribute.
It will be used into the `Validators` field of the `schema.ObjectAttribute` struct.

## How to use it

```go
import (
    fobjectvalidator "github.com/orange-cloudavenue/terraform-plugin-framework-validators/objectvalidator"
)
```

## List of Validators

- [`RequireIfAttributeIsOneOf`](../common/require_if_attribute_is_one_of.md) - This validator is used to require the attribute if another attribute is one of the given values.
- [`RequireIfAttributeIsSet`](../common/require_if_attribute_is_set.md) - This validator is used to require the attribute if another attribute is set.
- [`NullIfAttributeIsOneOf`](../common/null_if_attribute_is_one_of.md) - This validator is used to verify the attribute value is null if another attribute is one of the given values.
- [`NullIfAttributeIsSet`](../common/null_if_attribute_is_set.md) - This validator is used to verify the attribute value is null if another attribute is set.

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
