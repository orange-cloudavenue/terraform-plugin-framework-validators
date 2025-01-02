---
hide:
    - navigation
---

# `AttributeIsDivisibleByAnInteger`

!!! quote inline end "Released in v1.4.0"

This validator is used to check if the attribute is divisible by an integer.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "cpus": schema.Int64Attribute{
                Optional:            true,
                MarkdownDescription: "Number of CPUs",
            },
            "cpus_cores": schema.Int64Attribute{
                Optional:            true,
                MarkdownDescription: "Number of CPUs cores",
                Validators: []validator.Int64{
                    fint64validator.AttributeIsDivisibleByAnInteger(path.MatchRoot("cpus"))
                },
            },
```
