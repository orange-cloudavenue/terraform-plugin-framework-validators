---
hide:
    - navigation
---
# `ZeroRemainder`

!!! quote inline end "Released in v1.4.0"

This validator checks if the configured attribute is divisible by a specified integer X, and has zero remainder.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "memory": schema.Int64Attribute{
                Optional:            true,
                MarkdownDescription: "Memory",
                Validators: []validator.Int64{
                    fint64validator.ZeroRemainder(4),
                },
            },
```
