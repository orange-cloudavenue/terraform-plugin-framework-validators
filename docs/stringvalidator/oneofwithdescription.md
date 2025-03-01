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
                Validators: []validator.String{
                    fstringvalidator.OneOfWithDescription(
                        fstringvalidator.OneOfWithDescriptionValues{
                            Value:       "VMXNET3",
                            Description: "Vmware interface",
                        },
                        fstringvalidator.OneOfWithDescriptionValues{
                            Value:       "E1000E",
                            Description: "Intel interface",
                        },
                    )
                },
            },
```

## Description and Markdown description

### Description

Value must be one of: "VMXNET3" (Vmware interface), "E1000E" (Intel interface)

### Markdown description

Value must be one of:

- `VMXNET3` (Vmware interface)
- `E1000E` (Intel interface)
