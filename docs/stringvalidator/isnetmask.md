---
hide:
    - navigation
---
# `IsNetmask`

!!! quote inline end "Released in v1.1.0"

This validator is used to check if the string is a valid netmask.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "netmask": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Netmask for ...",
                Validators: []validator.String{
                    fstringvalidator.IsValidNetmask()
                },
            },
```
