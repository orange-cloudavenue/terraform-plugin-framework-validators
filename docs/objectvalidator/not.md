---
hide:
    - navigation
---
# `Not`

!!! quote inline end "Released in v1.11.0"

This validator is used to check if the validators passed as arguments are NOT met.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "my_object": schema.SingleNestedAttribute{
                Optional:            true,
                MarkdownDescription: "My Object ...",
                Validators:          []validator.Object{
                    fobjectvalidator.Not(...)
                },
            },
```
