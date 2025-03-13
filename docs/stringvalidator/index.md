---
hide:
    - navigation
---
# StringValidator

String validator are used to validate the plan of a string attribute.
It will be used into the `Validators` field of the `schema.StringAttribute` struct.

## How to use it

```go
import (
    fstringvalidator "github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator"
)
```

## List of Validators

- [`RequireIfAttributeIsOneOf`](../common/require_if_attribute_is_one_of.md) - This validator is used to require the attribute if another attribute is one of the given values.
- [`RequireIfAttributeIsSet`](../common/require_if_attribute_is_set.md) - This validator is used to require the attribute if another attribute is set.
- [`NullIfAttributeIsOneOf`](../common/null_if_attribute_is_one_of.md) - This validator is used to verify the attribute value is null if another attribute is one of the given values.
- [`NullIfAttributeIsSet`](../common/null_if_attribute_is_set.md) - This validator is used to verify the attribute value is null if another attribute is set.
- [`OneOfWithDescription`](oneofwithdescription.md) - This validator is used to check if the string is one of the given values and format the description and the markdown description.
- [`OneOfWithDescriptionIfAttributeIsOneOf`](../common/oneofwithdescriptionifattributeisoneof.md) - This validator is used to check if the string is one of the given values if the attribute is one of and format the description and the markdown description.

### Network

- [`IsNetwork`](isnetwork.md) - This validator is a generic validator for checking if the string is a valid network format.
- [`IsIP`](isip.md) - (**DEPRECATED**) This validator is used to check if the string is a valid IP address.
- [`IsNetmask`](isnetmask.md) - This validator is used to check if the string is a valid netmask.
- [`IsMacAddress`](ismacaddress.md) - This validator is used to check if the string is a valid MAC address.

### String

- [`IsURN`](isurn.md) - This validator is used to check if the string is a valid URN.
- [`IsUUID`](isuuid.md) - This validator is used to check if the string is a valid UUID.
- [`PrefixContains`](prefixcontains.md) - This validator is used to check if the string contains prefix in the given value.
- [`Cases`](cases.md) - This validator is a generic validator for checking if the string respects a case.

### Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.
- [`HTTPCode`](httpcode.md) - This validator is used to check if the string contains a valid http status code.
