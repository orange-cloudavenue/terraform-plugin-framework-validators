---
hide:
    - navigation
---
# `IsUUID`

!!! quote inline end "Released in v1.1.0"

!!! warning

    Now this validator is deprecated, please use `formats.IsUUIDv4` instead

This validator is used to check if the string is a valid (v4) UUID.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "uuid": schema.StringAttribute{  
            Optional:            true,  
            MarkdownDescription: "Unique identifier (UUID v4) for the resource.",  
            Validators: []validator.String{
                fstringvalidator.IsValidUUID()
            },
        },
        (...)
    }
}
```
