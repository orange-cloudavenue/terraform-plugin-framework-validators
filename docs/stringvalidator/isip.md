---
hide:
    - navigation
---
# `IsIP`

!!! danger inline end "Deprecated"
    Use [IsNetwork](./isnetwork.md) instead.

This validator is used to check if the string is a valid IP address.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "ip_address": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "IPV4 for ...",
                Validators: []validator.String{
                    fstringvalidator.IsIP()
                },
            },
```
