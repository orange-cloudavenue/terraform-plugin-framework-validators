---
hide:
    - navigation
---
# `IsMacAddress`

!!! quote inline end "Released in v1.2.0"

This validator is used to check if the string is a valid Mac Address.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "mac_address": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Mac Adresse for ...",
                Validators: []validator.String{
                    fstringvalidator.IsMacAddress()
                },
            },
```
