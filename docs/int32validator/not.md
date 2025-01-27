---
hide:
    - navigation
---
# `Not`

!!! quote inline end "Released in v1.13.0"

This validator is used to check if the validators passed as arguments are NOT met.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "count": schema.Int32Attribute{
                Optional:            true,
                MarkdownDescription: "Count of ...",
                Validators: []validator.Int32{
                    fint32validator.Not(int32validator.Between(10,20))
                },
            },
```
