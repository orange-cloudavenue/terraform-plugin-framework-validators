---
hide:
    - navigation
---

# `AttributeIsDivisibleByAnInteger`

!!! quote inline end "Released in v1.13.0"

This validator is used to check if the attribute is divisible by an integer.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "cpus": schema.Int32Attribute{
                Optional:            true,
                MarkdownDescription: "Number of CPUs",
            },
            "cpus_cores": schema.Int32Attribute{
                Optional:            true,
                MarkdownDescription: "Number of CPUs cores",
                Validators: []validator.Int32{
                    fint32validator.AttributeIsDivisibleByAnInteger(path.MatchRoot("cpus"))
                },
            },
```
