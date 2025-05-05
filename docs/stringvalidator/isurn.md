---
hide:
    - navigation
---
# `IsURN`

!!! quote inline end "Released in v1.1.0"

!!! warning

    Now this validator is deprecated, please use `formats.IsURN` instead

This validator is used to check if the string is a valid URN.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "urn": schema.StringAttribute{
            Optional:            true,
            MarkdownDescription: "Uniform Resource Name (URN) for the resource.",
            Validators: []validator.String{
                fstringvalidator.IsValidURN()
            },
        },
        (...)
    }
}
```
