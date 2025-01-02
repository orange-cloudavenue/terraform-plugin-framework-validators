---
hide:
    - navigation
---
# `NullIfAttributeIsSet`

!!! quote inline end "Released in v1.8.0"

This validator is used to verify the attribute value is null if another attribute is set.
Set could mean either the attribute is present in the configuration.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "network_type": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Network type ...",
                Validators: []validator.String{
                    fstringvalidator.OneOf("public", "private"),
                },
            },
            "enabled": schema.BoolAttribute{
                Optional:            true,
                MarkdownDescription: "Enable ...",
                Validators: []validator.Bool{
                    fboolvalidator.NullIfAttributeIsSet(path.MatchRoot("network_type"))
                },
            },
```

## Example of generated documentation

If the value of [`network_type`](#network_type) attribute is set this attribute is **NULL**.
