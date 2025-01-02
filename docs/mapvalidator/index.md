---
hide:
    - navigation
---
# MapValidator

Map validator are used to validate the plan of a map attribute.
It will be used into the `Validators` field of the `schema.MapAttribute` struct.

## How to use it

```go
import (
    fmapvalidator "github.com/orange-cloudavenue/terraform-plugin-framework-validators/mapvalidator"
)
```

## List of Validators

Every `string` validators are available for maps thanks to a generic validator provided by Hashicorp. See the section below for more details.

- [`RequireIfAttributeIsOneOf`](../common/require_if_attribute_is_one_of.md) - This validator is used to require the attribute if another attribute is one of the given values.
- [`RequireIfAttributeIsSet`](../common/require_if_attribute_is_set.md) - This validator is used to require the attribute if another attribute is set.
- [`NullIfAttributeIsOneOf`](../common/null_if_attribute_is_one_of.md) - This validator is used to verify the attribute value is null if another attribute is one of the given values.
- [`NullIfAttributeIsSet`](../common/null_if_attribute_is_set.md) - This validator is used to verify the attribute value is null if another attribute is set.

## Special

- [`Not`](not.md) - This validator is used to negate the result of another validator.

## Generic

### String

Hashicorp provides a generic validator for strings. It uses the validators already defined in string to validate a list of strings.
It is available in the [hashicorp stringvalidator](https://github.com/hashicorp/terraform-plugin-framework-validators/tree/main) package.

Example of usage:

```go
    // Used within a Schema method of a DataSource, Provider, or Resource
    _ = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "example_attr": schema.MapAttribute{
                ElementType: types.StringType,
                Required:    true,
                Validators: []validator.Map{
                    // Validate this Map must contain string values which are at least 3 characters.
                    mapvalidator.ValueStringsAre(fstringvalidator.IsUUID()),
                },
            },
        },
    }
```
