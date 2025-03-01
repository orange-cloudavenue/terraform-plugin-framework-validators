---
hide:
    - navigation
---
# `OneOfWithDescription`

!!! quote inline end "Released in v1.5.0"

This validator is used to check if the string is one of the given values and format the description and the markdown description.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "interface_type": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Type of ...",
                Validators: []validator.Int64{
                    fint64validator.OneOfWithDescription(
                        fint64validator.OneOfWithDescriptionValues{
                            Value:       1,
                            Description: "Vmware interface",
                        },
                        fint64validator.OneOfWithDescriptionValues{
                            Value:       2,
                            Description: "Intel interface",
                        },
                    )
                },
            },
```

## Description and Markdown description

### Description

Value must be one of: "1" (Vmware interface), "2" (Intel interface)

### Markdown description

Value must be one of:

- `1` (Vmware interface)
- `2` (Intel interface)
